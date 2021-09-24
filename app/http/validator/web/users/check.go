package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type Check struct {
	UserId       int64  `json:"user_id" form:"user_id" binding:"required"`
	IsCheck      int    `json:"is_check" form:"is_check" binding:"required"`
	MobgiAccount int    `json:"mobgi_account" form:"mobgi_account" binding:"required"`
	CheckMsg     string `json:"check_msg" form:"check_msg" binding:"required"`
}

func (c Check) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&c); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(c, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
	} else {
		(&web.Users{}).CheckUser(extraAddBindDataContext)
	}

}
