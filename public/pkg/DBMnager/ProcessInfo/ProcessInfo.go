package ProcessInfo

import (
	"gorm.io/gorm"
	"time"
)

type ProcessInfo struct {
	PID            int             `gorm:"primaryKey;column:pid;type:integer"`
	ProcessName    string          `gorm:"column:process_name;type:text;not null"`
	ParentPID      *int            `gorm:"column:parent_pid;type:integer"`
	ParentName     string          `gorm:"column:parent_name;type:text"`
	CreateTime     time.Time       `gorm:"column:create_time;type:timestamp with time zone;not null"`
	ExecutablePath string          `gorm:"column:executable_path;type:text;not null"`
	FileCreateTime *time.Time      `gorm:"column:file_create_time;type:timestamp with time zone"`
	FileModifyTime *time.Time      `gorm:"column:file_modify_time;type:timestamp with time zone"`
	FileMD5        string          `gorm:"column:file_md5;type:char(32)"`
	SignatureInfo  string          `gorm:"column:signature_info;type:text"`
	Imports        []ProcessImport `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreateAt       time.Time       `gorm:"autoCreateTime"`
	DeleteAt       gorm.DeletedAt  `gorm:"index"`
	ID             uint            `gorm:"primaryKey;autoIncrement"`
	UUID           string          `gorm:"type:varchar(255);column:uuid"`
}

func (ProcessInfo) TableName() string {
	return "process_info"
}
