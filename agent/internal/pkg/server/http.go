package server

import (
	"augeu/agent/internal/pkg/systeminfo"
	"augeu/public/pkg/augeuHttp"
	"augeu/public/pkg/swaggerCore/models"
	"augeu/public/util/utils"
	"encoding/json"
)

const (
	GetClientIdApiPath = "/getClientId"
)

func (s *Server) GetClientId() (string, string, error) {
	uuid, err := systeminfo.GetUuid()
	if err != nil {
		return "", "", err
	}
	ips, err := utils.GetIps()
	if err != nil {
		return "", "", err
	}
	info, err := systeminfo.GetSystemInfo()

	payload := models.GetClientIDRequest{
		Secret:     &s.Conf.Secret,
		UUID:       &uuid,
		IP:         *ips,
		SystemInfo: info,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", "", err
	}

	ret, err := augeuHttp.PostRequestWithJson(s.Conf.RemoteAddr+GetClientIdApiPath, map[string]string{}, string(jsonData))
	if err != nil {
		return "", "", err
	}
	var resp models.GetClientIDResponse
	err = json.Unmarshal([]byte(ret), &resp)
	if err != nil {
		return "", "", err
	}
	if resp == (models.GetClientIDResponse{}) {
		return "", "", err
	}
	return resp.Jwt, *resp.ClientID, nil

}
