package data

import (
	"goskeleton/app/model"
)

type FinanceStatementModel struct {
	model.BaseModel
	Id           int64   `json:"id"`
	YearMonths   string  `json:"year_months"`   // 年月
	SspId        int64   `json:"ssp_id"`        // 开发者ID
	ContractRate float64 `json:"contract_rate"` // 对账比例
	AdIncome     float64 `json:"ad_income"`     // 广告收入
	InvoiceType  int     `json:"invoice_type"`  // 发票种类：0无；1增值税普通发票；2税率3%增值税专用发票；3税率6%增值税专用发票
	Status       int     `json:"status"`        // 状态 1未确认 3待对账 4对账中 5待结算 6已结算
	Operator     string  `json:"operator"`      // 操作人
	UpdatedAt    string  `json:"updated_at" gorm:"column:update_time"`
}

func FinanceStatementDB() *FinanceStatementModel {
	return &FinanceStatementModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Data")}}
}

func (e *FinanceStatementModel) TableName() string {
	return "finance_statement"
}
