package data

import "goskeleton/app/model"

type ThirdDetailDataModel struct {
	model.BaseModel
	Id                    int64   `json:"id"`
	AppKey                string  `json:"app_key"`                 // 应用app_key
	Platform              int     `json:"platform"`                // 平台0安卓1IOS
	SspId                 int64   `json:"ssp_id"`                  // 开发者ID
	IsMobgi               int     `json:"is_mobgi"`                // 是否内部订单：0：外部，1：内部
	AdsGid                int     `json:"ads_gid"`                 // 广告商组ID
	AdsId                 string  `json:"ads_id"`                  // 广告商ID
	AdType                string  `json:"ad_type"`                 // 聚合类型 1视频聚合，2插图聚合
	PosKey                string  `json:"pos_key"`                 // 广告位
	ThirdBlockId          string  `json:"third_block_id"`          // 第三方广告位ID
	ChannelGid            int     `json:"channel_gid"`             // 渠道组
	AppVersion            string  `json:"app_version"`             //
	SdkVersion            string  `json:"sdk_version"`             //
	SysVersion            string  `json:"sys_version"`             //
	YearMonths            string  `json:"year_months"`             // 年月
	Days                  string  `json:"days"`                    // 按每日分
	Hours                 int     `json:"hours"`                   // 按小时分
	IsCustom              int     `json:"is_custom"`               //
	ThirdViews            int     `json:"third_views"`             // 展示
	ThirdClicks           int     `json:"third_clicks"`            // 点击数
	Revenue               float64 `json:"revenue"`                 // 广告收入
	AdIncomeAdjust        float64 `json:"ad_income_adjust"`        //
	DivisionRate          float64 `json:"division_rate"`           //
	DeveloperDivisionRate float64 `json:"developer_division_rate"` //
	SettlementType        int     `json:"settlement_type"`         // 结算类型 1结算 2不结算
	Source                int     `json:"source"`                  // 数据来源：0系统统计 1手工导入
	UpdatedAt             string  `json:"update_time" gorm:"column:update_time"`
}

func ThirdDetailDataDB() *ThirdDetailDataModel {
	return &ThirdDetailDataModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Data")}}
}

func (t *ThirdDetailDataModel) TableName() string {
	return "report_third_detail_data"
}
