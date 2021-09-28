package model

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/model/tool"
	"goskeleton/app/utils/md5_encrypt"
	"log"
	"strconv"
	"time"

	"go.uber.org/zap"
)

func CreateUserFactory() *UsersModel {
	return &UsersModel{BaseModel: BaseModel{DB: UseDbConn("")}}
}

type UsersModel struct {
	BaseModel
	UserBaseInfo
	Password     string `json:"-" gorm:"password"`
	Operator     int    `json:"operator"`
	RoleId       int64  `json:"role_id" gorm:"column:role_id"`
	UserType     int    `json:"user_type"`
	Token        string `json:"token"`
	GroupId      int    `json:"group_id"`
	IsParent     int    `json:"is_parent"`
	IsCheck      int    `json:"is_check" gorm:"column:is_check"`
	RegisterType int    `json:"register_type" gorm:"column:register_type"`
	CardId       string `json:"card_id" gorm:"column:card_id"`
	Contact      string `json:"contact" gorm:"column:contact"`
	Address      string `json:"address" gorm:"column:address"`
	ParentId     int64  `json:"parent_id"`
	TimeColumns
	UserExtendColumn
}

type UserBaseInfo struct {
	Id       int64  `json:"user_id" gorm:"column:user_id"`
	UserName string `json:"user_name" gorm:"column:user_name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	IsLock   int    `json:"is_lock"`
	IsAdmin  int    `json:"is_admin"`
}

type UserExtendColumn struct {
	ParentName   string `json:"parent_name"`
	UserTypeName string `json:"user_type_name"`
}

// TableName 表名
func (u *UsersModel) TableName() string {
	return "admin_user"
}

// Store 用户注册（写一个最简单的使用账号、密码注册即可）
func (u *UsersModel) Store(values map[string]interface{}) bool {
	result := u.Table(u.TableName()).Create(values)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// Login 用户登录
func (u *UsersModel) Login(userName, pass, ip string) (*UsersModel, string) {
	result := u.Table(u.TableName()).Where("user_name=? or email=?", userName, userName).Limit(1).First(u)
	if result.Error == nil {
		// 账号密码验证成功
		if u.IsLock == 1 {
			return nil, "账号已被停用"
		}
		if u.Password == md5_encrypt.MobgiPwd(pass) {
			return u, ""
		} else {
			return nil, consts.CurdLoginFailPassErr
		}
	} else {
		return nil, consts.CurdLoginFailUserNameErr
	}
}

// OauthLoginToken 记录用户登陆（login）生成的token，每次登陆记录一次token
func (u *UsersModel) OauthLoginToken(userId int64, token string, expiresAt int64, clientIp string) bool {
	sql := "INSERT INTO `oauth_access_tokens`(fr_user_id,`action_name`,token,expires_at,client_ip) " +
		"SELECT ?,'login',? ,?,? FROM DUAL WHERE NOT EXISTS(" + // 方便 select 的使用，不涉及表的时候可以使用 DUAL 当成虚拟表
		"SELECT 1 FROM `oauth_access_tokens` a WHERE a.fr_user_id=? AND a.action_name='login' AND a.token= ?" +
		")"
	//注意：token的精确度为秒，如果在一秒之内，一个账号多次调用接口生成的token其实是相同的，这样写入数据库，第二次的影响行数为0，知己实际上操作仍然是有效的。
	//所以这里只判断无错误即可，判断影响行数的话，>=0 都是ok的
	if u.Exec(sql, userId, token, time.Unix(expiresAt, 0).Format(variable.DateFormat), clientIp, userId, token).Error == nil {
		return true
	}
	return false
}

// OauthRefreshToken 用户刷新token
func (u *UsersModel) OauthRefreshToken(userId, expiresAt int64, oldToken, newToken, clientIp, platform string) bool {
	sql := "UPDATE oauth_access_tokens SET token=? ,expires_at=?,client_ip=?,updated_at=NOW(),action_name='refresh' WHERE fr_user_id=? AND token=?"
	if u.Exec(sql, newToken, time.Unix(expiresAt, 0).Format(variable.DateFormat), clientIp, userId, oldToken).Error == nil {
		go CreateLoginLogFactory().LogLogin(userId)
		return true
	}
	return false
}

// UpdateUserLoginInfo 更新用户登陆次数、最近一次登录ip、最近一次登录时间
func (u *UsersModel) UpdateUserLoginInfo(userId int64) {
	sql := "UPDATE `" + u.TableName() + "` SET last_login_time=? WHERE id=?"
	_ = u.Exec(sql, time.Now().Format(variable.DateFormat), userId)
}

// OauthResetToken 当用户更改密码后，所有的token都失效，必须重新登录
func (u *UsersModel) OauthResetToken(userId float64, newPass string) bool {
	//如果用户新旧密码一致，直接返回true，不需要处理
	userItem, err := u.ShowOneItem(userId)
	if userItem != nil && err == nil && userItem.Password == newPass {
		return true
	} else if userItem != nil {
		sql := "DELETE FROM oauth_access_tokens WHERE fr_user_id=?"
		if u.Exec(sql, userId).Error == nil {
			return true
		}
	}
	return false
}

// OauthDestroyToken 当users 删除数据，相关的token同步删除
func (u *UsersModel) OauthDestroyToken(userId float64) bool {
	// 如果用户新旧密码一致，直接返回true，不需要处理
	sql := "DELETE FROM oauth_access_tokens WHERE fr_user_id=?"
	// 判断>=0, 有些没有登录过的用户没有相关token，此语句执行影响行数为0，但是仍然是执行成功
	if u.Exec(sql, userId).Error == nil {
		return true
	}
	return false
}

// OauthCheckTokenIsOk 判断用户token是否在数据库存在+状态OK
func (u *UsersModel) OauthCheckTokenIsOk(userId int64, token string) bool {
	sql := "SELECT token FROM `oauth_access_tokens` WHERE fr_user_id=? AND revoked=0 AND expires_at>NOW() ORDER BY expires_at DESC, updated_at  DESC  LIMIT ?"
	maxOnlineUsers := variable.ConfigYml.GetInt("Token.JwtTokenOnlineUsers")
	rows, err := u.Raw(sql, userId, maxOnlineUsers).Rows()
	if err == nil && rows != nil {
		for rows.Next() {
			var tempToken string
			err := rows.Scan(&tempToken)
			if err == nil {
				if tempToken == token {
					_ = rows.Close()
					return true
				}
			}
		}
		// 凡是查询类记得释放记录集
		_ = rows.Close()
	}
	return false
}

// SetTokenInvalid 禁用一个用户的: 1.users表的 status 设置为 0，oauth_access_tokens 表的所有token删除
// 禁用一个用户的token请求（本质上就是把users表的 status 字段设置为 0 即可）
func (u *UsersModel) SetTokenInvalid(userId int) bool {
	sql := "delete from `oauth_access_tokens` where `fr_user_id`=?  "
	if u.Exec(sql, userId).Error == nil {
		if u.Exec("update `"+u.TableName()+"` set status=0 where id=?", userId).Error == nil {
			return true
		}
	}
	return false
}

// ShowOneItem 根据用户ID查询一条信息
func (u *UsersModel) ShowOneItem(userId float64) (*UsersModel, error) {
	sql := "SELECT `id`, `user_name`,`pass`, `real_name`, `phone`, `status` FROM `" + u.TableName() + "` WHERE id=? LIMIT 1"
	result := u.Raw(sql, userId).First(u)
	if result.Error == nil {
		return u, nil
	} else {
		return nil, result.Error
	}
}

// 查询数据之前统计条数
func (u *UsersModel) counts(where *tool.WhereQuery) (counts int64) {
	if res := u.Table(u.TableName()).Where(where.QuerySql, where.QueryParams...).Count(&counts); res.Error != nil {
		variable.ZapLog.Error("UsersModel - counts 查询数据条数出错", zap.Error(res.Error))
	}
	return counts
}

// Show 查询（根据关键词模糊查询）
func (u *UsersModel) Show(where *tool.WhereQuery, limitStart, size int) (counts int64, temp []*UsersModel) {
	if counts = u.counts(where); counts > 0 {
		if u.Table(u.TableName()).Where(where.QuerySql, where.QueryParams...).Order("update_time desc").Offset(limitStart).Limit(size).Find(&temp).RowsAffected > 0 {
			return counts, temp
		}
	}
	log.Println(counts)
	return 0, nil
}

// Update 更新
func (u *UsersModel) Update(values map[string]interface{}) bool {
	if u.Table(u.TableName()).Where("id=?", values["id"]).Updates(values).RowsAffected >= 0 {
		if u.OauthResetToken(values["id"].(float64), values["pass"].(string)) {
			return true
		}
	}
	return false
}

// Destroy 删除用户以及关联的token记录
func (u *UsersModel) Destroy(id float64) bool {
	if u.Table(u.TableName()).Delete(u, id).Error == nil {
		if u.OauthDestroyToken(id) {
			return true
		}
	}
	return false
}

// Logout 删除用户以及关联的token记录
func (u *UsersModel) Logout(token string, id float64) bool {
	sql := "delete from `oauth_access_tokens` where `fr_user_id`=? and token=?"
	if u.Exec(sql, id, token).Error == nil {
		return true
	}
	return false
}

// UserIsExists 注册/添加检测是否已存在
func (u *UsersModel) UserIsExists(userId float64, userName, email, phone string) string {
	whereSlice := map[string]interface{}{
		"or": map[string]interface{}{
			"user_name": userName,
			"email":     email,
			"phone":     phone,
		},
	}
	if userId > 0 {
		whereSlice["id"] = []string{"!=", strconv.Itoa(int(userId))}
	}
	where := (&tool.WhereQuery{Filter: true}).GenerateWhere(whereSlice)
	var user UsersModel
	if result := u.Table(u.TableName()).Where(where.QuerySql, where.QueryParams...).First(&user); result.Error != nil {
		return "用户数据查询出错，请重试"
	} else {
		if len(email) > 0 && user.Email == email {
			return "用户邮箱已存在"
		}
		if len(phone) > 0 && user.Mobile == phone {
			return "手机号码已存在"
		}
		if user.UserName == userName {
			return "用户名称已存在"
		}
	}

	return ""
}

// CheckPass 检查密码
func (u *UsersModel) CheckPass(id float64, pass string) bool {
	originalPass := ""
	u.Table(u.TableName()).Select("pass").Where("id = ?", id).Limit(1).First(&originalPass)
	return pass == originalPass
}

// ResetPass 修改密码
func (u *UsersModel) ResetPass(id float64, pass string) bool {
	err := u.Exec("UPDATE `"+u.TableName()+"` SET pass = ? WHERE id = ? LIMIT 1", pass, id).Error
	if err == nil {
		u.OauthDestroyToken(id)
		return true
	}
	return false
}

// EmailExists 检查密码
func (u *UsersModel) EmailExists(email string) (userId int) {
	u.Table(u.TableName()).Where("email = ?", email).Limit(1).First(&userId)
	return
}

// GetsByIds 以ID查询用户信息
func (u *UsersModel) GetsByIds(ids []int64) (users []*UserBaseInfo) {
	u.Table(u.TableName()).Where("user_id in ?", ids).Find(&users)
	return
}

// UpdateGroup 修改用户组
func (u *UsersModel) UpdateGroup(ids []int64, groupId float64) bool {
	return u.Table(u.TableName()).Where("user_id in ?", ids).Update("group_id", groupId).RowsAffected > 0
}

// DoCheckUser 审核用户
func (u *UsersModel) DoCheckUser(userId int64, values map[string]interface{}) bool {
	return u.Table(u.TableName()).Where("user_id = ?", userId).Updates(values).RowsAffected > 0
}
