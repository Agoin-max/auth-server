package models

type Permission struct {
	Base
	PermissionName string `json:"permission_name"` // 权限名称
	PermissionDesc string `json:"permission_desc"` // 权限描述
	AppId          int    `json:"app_id"`          // 应用ID
	Status         int    `json:"status"`          // 状态
}

func (Permission) TableName() string {
	return "sys_permission"
}
