package admin

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	AdminName       string `gorm:"column:adminname"`
	AdminPasswdHash string `gorm:"column:adminpasswdhash"`
}
