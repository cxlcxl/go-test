package model

import (
	"fmt"
	"gorm.io/gorm"
	"goskeleton/app/global/variable"
	"strings"
	"time"
)

func CreateRoleFactory() *RoleModel {
	return &RoleModel{BaseModel: BaseModel{DB: UseDbConn("")}}
}

type RoleModel struct {
	BaseModel `json:"-"`
	RoleName  string `gorm:"column:role_name" json:"role_name"`
	Des       string `json:"des"`
	*BaseColumns
	//Permissions []RolePermission
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

// GetRoles ..
func (r *RoleModel) GetRoles(roleName string) []RoleModel {
	var roles []RoleModel
	if res := r.Raw("SELECT * FROM roles WHERE role_name LIKE ?", "%"+roleName+"%").Find(&roles); res.RowsAffected > 0 {
		return roles
	}
	return nil
}

// GetRoleById ..
func (r *RoleModel) GetRoleById(roleId int) (role RoleModel) {
	if res := r.Raw("SELECT * FROM roles WHERE id = ?", roleId).First(&role); res.Error == nil {
		return role
	}
	return
}

// Update ..
func (r *RoleModel) Update(values map[string]interface{}, roleId int) bool {
	if r.Where("id = ?", roleId).Updates(values).RowsAffected > 0 {
		return true
	}
	return false
}

// Destroy ..
func (r *RoleModel) Destroy(roleId int) bool {
	if r.Delete(r, roleId).RowsAffected > 0 {
		return true
	}
	return false
}

// RolePermission 角色权限对象
type RolePermission struct {
	RoleId       int64 `gorm:"column:role_id" json:"role_id"`
	PermissionId int64 `gorm:"column:permission_id" json:"permission_id"`
}

// UpdatePermissions 修改权限
func (r *RoleModel) UpdatePermissions(permissionIds []int, roleId int) bool {
	insertSQL := make([]string, len(permissionIds))
	for i, id := range permissionIds {
		insertSQL[i] = fmt.Sprintf("(%d,%d)", roleId, id)
	}
	err := r.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM role_permissions WHERE role_id = ?", roleId).Error; err != nil {
			return err
		}
		if len(insertSQL) > 0 {
			if err := tx.Exec("INSERT INTO role_permissions(role_id,permission_id) VALUES " + strings.Join(insertSQL, ",")).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}

// GetPermissions 获取角色的权限
func (r *RoleModel) GetPermissions(roleId int) (permissions []RolePermission) {
	r.Raw("SELECT permission_id FROM role_permissions WHERE role_id = ?", roleId).Find(&permissions)
	return
}
