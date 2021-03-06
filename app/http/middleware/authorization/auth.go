package authorization

import (
	"goskeleton/app/global/consts"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	userstoken "goskeleton/app/service/users/token"
	"goskeleton/app/utils/response"
	"strings"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HeaderParams struct {
	Authorization string `header:"Authorization"`
}

// CheckTokenAuth 检查token权限
func CheckTokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerParams := HeaderParams{}
		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			variable.ZapLog.Error(my_errors.ErrorsValidatorBindParamsFail, zap.Error(err))
			context.Abort()
			return
		}
		if len(headerParams.Authorization) < 20 {
			response.ErrorTokenAuthFail(context)
			return
		}
		token := strings.Split(headerParams.Authorization, " ")
		if len(token) != 2 || len(token[1]) < 20 {
			response.ErrorTokenAuthFail(context)
			return
		}
		if tokenIsEffective := userstoken.CreateUserFactory().IsEffective(token[1]); !tokenIsEffective {
			response.ErrorTokenAuthFail(context)
			return
		}
		customerToken, err := userstoken.CreateUserFactory().ParseToken(token[1])
		if err != nil {
			response.ErrorTokenAuthFail(context)
			return
		}
		key := variable.ConfigYml.GetString("Token.BindContextKeyName")
		// token验证通过，同时绑定在请求上下文
		context.Set(key, customerToken)
		context.Set("token", token[1])
		context.Set("user_id", customerToken.UserId)
		context.Set("username", customerToken.Name)
		context.Set("email", customerToken.Email)
		context.Next()
	}
}

// CheckCasbinAuth casbin检查用户对应的角色权限是否允许访问接口
func CheckCasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUrl := c.Request.URL.Path // 路由例如 /admin/user/index
		method := c.Request.Method       // 方法 GET/POST/PUT...

		// 模拟请求参数转换后的角色（roleId=2）
		// 主线版本没有深度集成casbin的使用逻辑
		// GinSkeleton-Admin 系统则深度集成了casbin接口权限管控
		role := "2" // 这里模拟某个用户的roleId=2
		// 这里将用户的id解析为所拥有的的角色，判断是否具有某个权限即可
		isPass, err := variable.Enforcer.Enforce(role, requestUrl, method)
		if err != nil {
			response.ErrorCasbinAuthFail(c, err.Error())
			return
		} else if !isPass {
			response.ErrorCasbinAuthFail(c, "")
			return
		} else {
			c.Next()
		}
	}
}

// CheckCaptchaAuth 验证码中间件
func CheckCaptchaAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		captchaIdKey := variable.ConfigYml.GetString("Captcha.captchaId")
		captchaValueKey := variable.ConfigYml.GetString("Captcha.captchaValue")
		captchaId := c.PostForm(captchaIdKey)
		value := c.PostForm(captchaValueKey)
		if captchaId == "" || value == "" {
			response.Fail(c, consts.CaptchaCheckParamsInvalidCode, consts.CaptchaCheckParamsInvalidMsg, "")
			return
		}
		if captcha.VerifyString(captchaId, value) {
			c.Next()
		} else {
			response.Fail(c, consts.CaptchaCheckFailCode, consts.CaptchaCheckFailMsg, "")
			return
		}
	}
}
