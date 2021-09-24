package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/model"
	"goskeleton/app/utils/response"
)

type Group struct{}

// Show 用户组列表
func (g *Group) Show(c *gin.Context) {
	name := c.Query("name")
	showList := model.UserGroupDB().GetGroup(name)
	if showList != nil {
		response.Success(c, consts.CurdStatusOkMsg, showList)
	} else {
		response.Fail(c, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	}
}

// Store 角色新增
func (g *Group) Store(c *gin.Context) {
	//roleName := c.GetString("role_name")
	//des := c.GetString("des")
	//if model.UserGroupDB().Store(roleName, des) {
	//	response.Success(c, consts.CurdStatusOkMsg, "")
	//} else {
	//	response.Fail(c, consts.CurdRegisterFailCode, consts.CurdRegisterFailMsg, "")
	//}
}

// Update 用户更新
func (g *Group) Update(c *gin.Context) {
	//id := c.GetFloat64("id")
	//values := controller.GetValues(c, []string{"role_name", "des"})
	//if model.UserGroupDB().Update(values, int(id)) {
	//	response.Success(c, consts.CurdStatusOkMsg, "")
	//} else {
	//	response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, "")
	//}
}

// Destroy 删除记录
func (g *Group) Destroy(c *gin.Context) {
	//roleId, err := strconv.Atoi(c.PostForm("id"))
	//if err != nil {
	//	response.Fail(c, consts.CurdDeleteFailCode, "参数错误", "")
	//	return
	//}
	//if model.UserGroupDB().Destroy(roleId) {
	//	response.Success(c, consts.CurdStatusOkMsg, "")
	//} else {
	//	response.Fail(c, consts.CurdDeleteFailCode, consts.CurdDeleteFailMsg, "")
	//}
}

// UpdatePermissions 更新角色权限
func (g *Group) UpdatePermissions(c *gin.Context) {
	//roleId := c.GetFloat64("role_id")
	//permissions, ok := c.GetPostFormArray("permission_ids")
	//if !ok {
	//	response.Fail(c, consts.CurdDeleteFailCode, "权限参数获取失败", "")
	//}
	//var permissionIds = make([]int, len(permissions))
	//for i, permissionId := range permissions {
	//	id, err := strconv.Atoi(permissionId)
	//	if err != nil {
	//		continue // 转化不成功的丢弃
	//	}
	//	permissionIds[i] = id
	//}
	//if model.UserGroupDB().UpdatePermissions(permissionIds, int(roleId)) {
	//	response.Success(c, consts.CurdStatusOkMsg, "")
	//} else {
	//	response.Fail(c, consts.CurdDeleteFailCode, consts.CurdUpdateFailMsg, "")
	//}
}
