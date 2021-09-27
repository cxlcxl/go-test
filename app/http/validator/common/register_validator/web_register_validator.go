package register_validator

import (
	"goskeleton/app/core/container"
	"goskeleton/app/http/validator/common/upload_files"
	"goskeleton/app/http/validator/common/websocket"
	"goskeleton/app/http/validator/web/flow"
	"goskeleton/app/http/validator/web/groups"
	"goskeleton/app/http/validator/web/users"
)

// WebRegisterValidator 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func WebRegisterValidator() {
	// 创建容器
	containers := container.CreateContainersFactory()
	// 验证器注册 map
	validators := map[string]interface{}{
		"UsersLogin":            users.Login{},
		"UsersLogout":           users.Logout{},
		"RefreshToken":          users.RefreshToken{},
		"UsersStore":            users.Store{},
		"UsersUpdate":           users.Update{},
		"UsersDestroy":          users.Destroy{},
		"UsersResetPass":        users.ResetPass{},
		"UsersChangeGroup":      users.ChangeGroup{},
		"UsersCheck":            users.Check{},
		"GroupStore":            groups.Store{},
		"GroupUpdate":           groups.Update{},
		"PermissionGroupUpdate": groups.PermissionUpdate{},
		"UploadFiles":           upload_files.UpFiles{}, // 文件上传
		"WebsocketConnect":      websocket.Connect{},
		"FlowDestroy":           flow.Destroy{},
		"FlowDetail":            flow.DetailFlow{},
	}

	for validator, forStruct := range validators {
		//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
		containers.Set(validator, forStruct)
	}
}
