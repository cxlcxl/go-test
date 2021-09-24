package groups

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type PermissionUpdate struct {
	PermissionIds []int64 `form:"permission_ids" json:"permission_ids" binding:"required"`
	RoleId        int64   `form:"role_id" json:"role_id" binding:"required"`
}

func (pr PermissionUpdate) CheckParams(c *gin.Context) {
	//1.基本的验证规则没有通过
	if err := c.ShouldBind(&pr); err != nil {
		response.ErrorParam(c, validator.Translate(err))
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(pr, "", c)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(c, consts.JsonMarshalFailed, "")
	} else {
		(&web.Group{}).UpdatePermissions(extraAddBindDataContext)
	}
}
