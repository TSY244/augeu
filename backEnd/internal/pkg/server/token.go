package server

import (
	"augeu/backEnd/internal/pkg/DBMnager/TokenTable"
	"augeu/public/pkg/logger"
	"errors"
	"gorm.io/gorm"
)

// CheckToken 检查是否已经有已经存在的token
func (s *Server) CheckToken() (string, bool, error) {
	token, err := TokenTable.GetToken(s.DBM.DB)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", false, nil
		}
		logger.Errorf("Failed to get token: %v", err)
		return "", false, err
	}
	return token, true, nil
}
