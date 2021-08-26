package data_type

type Page struct {
	Page  float64 `form:"page" json:"page" binding:"min=1"`
	Limit float64 `form:"limit" json:"limit" binding:"min=1"`
}
