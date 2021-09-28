package model_api

import "goskeleton/app/model"

type FlowPosRelModel struct {
	model.BaseModel
	Id                int64  `gorm:"column:id" json:"id"`
	FlowId            int64  `gorm:"column:flow_id" json:"flow_id"`                           // 流量配置id
	AppKey            string `gorm:"column:app_key" json:"app_key"`                           //
	AdType            int    `gorm:"column:ad_type" json:"ad_type"`                           // 广告类型 1视频2插页
	PosKey            string `gorm:"column:pos_key" json:"pos_key"`                           // 广告位id
	AdsId             string `gorm:"column:ads_id" json:"ads_id"`                             // 广告商ID
	ThirdPartyBlockId string `gorm:"column:third_party_block_id" json:"third_party_block_id"` // 第三方广告位ID
	Del               int    `gorm:"column:del" json:"del"`                                   // 删除标志
	model.TimeColumns
}

func FlowPosRelDB() *FlowPosRelModel {
	return &FlowPosRelModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Api")}}
}

func (f *FlowPosRelModel) TableName() string {
	return "flow_pos_rel"
}
