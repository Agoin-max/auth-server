package models

type Resource struct {
	Base
	ResourceName string `json:"resource_name"` // 资源名称
	ResourceCode string `json:"resource_code"` // 资源编码
	AppId        int    `json:"app_id"`        // 应用ID
}

func (Resource) TableName() string {
	return "sys_resource"
}
