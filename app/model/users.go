package model

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/model/tool"
	"goskeleton/app/utils/md5_encrypt"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// 操作数据库喜欢使用gorm自带语法的开发者可以参考 GinSkeleton-Admin 系统相关代码
// Admin 项目地址：https://gitee.com/daitougege/gin-skeleton-admin-backend/
// gorm_v2 提供的语法+ ginskeleton 实践 ：  http://gitee.com/daitougege/gin-skeleton-admin-backend/blob/master/app/model/button_cn_en.go

// 创建 userFactory
// 参数说明： 传递空值，默认使用 配置文件选项：UseDbType（mysql）

func CreateUserFactory(sqlType string) *UsersModel {
	return &UsersModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type UsersModel struct {
	BaseModel     `json:"-"`
	UserName      string         `json:"user_name" gorm:"column:user_name"`
	Pass          string         `json:"-"`
	Email         string         `json:"email"`
	Phone         string         `json:"phone"`
	Avatar        string         `json:"avatar"`
	RealName      string         `json:"real_name" gorm:"column:real_name"`
	RoleId        int64          `json:"role_id" gorm:"column:role_id"`
	Status        int            `json:"status"`
	Token         string         `json:"token"`
	Role          RoleModel      `json:"roles"`
	LastLoginTime tool.LocalTime `json:"last_login_time" gorm:"column:last_login_time"`
	*BaseColumns
}

// TableName 表名
func (u *UsersModel) TableName() string {
	return "users"
}

// Store 用户注册（写一个最简单的使用账号、密码注册即可）
func (u *UsersModel) Store(values map[string]interface{}) bool {
	// sql := "INSERT INTO users(user_name,pass,email,phone,real_name,role_id) VALUES(?,?,?,?,?,?)"
	result := u.Create(values)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// Login 用户登录
func (u *UsersModel) Login(userName, pass, ip string) (*UsersModel, string) {
	sql := "select id,user_name,real_name,pass,phone,avatar,status from users where user_name=? limit 1"
	result := u.Raw(sql, userName).First(u)
	if result.Error == nil {
		// 账号密码验证成功
		if u.Status != 1 {
			return nil, "账号已被停用"
		}
		if u.Pass == md5_encrypt.Base64Md5(pass) {
			return u, ""
		} else {
			go CreateLoginLogFactory("").LogLogin(u.Id, ip, "web", 0)
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
		go CreateLoginLogFactory("").LogLogin(userId, clientIp, platform, 1)
		return true
	}
	return false
}

// UpdateUserLoginInfo 更新用户登陆次数、最近一次登录ip、最近一次登录时间
func (u *UsersModel) UpdateUserLoginInfo(userId int64) {
	sql := "UPDATE users SET last_login_time=? WHERE id=?"
	_ = u.Exec(sql, time.Now().Format(variable.DateFormat), userId)
}

// OauthResetToken 当用户更改密码后，所有的token都失效，必须重新登录
func (u *UsersModel) OauthResetToken(userId float64, newPass string) bool {
	//如果用户新旧密码一致，直接返回true，不需要处理
	userItem, err := u.ShowOneItem(userId)
	if userItem != nil && err == nil && userItem.Pass == newPass {
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
		if u.Exec("update users set status=0 where id=?", userId).Error == nil {
			return true
		}
	}
	return false
}

// ShowOneItem 根据用户ID查询一条信息
func (u *UsersModel) ShowOneItem(userId float64) (*UsersModel, error) {
	sql := "SELECT `id`, `user_name`,`pass`, `real_name`, `phone`, `status` FROM `users` WHERE id=? LIMIT 1"
	result := u.Raw(sql, userId).First(u)
	if result.Error == nil {
		return u, nil
	} else {
		return nil, result.Error
	}
}

// 查询数据之前统计条数
func (u *UsersModel) counts(where *tool.WhereQuery) (counts int64) {
	sql := "SELECT count(*) as counts FROM users WHERE " + where.QuerySql
	if res := u.Raw(sql, where.QueryParams...).First(&counts); res.Error != nil {
		variable.ZapLog.Error("UsersModel - counts 查询数据条数出错", zap.Error(res.Error))
	}
	return counts
}

// Show 查询（根据关键词模糊查询）
func (u *UsersModel) Show(values map[string]interface{}, limitStart, limitItems int) (counts int64, temp []*UsersModel) {
	whereSlice := map[string]interface{}{
		"user_name": []string{"like", values["user_name"].(string)},
	}
	where := (&tool.WhereQuery{}).GenerateWhere(whereSlice)
	if counts = u.counts(where); counts > 0 {
		sql := "SELECT * FROM `users` WHERE " + where.QuerySql + " LIMIT ?,?"
		where.QueryParams = append(where.QueryParams, limitStart, limitItems)
		if res := u.Raw(sql, where.QueryParams...).Order("updated_at desc").Find(&temp); res.RowsAffected > 0 {
			return counts, temp
		}
	}
	return 0, nil
}

// Update 更新
func (u *UsersModel) Update(values map[string]interface{}) bool {
	//sql := "update users set user_name=?,pass=?,real_name=?,phone=?,email=?,remark=?,avatar=?,updated_at=? WHERE id=?"
	if u.Where("id=?", values["id"]).Updates(values).RowsAffected >= 0 {
		if u.OauthResetToken(values["id"].(float64), values["pass"].(string)) {
			return true
		}
	}
	return false
}

// Destroy 删除用户以及关联的token记录
func (u *UsersModel) Destroy(id float64) bool {
	if u.Delete(u, id).Error == nil {
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
	sql := "SELECT * FROM users WHERE " + where.QuerySql

	var user UsersModel
	if result := u.Raw(sql, where.QueryParams...).First(&user); result.Error != nil {
		return "用户数据查询出错，请重试"
	} else {
		if len(email) > 0 && user.Email == email {
			return "用户邮箱已存在"
		}
		if len(phone) > 0 && user.Phone == phone {
			return "手机号码已存在"
		}
		if user.UserName == userName {
			return "用户名称已存在"
		}
	}

	return ""
}
