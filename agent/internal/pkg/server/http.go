package server

import (
	"augeu/agent/internal/pkg/systeminfo"
	"augeu/agent/internal/utils/convert"
	"augeu/agent/pkg/windowsLog"
	"augeu/public/pkg/augeuHttp"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	"augeu/public/util/utils"
	"encoding/json"
	"github.com/0xrawsec/golang-evtx/evtx"
)

const (
	GetClientIdApiPath      = "/getClientId"
	UploadLoginEventApiPath = "/upload/loginEvent"
	UploadRdpEventApiPath   = "/upload/rdpEvent"
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
		Secret: &s.Conf.Secret,
		ClientInfo: &models.ClientInfo{
			UUID:       &uuid,
			IP:         *ips,
			SystemInfo: info,
		},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", "", err
	}

	ret, err := augeuHttp.PostRequestWithJson(s.Conf.RemoteAddr+GetClientIdApiPath, s.Header, string(jsonData))
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

func (s *Server) PushLoginEvent(evtxMap chan *evtx.GoEvtxMap) error {
	events := windowsLog.GetEventsForLoginEvent(evtxMap)

	resq := convert.ArrayCopy(events, convert.LoginEvent2RLoginEventResq)
	jsonData, err := json.Marshal(resq)
	if err != nil {
		logger.Errorf("PushLoginEvent json.Marshal error: %v", err)
		return err
	}
	_, err = augeuHttp.PostRequestWithJson(s.Conf.RemoteAddr+UploadLoginEventApiPath, s.Header, string(jsonData))
	if err != nil {
		logger.Errorf("PushLoginEvent PostRequestWithJson error: %v", err)
		return err
	}
	logger.Infof("PushLoginEvent success")
	return nil
}

func (s *Server) PushRdpEvent(evtxMap chan *evtx.GoEvtxMap) error {
	events := windowsLog.GetEventsForRdpEvent(evtxMap)
	resq := convert.ArrayCopy(events, convert.RdpEvent2RRdpEventResq)
	jsonData, err := json.Marshal(resq)

}
