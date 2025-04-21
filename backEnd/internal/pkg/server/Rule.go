package server

func (s *Server) GetRule() (string, error) {
	s.RuleRWLocker.RLocker()
	defer s.RuleRWLocker.RUnlock()
	return s.Config.CoreConfig.Rule.Rules, nil
}
