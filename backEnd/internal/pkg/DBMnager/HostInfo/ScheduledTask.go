package HostInfo

import (
	"gorm.io/gorm"
	"time"
)

// 定时任务表
type ScheduledTask struct {
	Hostname           string    `gorm:"primaryKey;type:varchar(255)"`
	TaskName           string    `gorm:"primaryKey;type:text"`
	NextRunTime        time.Time `gorm:"type:timestamp"`
	Mode               string    `gorm:"type:varchar(50)"`
	LogonStatus        string    `gorm:"type:varchar(50)"`
	LastRunTime        time.Time `gorm:"type:timestamp"`
	LastResult         int       `gorm:"type:integer"`
	Creator            string    `gorm:"type:varchar(255)"`
	TaskAction         string    `gorm:"type:text"`
	StartCondition     string    `gorm:"type:varchar(50)"`
	Comment            string    `gorm:"type:text"`
	TaskStatus         string    `gorm:"type:varchar(50);not null"`
	IdleSettings       string    `gorm:"type:varchar(50)"`
	PowerManagement    string    `gorm:"type:text"`
	RunAsUser          string    `gorm:"type:varchar(255);not null"`
	DeleteExpiredTask  string    `gorm:"type:varchar(50)"`
	ExecutionTimeLimit string    `gorm:"type:interval"` // PostgreSQL interval类型
	ScheduleData       string    `gorm:"type:text"`
	ScheduleType       string    `gorm:"type:varchar(50)"`
	StartTime          time.Time `gorm:"type:time"`
	StartDate          time.Time `gorm:"type:date"`
	EndDate            time.Time `gorm:"type:date"`

	ID       uint           `gorm:"primaryKey;autoIncrement"`
	UUID     string         `gorm:"type:varchar(255);column:uuid"`
	CreateAt time.Time      `gorm:"autoCreateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

func (ScheduledTask) TableName() string {
	return "scheduled_tasks"
}
