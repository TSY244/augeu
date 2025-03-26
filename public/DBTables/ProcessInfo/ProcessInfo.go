package ProcessInfo

import "time"

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
	Imports        []ProcessImport `gorm:"foreignKey:PID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ParentProcess  *ProcessInfo    `gorm:"foreignKey:ParentPID;references:PID"`
}

func (ProcessInfo) TableName() string {
	return "process_info"
}
