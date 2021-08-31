package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller"
	"goskeleton/app/model"
	"goskeleton/app/model/tool"
	"goskeleton/app/utils/response"
	"strconv"
)

type Conf struct{}

// Show 列表展示
func (cf *Conf) Show(c *gin.Context) {
	pageStart, limit := controller.GetPage(c)
	values := controller.GetQueries(c, []string{"keyword", "in_use"})
	params := map[string]interface{}{
		"or": map[string]interface{}{
			"key":   []string{"like", values["keyword"].(string)},
			"name":  []string{"like", values["keyword"].(string)},
			"value": []string{"like", values["keyword"].(string)},
		},
		"in_use": values["in_use"],
	}
	query := (&tool.WhereQuery{}).GenerateWhere(params)
	counts, configs := model.ConfigDB().Show(query, pageStart, limit, "", "")
	response.Success(c, "请求成功", gin.H{
		"counts": counts,
		"list":   configs,
	})
}

// Destroy 删除
func (cf *Conf) Destroy(c *gin.Context) {
	values := controller.GetForms(c, []string{"id.int"})
	if ok := model.ConfigDB().Destroy(values["id"].(int)); !ok {
		response.Success(c, "删除成功", "")
	} else {
		response.Fail(c, consts.CurdDeleteFailCode, consts.CurdDeleteFailMsg, "")
	}
}

// Store 创建
func (cf *Conf) Store(c *gin.Context) {
	values := controller.GetForms(c, []string{"key", "name", "value", "des", "in_use.int", "bak0", "bak1", "bak2"})
	if ok := model.ConfigDB().Store(values); ok {
		response.Success(c, "创建成功", "")
	} else {
		response.Fail(c, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, "")
	}
}

func (cf *Conf) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Fail(c, consts.CurdUpdateFailCode, "请求失败，请重试", "")
		return
	}
	values := controller.GetForms(c, []string{"key", "name", "value", "des", "in_use.int", "bak0", "bak1", "bak2"})
	if ok := model.ConfigDB().Update(values, id); !ok {
		response.Success(c, "修改成功", "")
	} else {
		response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, "")
	}
}

func (cf *Conf) ConfInfo(c *gin.Context) {
	values := controller.GetQueries(c, []string{"id.int"})
	config := model.ConfigDB().Info(values["id"].(int))
	response.Success(c, "请求成功", config)
}
