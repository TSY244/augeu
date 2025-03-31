package server

type Config struct {
	ListenPort       int
	EnableStaticAuth bool
	DBHost           string
	DBPort           int
	DBUser           string
	DBPassword       string
	DBName           string
}
