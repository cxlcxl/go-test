package users

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"

	"github.com/gin-gonic/gin"
)

type Login struct {
	BaseField
}

func (l Login) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&l); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(l, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
	} else {
		(&web.Users{}).Login(extraAddBindDataContext)
	}

}
