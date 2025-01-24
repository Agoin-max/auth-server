package models

type App struct {
	Base
	AppName     string `json:"app_name"`     // 应用名称
	AppCode     string `json:"app_code"`     // 应用编码
	AppIcon     string `json:"app_icon"`     // 应用图标
	Status      int    `json:"status"`       // 状态
	Description string `json:"description"`  // 描述
	SwaggerUrl  string `json:"swagger_url"`  // swagger地址
	AppJumpUrl  string `json:"app_jump_url"` // 应用跳转地址
}

func (App) TableName() string {
	return "sys_app"
}
