package Log

import (
	"gorm.io/gorm"
	"time"
)

type LoginEvent struct {
	EventId         int64          `json:"event_id"`
	EventTime       time.Time      `json:"event_time"` // 这个事件创建的时间
	LoginType       string         `json:"login_type"`
	SourceIp        string         `json:"source_ip"`
	Username        string         `json:"username"`
	SubjectUsername string         `json:"subject_username"` // 主体用户名
	SubjectDomain   string         `json:"subject_domain"`   // 主体域
	ProcessName     string         `json:"process_name"`     // 进程名称
	CreateAt        time.Time      `gorm:"autoCreateTime"`   // 这条sql 记录创建的时间
	DeleteAt        gorm.DeletedAt `gorm:"index"`
	ID              uint           `gorm:"primaryKey;autoIncrement"`      // 主键
	UUID            string         `gorm:"type:varchar(255);column:uuid"` // windows 主机的uuid
}

func (LoginEvent) TableName() string {
	return "login_event"
}
