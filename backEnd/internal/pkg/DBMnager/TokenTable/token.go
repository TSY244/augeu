package TokenTable

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Token struct {
	CreateAt time.Time      `gorm:"autoCreateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
	Creator  string         `gorm:"column:creator"`
	Token    string         `gorm:"primarykey"`
}

func (Token) TableName() string {
	return "token"
}

func GeneratorToken(creator, token string, db *gorm.DB) error {
	// 忽视冲突
	return db.
		Clauses(clause.OnConflict{
			DoNothing: true, // 冲突时什么也不做
		}).
		Create(&Token{
			Creator:  creator,
			Token:    token,
			DeleteAt: gorm.DeletedAt{Time: time.Now().AddDate(0, 0, 7)},
		}).Error
}

func GetToken(db *gorm.DB) (string, error) {
	var token Token
	if err := db.First(&token).Error; err != nil {
		return "", err
	}
	return token.Token, nil
}
