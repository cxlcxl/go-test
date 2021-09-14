package routers

import (
	"goskeleton/app/global/variable"
	"goskeleton/app/http/controller/api"
	"goskeleton/app/http/middleware/authorization"
	"goskeleton/app/http/middleware/cors"
	validatorFactory "goskeleton/app/http/validator/core/factory"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// 该路由主要设置门户类网站等前台路由

func InitApiRouter() *gin.Engine {
	var router *gin.Engine
	// 非调试模式（生产模式） 日志写到日志文件
	if variable.ConfigYml.GetBool("AppDebug") == false {
		//1.将日志写入日志文件
		gin.DisableConsoleColor()
		f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		gin.DefaultWriter = io.MultiWriter(f)
		// 2.如果是有nginx前置做代理，基本不需要gin框架记录访问日志，开启下面一行代码，屏蔽上面的三行代码，性能提升 5%
		//gin.SetMode(gin.ReleaseMode)

		router = gin.Default()
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}

	//根据配置进行设置跨域
	if variable.ConfigYml.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Api 模块接口 hello word！")
	})

	//  创建一个门户类接口路由组
	vApi := router.Group("/api/v1/")
	{
		vApi.GET("news", validatorFactory.Create("HomeNews"))

		users := vApi.Group("user/")
		{
			users.POST("login", validatorFactory.Create("ApiUserLogin"))
			users.POST("email-verify", (&api.Users{}).EmailVerify)

			auth := users.Use(authorization.CheckTokenAuth())
			{
				auth.GET("info", (&api.Users{}).UserInfo)
				auth.POST("logout", (&api.Users{}).Logout)
			}
		}
	}
	return router
}
