package models

type OpLog struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	UserID uint   `json:"user_id"` // 用户ID
	OpType string `json:"op_type"` // 操作类型
	OpDesc string `json:"op_desc"` // 操作描述
	Ip     string `json:"ip"`      // IP地址

}

func (OpLog) TableName() string {
	return "sys_op_log"
}
