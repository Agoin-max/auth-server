package models

type UserRole struct {
	ID     uint `gorm:"primary_key" json:"id"`
	UserID int  `json:"user_id"` // 用户ID
	RoleID int  `json:"role_id"` // 角色ID
}

func (UserRole) TableName() string {
	return "sys_user_role"
}
