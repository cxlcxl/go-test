package web

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/controller"
	"goskeleton/app/model"
	"goskeleton/app/service/users/curd"
	userstoken "goskeleton/app/service/users/token"
	"goskeleton/app/utils/response"
	"time"

	"github.com/gin-gonic/gin"
)

type Users struct{}

// Login 用户登录
func (u *Users) Login(context *gin.Context) {
	userName := context.GetString("user_name")
	pass := context.GetString("pass")
	phone := context.GetString("phone")
	userModelFact := model.CreateUserFactory("")
	userModel, errMsg := userModelFact.Login(userName, pass, context.ClientIP())
	if errMsg != "" {
		response.Fail(context, consts.CurdLoginFailCode, errMsg, "")
		return
	}

	if userModel != nil {
		userTokenFactory := userstoken.CreateUserFactory()
		expireAt := variable.ConfigYml.GetInt64("Token.JwtTokenCreatedExpireAt")
		if userToken, err := userTokenFactory.GenerateToken(userModel.Id, userModel.UserName, userModel.RealName, userModel.Phone, userModel.Avatar, expireAt); err == nil {
			if userTokenFactory.RecordLoginToken(userToken, context.ClientIP()) {
				data := gin.H{
					"user_id":    userModel.Id,
					"user_name":  userName,
					"real_name":  userModel.RealName,
					"phone":      phone,
					"token":      userToken,
					"updated_at": time.Now().Format(variable.DateFormat),
				}
				response.Success(context, consts.CurdStatusOkMsg, data)
				go model.CreateLoginLogFactory("").LogLogin(userModel.Id, context.ClientIP(), "web", 1)
				return
			}
		}
	}
	response.Fail(context, 2001, "请求失败", "")
}

// RefreshToken 刷新用户token
func (u *Users) RefreshToken(context *gin.Context) {
	oldToken := context.GetString("token")
	if newToken, ok := userstoken.CreateUserFactory().RefreshToken(oldToken, context.ClientIP(), "web"); ok {
		res := gin.H{
			"token": newToken,
		}
		response.Success(context, consts.CurdStatusOkMsg, res)
	} else {
		response.Fail(context, consts.CurdRefreshTokenFailCode, consts.CurdRefreshTokenFailMsg, "")
	}
}

// Show 用户查询
func (u *Users) Show(c *gin.Context) {
	values := controller.GetValues(c, []string{"user_name"})
	limitStart, limit := controller.GetPage(c)
	counts, showList := model.CreateUserFactory("").Show(values, limitStart, limit)
	if counts > 0 && showList != nil {
		roles := model.CreateRoleFactory().GetRoles("")
		for _, user := range showList {
			for _, role := range roles {
				if user.RoleId == role.Id {
					user.Role.RoleName = role.RoleName
				}
			}
		}
		response.Success(c, consts.CurdStatusOkMsg, gin.H{"counts": counts, "list": showList})
	} else {
		response.Fail(c, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	}
}

// Store 用户注册/新增
func (u *Users) Store(c *gin.Context) {
	//  由于本项目骨架已经将表单验证器的字段(成员)绑定在上下文，因此可以按照 GetString()、GetInt64()、GetFloat64（）等快捷获取需要的数据类型，注意：相关键名规则：  前缀+验证器结构体中的 json 标签
	// 当然也可以通过gin框架的上下文原始方法获取，例如： c.PostForm("user_name") 获取，这样获取的数据格式为文本，需要自己继续转换
	values := controller.GetValues(c, []string{"user_name", "pass", "email", "phone", "real_name", "role_id.int"})
	if errMsg := model.CreateUserFactory("").UserIsExists(0, values["user_name"].(string), values["email"].(string), values["phone"].(string)); errMsg != "" {
		response.Fail(c, consts.CurdRegisterFailCode, errMsg, "")
		return
	}
	if curd.CreateUserCurdFactory().Store(values) {
		response.Success(c, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(c, consts.CurdRegisterFailCode, consts.CurdRegisterFailMsg, "")
	}
}

// Update 用户更新
func (u *Users) Update(c *gin.Context) {
	values := controller.GetValues(c, []string{"id.float64", "user_name", "pass", "real_name", "phone", "email", "avatar", "remark"})
	if errMsg := model.CreateUserFactory("").UserIsExists(values["id"].(float64), values["user_name"].(string), values["email"].(string), values["phone"].(string)); errMsg != "" {
		response.Fail(c, consts.CurdRegisterFailCode, errMsg, "")
		return
	}
	//注意：这里没有实现权限控制逻辑，例如：超级管理管理员可以更新全部用户数据，普通用户只能修改自己的数据。目前只是验证了token有效、合法之后就可以进行后续操作
	// 实际使用请根据真是业务实现权限控制逻辑、再进行数据库操作
	if curd.CreateUserCurdFactory().Update(values) {
		response.Success(c, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, "")
	}

}

// Destroy 删除记录
func (u *Users) Destroy(context *gin.Context) {
	userId := context.GetFloat64("id")
	if model.CreateUserFactory("").Destroy(userId) {
		response.Success(context, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(context, consts.CurdDeleteFailCode, consts.CurdDeleteFailMsg, "")
	}
}

// UserInfo 获取当前登陆的用户信息
func (u *Users) UserInfo(context *gin.Context) {
	if user, ok := context.Get(variable.ConfigYml.GetString("Token.BindContextKeyName")); !ok {
		response.Fail(context, 2001, "获取失败", "")
	} else {
		response.Success(context, "请求成功", user)
	}
}

// Logout 退出登录
func (u *Users) Logout(context *gin.Context) {
	if token, ok := context.Get("token"); !ok {
		response.Fail(context, 2001, "TOKEN 获取失败，请重试", "")
	} else {
		if ok := model.CreateUserFactory("").Logout(token.(string), context.GetFloat64("id")); !ok {
			response.Fail(context, 4001, "TOKEN 清除失败，请重试", "")
		} else {
			response.Success(context, "退出成功", "")
		}
	}
}
