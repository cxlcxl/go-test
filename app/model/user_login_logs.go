package model

import (
	"goskeleton/app/global/variable"
	"time"
)

func CreateLoginLogFactory() *LoginLogsModel {
	return &LoginLogsModel{BaseModel: BaseModel{DB: UseDbConn("")}}
}

type LoginLogsModel struct {
	BaseModel `json:"-"`
	ID        int    `json:"id"`
	LoginTime string `gorm:"column:login_time" json:"login_time"`
	UserId    int    `gorm:"column:user_id" json:"user_id"`
}

// TableName 表名
func (l *LoginLogsModel) TableName() string {
	return "admin_user_login_logs"
}

// LogLogin 写登陆记录
func (l *LoginLogsModel) LogLogin(loginUserId int64) bool {
	sql := "INSERT INTO `" + l.TableName() + "`(login_time,user_id) VALUES (?,?)"
	result := l.Exec(sql, time.Now().Format(variable.DateFormat), loginUserId)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
