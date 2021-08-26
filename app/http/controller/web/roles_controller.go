package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller"
	"goskeleton/app/model"
	"goskeleton/app/utils/response"
	"log"
	"strconv"
)

type Roles struct{}

// Show 角色列表
func (r *Roles) Show(c *gin.Context) {
	roleName := c.Query("role_name")
	showList := model.CreateRoleFactory("").GetRoles(roleName)
	if showList != nil {
		response.Success(c, consts.CurdStatusOkMsg, showList)
	} else {
		response.Fail(c, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	}
}

// Store 角色新增
func (r *Roles) Store(c *gin.Context) {
	roleName := c.GetString("role_name")
	des := c.GetString("des")
	if model.CreateRoleFactory("").Store(roleName, des) {
		response.Success(c, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(c, consts.CurdRegisterFailCode, consts.CurdRegisterFailMsg, "")
	}
}

// Update 用户更新
func (r *Roles) Update(c *gin.Context) {
	id := c.GetFloat64("id")
	values := controller.GetValues(c, []string{"role_name", "des"})
	if model.CreateRoleFactory("").Update(values, int(id)) {
		response.Success(c, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, "")
	}
}

// Destroy 删除记录
func (r *Roles) Destroy(c *gin.Context) {
	roleId, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Fail(c, consts.CurdDeleteFailCode, "参数错误", "")
		return
	}
	log.Println(roleId)
	if model.CreateRoleFactory("").Destroy(roleId) {
		response.Success(c, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(c, consts.CurdDeleteFailCode, consts.CurdDeleteFailMsg, "")
	}
}
