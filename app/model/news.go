package model

import (
	"goskeleton/app/global/variable"
	"time"
)

// 操作数据库喜欢使用gorm自带语法的开发者可以参考 GinSkeleton-Admin 系统相关代码
// Admin 项目地址：https://gitee.com/daitougege/gin-skeleton-admin-backend/
// gorm_v2 提供的语法+ ginskeleton 实践 ：  http://gitee.com/daitougege/gin-skeleton-admin-backend/blob/master/app/model/button_cn_en.go

// 参数说明： 传递空值，默认使用 配置文件选项：UseDbType（mysql）

func CreateNewFactory(sqlType string) *NewsModel {
	return &NewsModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type NewsModel struct {
	BaseModel `json:"-"`
	UserId    int    `gorm:"column:user_id" json:"user_id"`
	Title     string `json:"title"`
	Des       string `json:"des"`
	Content   string `json:"content"`
	State     int    `json:"state"`
}

// TableName 表名
func (n *NewsModel) TableName() string {
	return "news"
}

// Store 新增
func (n *NewsModel) Store(title, des, content string, userId int) bool {
	sql := "INSERT INTO news(user_id,title,des,content,created_at) VALUES (?,?,?,?,?)"
	if n.Exec(sql, userId, title, des, content, time.Now().Format(variable.DateFormat)).RowsAffected > 0 {
		return true
	}
	return false
}
