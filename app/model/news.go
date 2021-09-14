package model

import (
	"goskeleton/app/global/variable"
	"goskeleton/app/model/tool"
	"time"
)

func NewsDB() *NewsModel {
	return &NewsModel{BaseModel: BaseModel{DB: UseDbConn("")}}
}

type NewsModel struct {
	BaseModel `json:"-"`
	UserId    int    `gorm:"column:user_id" json:"user_id"`
	Title     string `json:"title"`
	Des       string `json:"des"`
	Content   string `json:"content"`
	State     int    `json:"state"`
	*BaseColumns
}

// TableName 表名
func (n *NewsModel) TableName() string {
	return "news"
}

// SearchNews ..
func (n *NewsModel) SearchNews(query *tool.WhereQuery, pageStart, limit int) (counts int, news []NewsModel) {
	if counts = n.count(query); counts > 0 {
		query.QueryParams = append(query.QueryParams, pageStart, limit)
		n.Raw("SELECT * FROM news WHERE "+query.QuerySql+" LIMIT ?,?", query.QueryParams...).Order("updated_at desc").Find(&news)
	}
	return
}

// Store 新增
func (n *NewsModel) Store(title, des, content string, userId int) bool {
	sql := "INSERT INTO news(user_id,title,des,content,created_at) VALUES (?,?,?,?,?)"
	if n.Exec(sql, userId, title, des, content, time.Now().Format(variable.DateFormat)).RowsAffected > 0 {
		return true
	}
	return false
}

// count 查询符合条件的总数
func (n *NewsModel) count(query *tool.WhereQuery) (counts int) {
	n.Raw("SELECT COUNT(*) AS counts FROM news WHERE "+query.QuerySql, query.QueryParams...).First(&counts)
	return
}
