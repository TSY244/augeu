package server

import (
	"augeu/agent/internal/pkg/systeminfo"
	"augeu/public/pkg/augeuHttp"
	"augeu/public/pkg/swaggerCore/models"
	"encoding/json"
)

const (
	GetClientIdApiPath = "/getClientId"
)

func (s *Server) GetClientId() (string, error) {
	uuid, err := systeminfo.GetUuid()
	if err != nil {
		return "", err
	}
	payload := models.GetClientIDRequest{
		Secret: &s.Conf.Secret,
		UUID:   &uuid,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	ret, err := augeuHttp.PostRequestWithJson(s.Conf.RemoteAddr+GetClientIdApiPath, map[string]string, string(jsonData))
	if err != nil {
		return "", err
	}
	var resp models.GetClientIDResponse
	err = json.Unmarshal([]byte(ret), &resp)
	if err != nil {
		return "", err
	}
	return *resp.ClientID, nil

}
