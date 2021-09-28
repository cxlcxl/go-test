package data

import "goskeleton/app/model"

type ThirdOriginalDataModel struct {
	model.BaseModel
	Id             int64   `json:"id"`
	AdsGid         int     `json:"ads_gid"`                              // 广告商组ID
	AdsId          string  `json:"ads_id" gorm:"column:ads_key"`         // 广告商组标识
	IsMobgi        int     `json:"is_mobgi"`                             // 是否内部订单：0：外部，1：内部
	AppId          string  `json:"app_id"`                               // 应用ID
	AppName        string  `json:"app_name"`                             // 应用名称
	MediumId       string  `json:"medium_id"`                            // 媒体ID
	PlacementId    string  `json:"placement_id"`                         // 广告位ID
	PlacementName  string  `json:"placement_name"`                       // 广告位名称
	PlacementType  string  `json:"placement_type"`                       // 广告位类型
	UnitId         string  `json:"unit_id"`                              // 广告单元ID
	Days           string  `json:"days"`                                 // 按每日分
	Hours          int     `json:"hours"`                                // 按小时分
	Views          int     `json:"views"`                                // 展示
	Clicks         int     `json:"clicks"`                               // 点击数
	Revenue        float64 `json:"revenue"`                              // 广告收入
	Requests       int     `json:"requests"`                             // 请求
	Filled         int     `json:"filled"`                               // 填充次数
	Ecpm           float64 `json:"ecpm"`                                 // ECPM
	ClickRate      float64 `json:"click_rate"`                           //
	FillRate       float64 `json:"fill_rate"`                            //
	CurrencyType   int     `json:"currency_type"`                        // 币种 1美元 2人民币
	ExchangeRate   float64 `json:"exchange_rate"`                        // 汇率
	SettlementType int     `json:"settlement_type"`                      // 结算类型 1结算 2不结算
	Source         int     `json:"source"`                               // 数据来源：0系统统计 1手工导入
	UpdatedAt      string  `json:"updated_at" gorm:"column:update_time"` //
}

func ThirdOriginalDataDB() *ThirdOriginalDataModel {
	return &ThirdOriginalDataModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Data")}}
}

func (t *ThirdOriginalDataModel) TableName() string {
	return "report_third_original_data"
}
