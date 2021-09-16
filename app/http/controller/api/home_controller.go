package api

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller"
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
	"goskeleton/app/utils/response"
	"strconv"
)

type Home struct{}

// News ...
func (h *Home) News(c *gin.Context) {
	pageStart, limit := controller.GetPage(c)
	title := c.Query("title")
	where := map[string]interface{}{
		"state": 1,
		"title": []string{"like", title},
	}
	query := (&tool.WhereQuery{Filter: true}).GenerateWhere(where)
	total, news := model.NewsDB().SearchNews(query, pageStart, limit)
	response.Success(c, "OK", gin.H{
		"total": total,
		"list":  news,
	})
}

func (h *Home) Detail(c *gin.Context) {
	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		response.Fail(c, 400, "参数错误", "")
		return
	}

	if newInfo, ok := model.NewsDB().Show(newId); !ok {
		response.Fail(c, 400, "请求的内容不存在", "")
		return
	} else {
		response.Success(c, "OK", newInfo)
	}
}
