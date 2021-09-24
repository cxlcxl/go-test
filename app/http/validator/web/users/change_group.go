package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
	"log"
)

type ChangeGroup struct {
	UserId  []int64 `json:"user_id" form:"user_id" binding:"required"`
	GroupId int64   `json:"group_id" form:"group_id" binding:"required"`
}

func (c ChangeGroup) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&c); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}
	log.Println(c.GroupId, c.UserId)
	log.Printf("%#T", c.UserId)

	extraAddBindDataContext := data_transfer.DataAddContext(c, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
	} else {
		(&web.Users{}).ChangeGroup(extraAddBindDataContext)
	}

}
