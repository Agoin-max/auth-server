package models

type User struct {
	Base
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Sex      string `json:"sex"`      // 性别
	Phone    string `json:"phone"`    // 电话
	Email    string `json:"email"`    // 邮箱
	Super    bool   `json:"super"`    // 是否超级管理员
	Status   int    `json:"status"`   // 状态
}

func (User) TableName() string {
	return "sys_user"
}
