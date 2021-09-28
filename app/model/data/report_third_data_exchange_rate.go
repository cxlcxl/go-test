package data

import "goskeleton/app/model"

type ThirdDataExchangeRateModel struct {
	model.BaseModel
	Id           int     `json:"id"`
	Year         int     `json:"year"`
	Month        int     `json:"month"`
	ExchangeRate float64 `json:"exchange_rate"` // 汇率
}

func ThirdDataExchangeRateDB() *ThirdDataExchangeRateModel {
	return &ThirdDataExchangeRateModel{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Data")}}
}

func (e *ThirdDataExchangeRateModel) TableName() string {
	return "report_third_data_exchange_rate"
}
