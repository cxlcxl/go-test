package data

import (
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
)

func ReportThirdDataModel() *ReportThirdData {
	return &ReportThirdData{BaseModel: model.BaseModel{DB: model.CreateMysqlDB("Data")}}
}

type ReportThirdData struct {
	model.BaseModel `json:"-"`
	AppKey          string         `json:"app_key"`
	Platform        int            `json:"platform"`
	SspId           int            `json:"ssp_id"`
	IsMobgi         int            `json:"is_mobgi"`
	AdsId           string         `json:"ads_id"`
	AdType          int            `json:"ad_type"`
	PosKey          string         `json:"pos_key"`
	Days            tool.LocalDate `json:"days"`
	ThirdViews      int            `json:"third_views"`
	ThirdClicks     int            `json:"third_clicks"`
	AdIncome        float64        `json:"ad_income"`
	SettlementType  int            `json:"settlement_type"`
	Source          int            `json:"source"`
	Hours           int            `json:"hours"`
	IsCustom        int            `json:"is_custom"`
	DivisionRate    float64        `json:"division_rate"`
}

func (r *ReportThirdData) TableName() string {
	return "report_third_data"
}

func (r *ReportThirdData) Show(where *tool.WhereQuery) (data []ReportThirdData) {
	sql := "SELECT days,SUM(third_views) AS third_views,SUM(third_clicks) AS third_clicks,SUM(ad_income) AS ad_income,app_key,ads_id,ad_type " +
		"FROM report_third_data WHERE " + where.QuerySql + " GROUP BY days"
	r.Raw(sql, where.QueryParams...).Find(&data)
	return data
}

func (r *ReportThirdData) counts(where *tool.WhereQuery) (counts int) {
	sql := "SELECT COUNT(*) AS counts FROM report_third_data WHERE " + where.QuerySql
	r.Raw(sql, where.QueryParams...).First(&counts)
	return counts
}
