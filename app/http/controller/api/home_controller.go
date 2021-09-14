package api

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller"
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
	"goskeleton/app/utils/response"
)

type Home struct{}

// News 门户类首页新闻
func (h *Home) News(c *gin.Context) {
	pageStart, limit := controller.GetPage(c)
	title := c.Query("title")
	where := map[string]interface{}{
		"state": 1,
		"title": []string{"like", title},
	}
	query := (&tool.WhereQuery{Filter: true}).GenerateWhere(where)
	total, news := model.NewsDB().SearchNews(query, pageStart, limit)
	// 这里随便模拟一条数据返回
	response.Success(c, "OK", gin.H{
		"total": total,
		"list":  news,
	})
}
