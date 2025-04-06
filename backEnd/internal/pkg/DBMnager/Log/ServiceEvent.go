package Log

// ServiceInfo represents the ServiceInfo table in the database.
type ServiceInfo struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"` // 主键，自增
	ServiceName string `gorm:"column:ServiceName"`       // 服务名称
	ImagePath   string `gorm:"column:ImagePath"`         // 镜像路径
	StartType   string `gorm:"column:StartType"`         // 启动类型
	Account     string `gorm:"column:Account"`           // 账户
	Description string `gorm:"column:description"`       // 描述
	UUID        string `gorm:"type:varchar(255);column:uuid"`
	EventID     string `gorm:"column:event_id"`
}

// TableName specifies the table name for the ServiceInfo model.
func (ServiceInfo) TableName() string {
	return "service_info"
}
