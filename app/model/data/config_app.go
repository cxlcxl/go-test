package data

import "goskeleton/app/model"

type ConfigAppModel struct {
	model.BaseModel
	AppKey    string `json:"app_key" gorm:"column:app_key"`
	AppId     int64  `json:"app_id"`
	AppName   string `json:"app_name"`
	Status    int    `json:"status"` // 应用状态1开启0关闭
	Type      int    `json:"type"`
	AppcateId int    `json:"appcate_id"`
	Platform  int    `json:"platform"`
	Developer int64  `json:"developer"`
	DgcGameId int    `json:"dgc_game_id"` // DGC游戏ID
	SynFlag   int    `json:"syn_flag"`    // 同步标志
	AppType   int    `json:"app_type"`    // 应用自定义类型1休闲游戏2独立游戏3联盟流量
	SspId     int64  `json:"ssp_id"`      // 开发者ID
}

func ConfigAppDB() *ConfigAppModel {
	return &ConfigAppModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Data")}}
}

func (c *ConfigAppModel) TableName() string {
	return "config_app"
}
