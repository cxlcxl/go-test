package curd

import (
	"goskeleton/app/model"
	"goskeleton/app/utils/md5_encrypt"
)

func CreateUserCurdFactory() *UsersCurd {
	return &UsersCurd{}
}

type UsersCurd struct {
}

func (u *UsersCurd) Store(values map[string]interface{}) bool {
	values["pass"] = md5_encrypt.MobgiPwd(values["pass"].(string)) // 预先处理密码加密，然后存储在数据库
	return model.CreateUserFactory("").Store(values)
}

func (u *UsersCurd) Update(values map[string]interface{}) bool {
	// 预先处理密码加密等操作，然后进行更新
	if len(values["pass"].(string)) > 0 {
		// 预先处理密码加密，然后存储在数据库
		values["pass"] = md5_encrypt.MobgiPwd(values["pass"].(string))
	}

	return model.CreateUserFactory("").Update(values)
}
