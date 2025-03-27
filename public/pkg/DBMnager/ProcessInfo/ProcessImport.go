package ProcessInfo

import (
	"gorm.io/gorm"
	"time"
)

type ProcessImport struct {
	ImportID        uint           `gorm:"autoIncrement;column:import_id;type:serial"`
	PID             int            `gorm:"column:pid;type:integer;not null"`
	ImportPath      string         `gorm:"column:import_path;type:text;not null"`
	ImportMD5       string         `gorm:"column:import_md5;type:char(32)"`
	ImportSignature string         `gorm:"column:import_signature;type:text"`
	CreateAt        time.Time      `gorm:"autoCreateTime"`
	DeleteAt        gorm.DeletedAt `gorm:"index"`
	ID              uint           `gorm:"primaryKey;autoIncrement"`
	UUID            string         `gorm:"type:varchar(255);column:uuid"`
}

func (ProcessImport) TableName() string {
	return "process_imports"
}
