package api_users

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/api"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
	"regexp"

	"github.com/gin-gonic/gin"
)

type ResetPass struct {
	Id           int    `form:"id" json:"id" binding:"required"`
	Pass         string `form:"pass" json:"pass" binding:"min=6,max=20"`
	OriginalPass string `form:"original_pass" json:"original_pass" binding:"min=6,max=20"`
}

func (r ResetPass) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&r); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}
	if ok, _ := regexp.MatchString(consts.PassRule, r.Pass); !ok {
		response.ErrorSystem(context, consts.PassError, "")
		return
	}
	if ok, _ := regexp.MatchString(consts.PassRule, r.OriginalPass); !ok {
		response.ErrorSystem(context, consts.PassError, "")
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(r, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
	} else {
		(&api.Users{}).ResetPass(extraAddBindDataContext)
	}
}
