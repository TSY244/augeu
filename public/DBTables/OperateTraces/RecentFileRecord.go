package OperateTraces

import "time"

type RecentFileRecord struct {
	ID               uint      `gorm:"primaryKey;autoIncrement;type:serial"`
	FileName         string    `gorm:"type:varchar(255);not null"`
	FilePath         string    `gorm:"type:text;not null"`
	CreationTime     time.Time `gorm:"type:timestamp;not null"`
	ModificationTime time.Time `gorm:"type:timestamp;not null"`
	TargetPath       string    `gorm:"type:text;not null"`
	CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

func (RecentFileRecord) TableName() string {
	return "recent_file_record"
}
