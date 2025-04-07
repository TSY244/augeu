package server

import "augeu/agent/pkg/windowsLog"

func (s *Server) GetLoginEventInfo() error {
	return windowsLog.Run(windowsLog.LoginEvenType)
}

func (s *Server) GetEventInfo(eventName windowsLog.EventNameType, mapChanFunctions ...windowsLog.ExternalFunctionForMapChan) error {
	return windowsLog.Run(eventName, mapChanFunctions...)
}
