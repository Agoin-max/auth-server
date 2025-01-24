package models

type Dept struct {
	Base
	DeptName      string `json:"dept_name"`       // 部门名称
	DeptShortName string `json:"dept_short_name"` // 部门简称
	DeptCode      string `json:"dept_code"`       // 部门编码
	Status        int    `json:"status"`          // 状态
	Description   string `json:"description"`     // 描述
}

func (Dept) TableName() string {
	return "sys_dept"
}
