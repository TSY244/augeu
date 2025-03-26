package HostInfo

// 镜像劫持表
type ImageHijack struct {
	HijackName   string `gorm:"primaryKey;type:varchar(255)"`
	HijackPath   string `gorm:"type:text;not null"`
	HijackStatus bool   `gorm:"not null;default:false"`
}

func (ImageHijack) TableName() string {
	return "image_hijacks"
}
