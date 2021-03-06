package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type Destroy struct {
	// 表单参数验证结构体支持匿名结构体嵌套、以及匿名结构体与普通字段组合
	Id
}

// 验证器语法，参见 Register.go文件，有详细说明

func (d Destroy) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&d); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(d, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserShow表单参数验证器json化失败", "")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Destroy(extraAddBindDataContext)
	}
}
