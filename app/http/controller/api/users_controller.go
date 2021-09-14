package api

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/controller"
	"goskeleton/app/model"
	"goskeleton/app/service/users/curd"
	userstoken "goskeleton/app/service/users/token"
	"goskeleton/app/utils/md5_encrypt"
	"goskeleton/app/utils/response"
	"time"

	"github.com/gin-gonic/gin"
)

type Users struct{}

// Login 用户登录
func (u *Users) Login(context *gin.Context) {
	userName := context.GetString("user_name")
	pass := context.GetString("pass")
	userModelFact := model.CreateUserFactory("")
	userModel, errMsg := userModelFact.Login(userName, pass, context.ClientIP())
	if errMsg != "" {
		response.Fail(context, consts.CurdLoginFailCode, errMsg, "")
		return
	}

	if userModel != nil {
		userTokenFactory := userstoken.CreateUserFactory()
		expireAt := variable.ConfigYml.GetInt64("Token.JwtTokenCreatedExpireAt")
		if userToken, err := userTokenFactory.GenerateToken(userModel.Id, userModel.UserName, userModel.RealName, userModel.Email, userModel.Avatar, expireAt); err == nil {
			if userTokenFactory.RecordLoginToken(userToken, context.ClientIP()) {
				data := gin.H{
					"user_id":    userModel.Id,
					"user_name":  userName,
					"email":      userModel.Email,
					"real_name":  userModel.RealName,
					"token":      userToken,
					"updated_at": time.Now().Format(variable.DateFormat),
				}
				response.Success(context, consts.CurdStatusOkMsg, data)
				go model.CreateLoginLogFactory("").LogLogin(userModel.Id, context.ClientIP(), "api", 1)
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

// Update 用户更新
func (u *Users) Update(c *gin.Context) {
	values := controller.GetValues(c, []string{"id.float64", "user_name", "real_name", "phone", "email", "avatar", "remark"})
	if errMsg := model.CreateUserFactory("").UserIsExists(values["id"].(float64), values["user_name"].(string), values["email"].(string), values["phone"].(string)); errMsg != "" {
		response.Fail(c, consts.CurdRegisterFailCode, errMsg, "")
		return
	}
	if curd.CreateUserCurdFactory().Update(values) {
		response.Success(c, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, "")
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
		userId, ok := context.Get("user_id")
		if !ok {
			response.Fail(context, 4001, "用户信息读取失败", "")
			return
		}
		if ok = model.CreateUserFactory("").Logout(token.(string), float64(userId.(int64))); !ok {
			response.Fail(context, 4001, "TOKEN 清除失败，请重试", "")
		} else {
			response.Success(context, "退出成功", "")
		}
	}
}

// EmailVerify 邮件验证码
func (u *Users) EmailVerify(c *gin.Context) {

}

// Register 注册
func (u *Users) Register(c *gin.Context) {
	values := controller.GetValues(c, []string{"user_name", "pass", "email", "phone", "real_name"})
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

// ResetPass 修改密码
func (u *Users) ResetPass(c *gin.Context) {
	values := controller.GetValues(c, []string{"id.float64", "pass", "original_pass"})
	// 检查旧密码是否正确
	original := md5_encrypt.Base64Md5(values["original_pass"].(string))
	if !model.CreateUserFactory("").CheckPass(values["id"].(float64), original) {
		response.Fail(c, consts.CurdUpdateFailCode, "原密码不正确", "")
		return
	}
	pass := md5_encrypt.Base64Md5(values["pass"].(string))
	if model.CreateUserFactory("").ResetPass(values["id"].(float64), pass) {
		response.Success(c, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, "")
	}
}
