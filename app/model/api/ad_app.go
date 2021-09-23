package model_api

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

type AdAppModel struct {
	model.BaseModel `json:"-"`
	AppBaseInfo
	PackageName     string `json:"package_name" gorm:"column:package_name"`         // 包名
	AppDesc         string `json:"app_desc" gorm:"column:app_desc"`                 // 应用描述
	AppcateId       int    `json:"appcate_id" gorm:"column:appcate_id"`             // 应用类型
	State           int    `json:"state" gorm:"column:state"`                       // 状态1为开启,0为关闭
	DeveloperId     int64  `json:"dev_id" gorm:"column:dev_id"`                     // 所属开发者id
	Operator        string `json:"operator" gorm:"column:operator"`                 //
	AcountingMethod int    `json:"acounting_method" gorm:"column:acounting_method"` // 核算方式1:CPM  2:cpc
	AppType         int    `json:"app_type" gorm:"column:app_type"`                 // 应用自定义类型1休闲游戏2独立游戏3联盟流量
	model.TimeColumns
	/*
	  `income_rate` float NOT NULL DEFAULT '0.8' COMMENT '收入调整比例百分比',
	  `denominated` int(11) NOT NULL DEFAULT '0' COMMENT '计价单位分',
	  `source` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0为后台系统添加,1为网站开发者注册的应用',
	  `keyword` varchar(150) DEFAULT '' COMMENT '关键字',
	  `check_msg` varchar(100) DEFAULT '' COMMENT '审批意见',
	  `apk_url` varchar(200) DEFAULT '',
	  `out_game_id` int(10) NOT NULL DEFAULT '0' COMMENT '外部的游戏ID',
	  `is_track` tinyint(2) DEFAULT '0' COMMENT '是否接入监控，0：否，1：是',
	  `delivery_type` tinyint(2) NOT NULL DEFAULT '1' COMMENT '投放类型：1:普通，2:代理(SDK不植入)，3:代理(SDK植入)',
	  `appstore_id` int(11) DEFAULT '0' COMMENT '应用appstore_id',
	  `consumer_key` varchar(32) DEFAULT NULL COMMENT '应用key(投放)',
	  `way` tinyint(2) NOT NULL DEFAULT '1' COMMENT '接入方式',
	  `screen_direction` tinyint(2) NOT NULL DEFAULT '0' COMMENT '屏幕方向 0:未知1:横屏2:竖屏3:横竖屏都支持',
	  `is_change` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否有更新 0:否1:是',
	  `sha1_value` varchar(200) DEFAULT '' COMMENT 'sha1值',
	  `extra` text COMMENT '自定义配置'
	*/
}

type AppBaseInfo struct {
	AppId    int64  `json:"app_id" gorm:"column:app_id"`
	AppName  string `json:"app_name" gorm:"column:app_name"`   // 应用名称
	AppKey   string `json:"app_key" gorm:"column:app_key"`     // 应用key
	Platform int    `json:"platform" gorm:"column:platform"`   // 1:Android 2:IOS
	IsCheck  int    `json:"is_check" gorm:"column:is_check"`   // -1-未通过 1-通过,2为申请中,3为编辑后再申请
	IsOnline int    `json:"is_online" gorm:"column:is_online"` // 1-审核中 2-已上线 3-未上线
	TestMode int    `json:"test_mode" gorm:"column:test_mode"` // 测试模式 0:关闭1:打开
	IsConfig int    `json:"is_config"`                         // 是否已经配置 1 已配置
	Icon     string `json:"icon"`                              // 图标
}

// AdAppDB 创建DB对象
func AdAppDB() *AdAppModel {
	return &AdAppModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Api")}}
}

// TableName 表名
func (a *AdAppModel) TableName() string {
	return "ad_app"
}

// AppList ...
func (a *AdAppModel) AppList(query *tool.WhereQuery, limitStart, size int) (apps []*AppBaseInfo, counts int) {
	if counts = a.counts(query); counts > 0 {
		a.Table(a.TableName()).Where(query.QuerySql, query.QueryParams...).Order("update_time desc").Offset(limitStart).Limit(size).Find(&apps)
		return
	} else {
		return
	}
}

// Show ...
func (a *AdAppModel) Show(query *tool.WhereQuery, limitStart, size int) (apps []*AdAppModel, counts int) {
	if counts = a.counts(query); counts > 0 {
		query.QueryParams = append(query.QueryParams, limitStart, size)
		a.Raw("SELECT * FROM `"+a.TableName()+"` WHERE "+query.QuerySql+" LIMIT ?,?", query.QueryParams...).Order("update_time desc").Find(&apps)
		return
	} else {
		return
	}
}

func (a *AdAppModel) counts(query *tool.WhereQuery) (counts int) {
	a.Raw("SELECT COUNT(*) AS counts FROM `"+a.TableName()+"` WHERE "+query.QuerySql+" LIMIT 1", query.QueryParams...).First(&counts)
	return
}
