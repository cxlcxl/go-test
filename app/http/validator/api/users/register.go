package api_users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/api"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
	"regexp"
)

type Store struct {
	Username string `form:"user_name" json:"user_name" binding:"required,max=35,min=2"`
	Phone    string `form:"phone" json:"phone"`
	Email    string `form:"email" json:"email" binding:"required"`
	RealName string `form:"real_name" json:"real_name" binding:"max=30"`
	Remark   string `form:"remark" json:"remark"`
	Password string `form:"pass" json:"pass" binding:"required,min=6,max=20"`
	Verify   string `form:"verify" json:"verify" binding:"required,min=4,max=6"`
}

func (s Store) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&s); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}

	if errMsg := s.customCheck(); errMsg != "" {
		response.ErrorParam(context, errMsg)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(s, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
	} else {
		(&api.Users{}).Register(extraAddBindDataContext)
	}
}

func (s Store) customCheck() string {
	if len(s.Phone) > 0 {
		if ok, _ := regexp.MatchString(consts.PhoneRule, s.Phone); !ok {
			return consts.PhoneError
		}
	}
	if len(s.Email) > 0 {
		if ok, _ := regexp.MatchString(consts.EmailRule, s.Email); !ok {
			return consts.EmailError
		}
	}
	return ""
}
