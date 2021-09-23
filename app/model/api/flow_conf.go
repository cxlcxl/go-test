package model_api

import (
	"gorm.io/gorm"
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

type FlowConfModel struct {
	model.BaseModel `json:"-"`
	FlowList
	UserConfType    int    `gorm:"column:user_conf_type" json:"user_conf_type"`       // 用户行为配置类型0全部，1定向
	UserConf        string `gorm:"column:user_conf" json:"user_conf"`                 // 用户行为配置
	ChannelConfType int    `gorm:"column:channel_conf_type" json:"channel_conf_type"` // 渠道配置类型0全部，1定向
	ChannelConf     string `gorm:"column:channel_conf" json:"channel_conf"`           //
	AreaConfType    int    `gorm:"column:area_conf_type" json:"area_conf_type"`       // 区域配置类型0全部，1定向
	AreaConf        string `gorm:"column:area_conf" json:"area_conf"`                 //
	GameConfType    int    `gorm:"column:game_conf_type" json:"game_conf_type"`       // 游戏配置类型0全部，1定向
	GameConf        string `gorm:"column:game_conf" json:"game_conf"`                 //
	SdkConfType     int    `gorm:"column:sdk_conf_type" json:"sdk_conf_type"`         // sdk类型0全部，1定向
	SdkConf         string `gorm:"column:sdk_conf" json:"sdk_conf"`                   //
	SysConfType     int    `gorm:"column:sys_conf_type" json:"sys_conf_type"`         // 系统版本的类型 0全部1 定向
	SysConf         string `gorm:"column:sys_conf" json:"sys_conf"`                   //
	BrandConfType   int    `gorm:"column:brand_conf_type" json:"brand_conf_type"`     // 品牌定向 0全部1 定向
	BrandConf       string `gorm:"column:brand_conf" json:"brand_conf"`               //
	ConfNum         int    `gorm:"column:conf_num" json:"conf_num"`                   // 条件个数，冗余字段
	Del             int    `json:"del"`
	IdfaConfType    int    `gorm:"column:idfa_conf_type" json:"idfa_conf_type"`
	IdfaConf        string `gorm:"column:idfa_conf" json:"idfa_conf"`
}

type FlowList struct {
	model.IdColumns
	ConfType       int    `gorm:"column:conf_type" json:"conf_type"` // 1全局配置，2定向配置
	ConfName       string `gorm:"column:conf_name" json:"conf_name"`
	AppKey         string `gorm:"column:app_key" json:"app_key"`
	RelVersion     int    `gorm:"column:rel_version" json:"rel_version"`
	ConfTypeName   string `json:"conf_type_name"`                        // 扩展字段
	RelVersionName string `json:"rel_version_name"`                      // 扩展字段
	OperatorId     int64  `gorm:"column:operator_id" json:"operator_id"` // 操作人
	Operator       string `json:"operator"`                              //
	model.TimeColumns
}

// FlowConfDB 创建DB对象
func FlowConfDB() *FlowConfModel {
	return &FlowConfModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Api")}}
}

// TableName 表名
func (f *FlowConfModel) TableName() string {
	return "flow_conf"
}

// Show ...
func (f *FlowConfModel) Show(query *tool.WhereQuery, limitStart, size int) (apps []*FlowConfModel, counts int) {
	if counts = f.counts(query); counts > 0 {
		query.QueryParams = append(query.QueryParams, limitStart, size)
		f.Raw("SELECT * FROM `"+f.TableName()+"` WHERE "+query.QuerySql+" LIMIT ?,?", query.QueryParams...).Order("update_time desc").Find(&apps)
		return
	} else {
		return
	}
}

func (f *FlowConfModel) counts(query *tool.WhereQuery) (counts int) {
	f.Raw("SELECT COUNT(*) AS counts FROM `"+f.TableName()+"` WHERE "+query.QuerySql+" LIMIT 1", query.QueryParams...).First(&counts)
	return
}

type FlowSet struct {
	Counts int    `json:"counts" gorm:"counts"`
	AppKey string `json:"app_key" gorm:"app_key"`
}

// AppFlowConfig 组合应用列表
func (f *FlowConfModel) AppFlowConfig(appKeys []string) (flows []*FlowSet) {
	f.Table(f.TableName()).Select("COUNT(*) AS counts", "app_key").Where("conf_type = ? and app_key in ?", 1, appKeys).Group("app_key").Find(&flows)
	return
}

func (f *FlowConfModel) GetFlowByAppKey(appKey string) (flows []*FlowList) {
	f.Raw("SELECT * FROM `"+f.TableName()+"` WHERE app_key = ?", appKey).Find(&flows)
	return
}

// DeleteFlowById 删除配置信息要删除相关联的配置
func (f *FlowConfModel) DeleteFlowById(id int64) bool {
	err := f.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM flow_conf WHERE id = ?", id).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM flow_ad_type_rel WHERE flow_id = ?", id).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM flow_ads_rel WHERE flow_id = ?", id).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM flow_app_rel WHERE flow_id = ?", id).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM flow_pos_rel WHERE flow_id = ?", id).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM flow_pos_policy_rel WHERE flow_id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false
	}
	return true
}
