package model

import (
	"goskeleton/app/global/variable"
	"time"
)

func CreateLoginLogFactory(sqlType string) *LoginLogsModel {
	return &LoginLogsModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type LoginLogsModel struct {
	BaseModel `json:"-"`
	ID        int    `json:"id"`
	LoginIp   string `gorm:"column:login_ip" json:"login_ip"`
	LoginTime string `gorm:"column:login_time" json:"login_time"`
	UserId    int    `gorm:"column:user_id" json:"user_id"`
	Platform  string `gorm:"column:platform" json:"platform"`
	IsSuccess string `gorm:"column:is_success" json:"is_success"`
}

// TableName 表名
func (l *LoginLogsModel) TableName() string {
	return "user_login_logs"
}

// LogLogin 写登陆记录
func (l *LoginLogsModel) LogLogin(loginUserId int64, loginIp, platform string, state int) bool {
	sql := "INSERT INTO user_login_logs(login_ip,login_time,user_id,platform,is_success) VALUES (?,?,?,?,?)"
	result := l.Exec(sql, loginIp, time.Now().Format(variable.DateFormat), loginUserId, platform, state)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// UserLoginFailTimes 用户登陆失败次数
func (l *LoginLogsModel) UserLoginFailTimes(loginUserId int64) (counts int64) {
	sql := "SELECT COUNT(*) AS counts FROM user_login_logs WHERE user_id = ? AND is_success = 0"
	_ = l.Raw(sql, loginUserId).First(&counts)
	return counts
}
