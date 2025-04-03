package server

import "augeu/agent/pkg/windowsLog"

func (s *Server) GetLoginEventInfo() error {
	return windowsLog.Run(windowsLog.LoginEvenType)
}
