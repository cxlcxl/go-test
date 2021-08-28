package roles

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	validatorWeb "goskeleton/app/http/validator/web"
	"goskeleton/app/utils/response"
)

type PermissionUpdate struct {
	PermissionIds []int64 `form:"permission_ids" json:"permission_ids" binding:"required"`
	RoleId        int64   `form:"role_id" json:"role_id" binding:"required"`
}

func (pr PermissionUpdate) CheckParams(c *gin.Context) {
	//1.基本的验证规则没有通过
	if err := c.ShouldBind(&pr); err != nil {
		response.ErrorParam(c, validatorWeb.Translate(err))
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(pr, "", c)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(c, "表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Roles{}).UpdatePermissions(extraAddBindDataContext)
	}
}
