package model

import (
	"gorm.io/gorm"
	"goskeleton/app/global/variable"
	"time"
)

func CreateRoleFactory(sqlType string) *RoleModel {
	return &RoleModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type RoleModel struct {
	BaseModel `json:"-"`
	RoleName  string `gorm:"column:role_name" json:"role_name"`
	Des       string `json:"des"`
	*BaseColumns
}

func (r *RoleModel) TableName() string {
	return "roles"
}

// Store 新增
func (r *RoleModel) Store(roleName, des string) bool {
	sql := "INSERT INTO roles(role_name,des,created_at) VALUES (?,?,?)"
	if r.Exec(sql, roleName, des, time.Now().Format(variable.DateFormat)).RowsAffected > 0 {
		return true
	}
	return false
}

func (r *RoleModel) GetRoles(roleName string) []RoleModel {
	var roles []RoleModel
	if res := r.Raw("SELECT * FROM roles WHERE role_name LIKE ?", "%"+roleName+"%").Find(&roles); res.RowsAffected > 0 {
		return roles
	}
	return nil
}

func (r *RoleModel) GetRoleById(roleId int) (role RoleModel) {
	if res := r.Raw("SELECT * FROM roles WHERE id = ?", roleId).First(&role); res.Error == nil {
		return role
	}
	return
}

func (r *RoleModel) Update(values map[string]interface{}, roleId int) bool {
	if r.Where("id = ?", roleId).Updates(values).RowsAffected > 0 {
		return true
	}
	return false
}

func (r *RoleModel) Destroy(roleId int) bool {
	if r.Delete(r, roleId).RowsAffected > 0 {
		return true
	}
	return false
}

type RolePermission struct {
	RoleId       int64 `gorm:"column:role_id" json:"role_id"`
	PermissionId int64 `gorm:"column:permission_id" json:"permission_id"`
}

// UpdatePermissions 修改权限
func (r *RoleModel) UpdatePermissions(permissionIds []int64, roleId int64) bool {
	var permissions = make([]RolePermission, len(permissionIds))
	for i, id := range permissionIds {
		permissions[i].PermissionId = id
		permissions[i].RoleId = roleId
	}
	err := r.Transaction(func(tx *gorm.DB) error {
		if err := tx.Raw("DELETE FROM role_permissions WHERE role_id = ?", roleId).Error; err != nil {
			return err
		}
		if err := tx.Create(&permissions).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}

func (r *RoleModel) GetPermissions(roleId int) (permissions []RolePermission) {
	r.Raw("SELECT permission_id FROM role_permissions WHERE role_id = ?", roleId).Find(&permissions)
	return
}
