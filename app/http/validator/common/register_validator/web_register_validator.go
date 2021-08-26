package register_validator

import (
	"goskeleton/app/core/container"
	"goskeleton/app/http/validator/common/upload_files"
	"goskeleton/app/http/validator/common/websocket"
	"goskeleton/app/http/validator/web/news"
	"goskeleton/app/http/validator/web/users"
)

// WebRegisterValidator 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func WebRegisterValidator() {
	// 创建容器
	containers := container.CreateContainersFactory()
	// 验证器注册 map
	validators := map[string]interface{}{
		// Users 模块表单验证器按照 key => value 形式注册在容器，方便路由模块中调用
		"UsersLogin":   users.Login{},
		"UsersLogout":  users.Logout{},
		"RefreshToken": users.RefreshToken{},
		// Users基本操作（CURD）
		"UsersShow":    users.Show{},
		"UsersStore":   users.Store{},
		"UsersUpdate":  users.Update{},
		"UsersDestroy": users.Destroy{},
		// 文件上传
		"UploadFiles": upload_files.UpFiles{},
		// Websocket 连接验证器
		"WebsocketConnect": websocket.Connect{},
		// News 验证器
		"NewsStore": news.Store{},
	}

	for validator, forStruct := range validators {
		//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
		containers.Set(validator, forStruct)
	}
}
