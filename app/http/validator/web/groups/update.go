package groups

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type Update struct {
	GroupId int64  `form:"group_id" json:"group_id"  binding:"required,min=1"`
	Name    string `form:"name" json:"name"  binding:"required,min=1"`
	Desc    string `form:"desc" json:"desc"`
}

func (u Update) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&u); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(u, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
	} else {
		(&web.Group{}).Update(extraAddBindDataContext)
	}
}
