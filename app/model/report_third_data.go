package model

import (
	"goskeleton/app/model/tool"
)

func ReportThirdDataDb() *ReportThirdData {
	return &ReportThirdData{BaseModel: BaseModel{DB: CreateMysqlDB("Data")}}
}

type ReportThirdData struct {
	BaseModel      `json:"-"`
	AppKey         string         `json:"app_key"`
	Platform       int            `json:"platform"`
	SspId          int            `json:"ssp_id"`
	IsMobgi        int            `json:"is_mobgi"`
	AdsId          string         `json:"ads_id"`
	AdType         int            `json:"ad_type"`
	PosKey         string         `json:"pos_key"`
	Days           tool.LocalDate `json:"days"`
	ThirdViews     int            `json:"third_views"`
	ThirdClicks    int            `json:"third_clicks"`
	AdIncome       float64        `json:"ad_income"`
	SettlementType int            `json:"settlement_type"`
	Source         int            `json:"source"`
	Hours          int            `json:"hours"`
	IsCustom       int            `json:"is_custom"`
	DivisionRate   float64        `json:"division_rate"`
}

func (r *ReportThirdData) TableName() string {
	return "report_third_data"
}

func (r *ReportThirdData) Show(where *tool.WhereQuery, limitStart, limit int) (counts int, data []ReportThirdData) {
	if counts = r.counts(where); counts > 0 {
		sql := "SELECT * FROM report_third_data WHERE " + where.QuerySql + " LIMIT ?,?"
		where.QueryParams = append(where.QueryParams, limitStart, limit)
		r.Raw(sql, where.QueryParams...).Find(&data)
		return counts, data
	} else {
		return 0, nil
	}
}

func (r *ReportThirdData) counts(where *tool.WhereQuery) (counts int) {
	sql := "SELECT COUNT(*) AS counts FROM report_third_data WHERE " + where.QuerySql
	r.Raw(sql, where.QueryParams...).First(&counts)
	return counts
}
