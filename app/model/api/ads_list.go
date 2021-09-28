package model_api

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

type AdsListModel struct {
	model.BaseModel  `json:"-"`
	Id               int64  `gorm:"column:id" json:"id"`
	AdsId            string `gorm:"column:ads_id" json:"ads_id"`                           // 视频广告商编号
	Name             string `gorm:"column:name" json:"name"`                               // 视频广告商名称
	AdType           int    `gorm:"column:ad_type" json:"ad_type"`                         // 聚合类型：1聚合广告商2渠道广告商3DSP广告商
	AdSubType        string `gorm:"column:ad_sub_type" json:"ad_sub_type"`                 // 支持广告类型 1视频广告2插图广告 3自定义，4开屏 5原生流式
	AdSubTypeVersion string `gorm:"column:ad_sub_type_version" json:"ad_sub_type_version"` // 广告类型版本
	OutUrl           string `gorm:"column:out_url" json:"out_url"`                         // 外部的跳转的url
	SettlementMethod int    `gorm:"column:settlement_method" json:"settlement_method"`     // 结算方式,1cpm,2cpc
	GroupId          int    `gorm:"column:group_id" json:"group_id"`                       // 广告商组id
	Del              int    `gorm:"column:del" json:"del"`                                 // 软删除 0正常 1删除
	IsForeign        int    `gorm:"column:is_foreign" json:"is_foreign"`                   // 是否国外0国内1国外
	IsBid            int    `gorm:"column:is_bid" json:"is_bid"`                           // 是否实时竞价 1 是 0否【enum枚举类型】
	InterfaceUrl     string `gorm:"column:interface_url" json:"interface_url"`             // 接口地址
	BidPrice         string `gorm:"column:bid_price" json:"bid_price"`                     // 竞价价格
	model.TimeColumns
}

func AdsListDB() *AdsListModel {
	return &AdsListModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Api")}}
}

func (a *AdsListModel) TableName() string {
	return "ads_list"
}

func (a *AdsListModel) GetsBy(query *tool.WhereQuery, order string) (list []*AdsListModel) {
	if order == "" {
		order = "update_time desc"
	}
	a.Table(a.TableName()).Where(query.QuerySql, query.QueryParams...).Order(order).Find(&list)
	return
}
