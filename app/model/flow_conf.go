package model

import (
	"fmt"
	"goskeleton/app/global/variable"
	"time"

	"go.uber.org/zap"
)

func CreateFlowConfFactory(sqlType string) *FlowConfModel {
	return &FlowConfModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type FlowConfModel struct {
	BaseModel       `json:"-"`
	ConfType        int    `gorm:"column:conf_type" json:"conf_type"` // 1默认配置，2定向配置
	ConfName        string `gorm:"column:conf_name" json:"conf_name"`
	AppKey          string `gorm:"column:app_key" json:"app_key"`
	UserConfType    int    `gorm:"column:user_conf_type" json:"user_conf_type"`       // 用户行为配置类型0全部，1定向
	UserConf        string `gorm:"column:user_conf" json:"user_conf"`                 // 用户行为配置
	ChannelConfType int    `gorm:"column:channel_conf_type" json:"channel_conf_type"` // 渠道配置类型0全部，1定向
	ChannelConf     string `gorm:"column:channel_conf" json:"channel_conf"`           //
	AreaConfType    int    `gorm:"column:area_conf_type" json:"area_conf_type"`       // 区域配置类型0全部，1定向
	AreaConf        string `gorm:"column:area_conf" json:"area_conf"`                 //
	GameConfType    int    `gorm:"column:game_conf_type" json:"game_conf_type"`       // 游戏配置类型0全部，1定向
	GameConf        string `gorm:"column:game_conf" json:"game_conf"`                 //
	SdkConfType     int    `gorm:"column:sdk_conf_type" json:"sdk_conf_type"`         // sdk类型0全部，1定向
	SdkConf         string `gorm:"column:sdk_conf" json:"sdk_conf"`                   //
	SysConfType     int    `gorm:"column:sys_conf_type" json:"sys_conf_type"`         // 系统版本的类型 0全部1 定向
	SysConf         string `gorm:"column:sys_conf" json:"sys_conf"`                   //
	BrandConfType   int    `gorm:"column:brand_conf_type" json:"brand_conf_type"`     // 品牌定向 0全部1 定向
	BrandConf       string `gorm:"column:brand_conf" json:"brand_conf"`               //
	ConfNum         int    `gorm:"column:conf_num" json:"conf_num"`                   // 条件个数，冗余字段
	OperatorId      int    `gorm:"column:operator_id" json:"operator_id"`             // 操作人
	Del             int    `json:"del"`
	RelVersion      int    `gorm:"column:rel_version" json:"rel_version"`
	IdfaConfType    int    `gorm:"column:idfa_conf_type" json:"idfa_conf_type"`
	IdfaConf        string `gorm:"column:idfa_conf" json:"idfa_conf"`
}

// 表名
func (f *FlowConfModel) TableName() string {
	return "flow_conf"
}

// 用户注册（写一个最简单的使用账号、密码注册即可）
func (u *FlowConfModel) Register(userName, pass, userIp string) bool {
	sql := "INSERT  INTO tb_users(user_name,pass,last_login_ip) SELECT ?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  user_name=?)"
	result := u.Exec(sql, userName, pass, userIp, userName)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// 查询
func (f *FlowConfModel) FetchRows(sql string, query ...interface{}) *FlowConfModel {
	// result := f.Select(sql, query...).Find()
	result := f.Raw(sql, query...).First(f)
	if result.Error == nil {
		fmt.Println(f)
	} else {
		fmt.Println("err0r")
	}

	return f
}

// 查询单条数据
func (f *FlowConfModel) FetchOne(sql string, query ...interface{}) *FlowConfModel {
	result := f.Raw(sql, query...).First(f)
	if result.Error == nil {
		return f
	} else {
		variable.ZapLog.Error("查询出错:", zap.Error(result.Error))
	}
	return nil
}

// 用户刷新token
func (u *FlowConfModel) OauthRefreshToken(userId, expiresAt int64, oldToken, newToken, clientIp string) bool {
	sql := "UPDATE   tb_oauth_access_tokens   SET  token=? ,expires_at=?,client_ip=?,updated_at=NOW(),action_name='refresh'  WHERE   fr_user_id=? AND token=?"
	if u.Exec(sql, newToken, time.Unix(expiresAt, 0).Format(variable.DateFormat), clientIp, userId, oldToken).Error == nil {
		return true
	}
	return false
}
