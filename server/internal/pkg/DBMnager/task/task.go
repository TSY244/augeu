package task

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskTime    string `gorm:"column:tasktime"` // yyyy-mm-dd hh:mm:ss
	VolunteerId string `gorm:"column:volunteerid"`
}
