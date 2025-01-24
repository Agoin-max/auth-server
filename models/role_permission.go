package models

type RolePermission struct {
	ID           uint `gorm:"primary_key" json:"id"`
	RoleID       int  `json:"role_id"`       // 角色ID
	PermissionID int  `json:"permission_id"` // 权限ID
}

func (RolePermission) TableName() string {
	return "sys_role_permission"
}
