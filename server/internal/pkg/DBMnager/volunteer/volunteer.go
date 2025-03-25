package volunteer

import (
	"augeu/server/internal/pkg/DBMnager/task"
	"gorm.io/gorm"
)

type Volunteer struct {
	gorm.Model
	VolunteerId   string      `gorm:"column:volunteerid;primaryKey;unique"`
	VolunteerName string      `gorm:"column:volunteername"`
	VolunteerDesc string      `gorm:"column:volunteerdesc"`
	Password      string      `gorm:"column:password"` //Should be hash
	Email         string      `gorm:"column:email"`
	Phone         string      `gorm:"column:phone"`
	Avatar        string      `gorm:"column:avatar"`
	IdCardFront   string      `gorm:"column:idcardfront"`
	IdCardBack    string      `gorm:"column:idcardback"`
	IDNumber      string      `gorm:"column:idnumber"`
	RealName      string      `gorm:"column:realname"`
	VolunteerTime string      `gorm:"column:volunteertime"`
	IsRegister    bool        `gorm:"column:isregister"`
	Integral      uint        `gorm:"column:integral"`
	Tasks         []task.Task `gorm:"foreignKey:VolunteerId;references:VolunteerId"`
}
