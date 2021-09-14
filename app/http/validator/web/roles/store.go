package roles

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type Store struct {
	RoleName string `form:"role_name" json:"role_name" binding:"required,min=1"`
	Des      string `form:"des" json:"des"`
}

func (s Store) CheckParams(c *gin.Context) {
	//1.基本的验证规则没有通过
	if err := c.ShouldBind(&s); err != nil {
		response.ErrorParam(c, validator.Translate(err))
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(s, "", c)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(c, "表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Roles{}).Store(extraAddBindDataContext)
	}
}
