package TokenTable

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	CreateAt time.Time      `gorm:"autoCreateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
	Creator  string         `gorm:"column:creator"`
	Token    string
}

func (Token) TableName() string {
	return "token"
}

func GeneratorToken(creator, token string, db *gorm.DB) error {
	return db.Create(&Token{
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
