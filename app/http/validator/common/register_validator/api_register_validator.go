package register_validator

import (
	"goskeleton/app/core/container"
	"goskeleton/app/http/validator/api/home"
	apiUsers "goskeleton/app/http/validator/api/users"
)

// ApiRegisterValidator 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func ApiRegisterValidator() {
	// 创建容器
	containers := container.CreateContainersFactory()
	// 验证器注册 map
	validators := map[string]interface{}{
		"HomeNews":     home.News{},
		"ApiUserLogin": apiUsers.Login{},
	}

	for validator, forStruct := range validators {
		//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
		containers.Set(validator, forStruct)
	}
}
