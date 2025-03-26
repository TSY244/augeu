package OperateTraces

import "time"

type PrefetchRecord struct {
	ID             uint      `gorm:"primaryKey;autoIncrement;type:serial"`
	LaunchTime     time.Time `gorm:"type:timestamp;not null"`
	ExecutableName string    `gorm:"type:varchar(255);not null"`
	ExecutableMD5  string    `gorm:"type:char(32);not null;column:executable_md5"`
	ExecutablePath string    `gorm:"type:text;not null"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

func (PrefetchRecord) TableName() string {
	return "prefetch_record"
}
