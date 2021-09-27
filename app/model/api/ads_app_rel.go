package model_api

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

type AdsAppRelModel struct {
	model.BaseModel `json:"-"`
	Id              int64  `gorm:"column:id" json:"id"`
	AdsId           string `gorm:"column:ads_id" json:"ads_id"` // 广告商id
	/**
	  `app_name` varchar(30) NOT NULL DEFAULT '' COMMENT '应用名称',
	  `app_key` varchar(30) NOT NULL DEFAULT '' COMMENT '应用appkey',
	  `platform` tinyint(2) NOT NULL DEFAULT '1' COMMENT '平台1安卓2ios',
	  `ad_sub_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '1视频广告,2插图广告,3,自定义，4开屏广告,5原生信息流'
	  `third_party_app_key` varchar(300) NOT NULL DEFAULT '' COMMENT '第三方的appkey',
	  `third_party_secret` varchar(300) NOT NULL DEFAULT '' COMMENT '第三方的密钥',
	  `third_party_report_id` varchar(300) NOT NULL DEFAULT '' COMMENT '第三方报表id',
	  `play_network` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0wifi下，1全网',
	  `life_cycle` int(11) NOT NULL DEFAULT '1800' COMMENT '生命周期，单位秒',
	  `is_show_view` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否显示悬浮窗口 1显示，0不显示',
	  `show_view_time` int(11) NOT NULL DEFAULT '0' COMMENT '显示悬浮窗口的时间 单位为秒',
	  `is_use_template` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否使用模板',
	  `template_show_time` int(10) NOT NULL DEFAULT '3' COMMENT '模板展示时间',
	  `template_id` int(10) DEFAULT NULL COMMENT '模板id',
	  `desc` varchar(30) NOT NULL DEFAULT '' COMMENT '描述',
	  `interval` int(11) NOT NULL DEFAULT '15' COMMENT '轮播时间，单位秒',
	  `version` int(11) NOT NULL DEFAULT '0' COMMENT '配置版本号',
	*/
	model.TimeColumns
}

func AdsAppRelDB() *AdsAppRelModel {
	return &AdsAppRelModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Api")}}
}

func (a *AdsAppRelModel) TableName() string {
	return "ads_app_rel"
}

func (a *AdsAppRelModel) GetsBy(query *tool.WhereQuery, order string) (list []*AdsAppRelModel) {
	if order == "" {
		order = "update_time desc"
	}
	a.Table(a.TableName()).Where(query.QuerySql, query.QueryParams...).Order(order).Find(&list)
	return
}
