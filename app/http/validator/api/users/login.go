package api_users

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/api"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `form:"user_name" json:"user_name" binding:"min=2"`
	Password string `form:"pass" json:"pass" binding:"min=6,max=20"`
}

func (l Login) CheckParams(context *gin.Context) {
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&l); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(l, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&api.Users{}).Login(extraAddBindDataContext)
	}
}
