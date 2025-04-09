package Log

import "time"

// ServiceInfo represents the ServiceInfo table in the database.
type ServiceInfo struct {
	ID          uint      `gorm:"primaryKey" json:"id"`                    // 自增主键
	EventID     int       `gorm:"column:event_id" json:"event_id"`         // 事件ID (7045)
	EventTime   time.Time `gorm:"column:event_time" json:"event_time"`     // 事件时间
	MachineUUID string    `gorm:"column:machine_uuid" json:"machine_uuid"` // 机器唯一标识
	ServiceName string    `gorm:"column:service_name" json:"service_name"` // 服务名称
	ImagePath   string    `gorm:"column:image_path" json:"image_path"`     // 可执行文件路径
	ServiceType string    `gorm:"column:service_type" json:"service_type"` // 服务类型（内核模式驱动程序）
	StartType   string    `gorm:"column:start_type" json:"start_type"`     // 启动类型（系统启动）
	AccountName string    `gorm:"column:account_name" json:"account_name"` // 运行账户（空）
	UUID        string    `gorm:"type:varchar(255);column:uuid"`
}

// TableName specifies the table name for the ServiceInfo model.
func (ServiceInfo) TableName() string {
	return "service_info"
}
