package model

type ReportThirdData struct {
	BaseModel `json:"-"`
}

func ReportThirdDataDb() *ReportThirdData {
	return &ReportThirdData{BaseModel{DB: UseDbConn("")}}
}

func (r *ReportThirdData) TableName() string {
	return "report_third_data"
}
