package model

type ReportThirdData struct {
	BaseModel `json:"-"`
}

func ReportThirdDataDb(sqlType string) *ReportThirdData {
	return &ReportThirdData{BaseModel{DB: UseDbConn(sqlType)}}
}

func (r *ReportThirdData) TableName() string {
	return "report_third_data"
}
