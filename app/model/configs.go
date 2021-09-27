package model

import (
	"goskeleton/app/model/tool"
)

func ConfigDB() *ConfigModel {
	return &ConfigModel{BaseModel: BaseModel{DB: UseDbConn("")}}
}

type ConfigModel struct {
	BaseModel `json:"-"`
	Id        int64  `gorm:"primarykey" json:"id"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Des       string `json:"des"`
	InUse     int    `json:"in_use"`
	Bak0      string `json:"bak0"`
	Bak1      string `json:"bak1"`
	Bak2      string `json:"bak2"`
	TimeColumns
}

func (c *ConfigModel) TableName() string {
	return "configs"
}

func (c *ConfigModel) Show(query *tool.WhereQuery, pageStart, limit int, orderBy, orderSort string) (counts int, configs []ConfigModel) {
	if counts = c.counts(query); counts > 0 {
		if orderBy == "" {
			orderBy = "updated_at"
		}
		if orderSort == "" {
			orderSort = "desc"
		}
		query.QueryParams = append(query.QueryParams, pageStart, limit)
		c.Raw("SELECT * FROM configs WHERE "+query.QuerySql+" LIMIT ?,?", query.QueryParams...).Order(orderBy + " " + orderSort).Find(&configs)
	}
	return
}

func (c *ConfigModel) counts(query *tool.WhereQuery) (counts int) {
	c.Raw("SELECT COUNT(*) AS counts FROM configs WHERE "+query.QuerySql, query.QueryParams...).First(&counts)
	return
}

func (c *ConfigModel) Info(id int) (config ConfigModel) {
	c.Raw("SELECT * FROM configs WHERE id = ?", id).First(&config)
	return
}

func (c *ConfigModel) Destroy(id int) bool {
	if c.Delete(c, id).Error != nil {
		return false
	}
	return true
}

func (c *ConfigModel) Store(values map[string]interface{}) bool {
	if c.Table(c.TableName()).Create(values).RowsAffected > 0 {
		return true
	}
	return false
}

func (c *ConfigModel) Update(values map[string]interface{}, id int) bool {
	if c.Where("id = ?", id).Updates(values).RowsAffected > 0 {
		return true
	}
	return false
}
