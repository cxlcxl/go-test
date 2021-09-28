package data

import (
	"goskeleton/app/model"
)

type FinanceStatementInfoModel struct {
	model.BaseModel
	Id                 int64   `json:"id"`
	StartTime          string  `json:"start_time"`           // 结束年月
	EndTime            string  `json:"end_time"`             // 结束年月
	SspId              int64   `json:"ssp_id"`               // 开发者ID
	ContractRate       float64 `json:"contract_rate"`        // 对账比例
	AdIncome           float64 `json:"ad_income"`            // 广告收入
	InvoiceType        int     `json:"invoice_type"`         // 发票种类：0无；1增值税普通发票；2税率3%增值税专用发票；3税率6%增值税专用发票
	Status             int     `json:"status"`               // 状态 1待审核 3待对账 4对账中 5待结算 6已结算
	Operator           string  `json:"operator"`             // 操作人
	FinanceStatementId int64   `json:"finance_statement_id"` // 月度明细表ID，多个 “,” 分隔
	model.TimeColumns
}

func FinanceStatementInfoDB() *FinanceStatementInfoModel {
	return &FinanceStatementInfoModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Data")}}
}

func (e *FinanceStatementInfoModel) TableName() string {
	return "finance_statement_info"
}
