package server

import config2 "augeu/backEnd/internal/pkg/config"

type Config struct {
	CoreConfig       *config2.Config
	ListenPort       int
	EnableStaticAuth bool
	DBHost           string
	DBPort           int
	DBUser           string
	DBPassword       string
	DBName           string
}
