package models

type Role struct {
	Base
	RoleName string `json:"role_name"` // 角色名称
	RoleCode string `json:"role_code"` // 角色编码
	RoleDesc string `json:"role_desc"` // 角色描述
	Status   int    `json:"status"`    // 状态
	AppId    int    `json:"app_id"`    // 应用ID
}

func (Role) TableName() string {
	return "sys_role"
}
