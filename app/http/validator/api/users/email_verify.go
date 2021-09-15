package api_users

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/api"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"

	"github.com/gin-gonic/gin"
)

type EmailVerify struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

func (e EmailVerify) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&e); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(e, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
	} else {
		(&api.Users{}).EmailVerify(extraAddBindDataContext)
	}
}
