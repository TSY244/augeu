package server

func (s *Server) GetRule() (string, error) {
	return s.Config.CoreConfig.Rule.Rules, nil
}
