package api_users

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/api"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
)

type ForgotReset struct {
	Email  string `form:"email" json:"email" binding:"required,email"`
	Pass   string `form:"pass" json:"pass" binding:"min=6,max=20"`
	Verify string `form:"verify" json:"verify" binding:"min=4,max=6"`
}

func (f ForgotReset) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&f); err != nil {
		log.Println(err)
		response.ErrorParam(context, validator.Translate(err))
		return
	}
	if ok, _ := regexp.MatchString(consts.PassRule, f.Pass); !ok {
		response.ErrorParam(context, consts.PassError)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(f, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
	} else {
		(&api.Users{}).ForgotPass(extraAddBindDataContext)
	}
}
