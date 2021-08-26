package web

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/model"
	"goskeleton/app/utils/response"

	"github.com/gin-gonic/gin"
)

type News struct{}

// Store 创建内容
func (n *News) Store(context *gin.Context) {
	userId := context.GetInt("user_id")
	title := context.GetString("title")
	des := context.GetString("des")
	content := context.GetString("content")
	if model.CreateNewFactory("").Store(title, des, content, userId) {
		response.Success(context, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(context, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, "")
	}
}

// Test
func (n *News) Test(context *gin.Context) {
	response.Success(context, "调用成功", "")
}
