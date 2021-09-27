package model_api

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

type AdsListModel struct {
	model.BaseModel `json:"-"`
	Id              int64  `gorm:"column:id" json:"id"`
	AdsId           string `gorm:"column:ads_id" json:"ads_id"`   // 视频广告商编号
	Name            string `gorm:"column:name" json:"name"`       // 视频广告商名称
	AdType          int    `gorm:"column:ad_type" json:"ad_type"` // 聚合类型：1聚合广告商2渠道广告商3DSP广告商
	/*
	  `ad_sub_type` varchar(100) NOT NULL DEFAULT '' COMMENT '支持广告类型 1视频广告2插图广告 3自定义，4开屏 5原生流式',
	  `ad_sub_type_version` varchar(100) NOT NULL DEFAULT '' COMMENT '广告类型版本',
	  `out_url` varchar(100) NOT NULL DEFAULT '' COMMENT '外部的跳转的url',
	  `settlement_method` tinyint(2) NOT NULL DEFAULT '1' COMMENT '结算方式,1cpm,2cpc',
	  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '广告商组id',
	  `del` int(11) NOT NULL DEFAULT '0' COMMENT '软删除 0正常 1删除',
	  `is_foreign` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否国外0国内1国外',
	  `is_bid` enum('1','0') NOT NULL DEFAULT '1' COMMENT '是否实时竞价 1 是 0否',
	  `interface_url` varchar(100) DEFAULT NULL COMMENT '接口地址',
	  `bid_price` varchar(500) NOT NULL COMMENT '竞价价格',
	*/
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
