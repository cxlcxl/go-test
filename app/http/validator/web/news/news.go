package news

type BaseField struct {
	Title string `form:"title" json:"title" binding:"required"`
}
