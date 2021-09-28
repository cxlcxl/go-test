package data

import (
	"goskeleton/app/model"
)

type FinanceReportDataModel struct {
	model.BaseModel
	Id           int64   `json:"id"`
	YearMonths   string  `json:"year_months"`   // 年月
	AppKey       string  `json:"app_key"`       // 应用app_key
	AdsId        string  `json:"ads_id"`        // 广告商的标识
	SspId        int64   `json:"ssp_id"`        // 开发者ID
	ExchangeRate float64 `json:"exchange_rate"` // 汇率
	ContractRate float64 `json:"contract_rate"` // 对账比例
	AdIncome     float64 `json:"ad_income"`     // 广告收入
	Status       int     `json:"status"`        // 状态 1待审核 3待对账 4对账中 5待结算 6已结算
	UpdatedAt    string  `json:"updated_at" gorm:"column:update_time"`
}

func FinanceReportDataDB() *FinanceReportDataModel {
	return &FinanceReportDataModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Data")}}
}

func (e *FinanceReportDataModel) TableName() string {
	return "finance_report_data"
}
