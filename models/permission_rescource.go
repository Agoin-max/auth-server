package models

type PermissionResource struct {
	ID           uint `gorm:"primary_key" json:"id"`
	PermissionId uint `json:"permission_id"` // 权限ID
	ResourceId   uint `json:"resource_id"`   // 资源ID
}

func (PermissionResource) TableName() string {
	return "sys_permission_resource"
}
