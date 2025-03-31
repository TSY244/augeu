//go:build windows
// +build windows

package main

import (
	"augeu/agent/internal/pkg/config"
	"augeu/agent/internal/pkg/server"
	//config2 "augeu/public/pkg/config"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		panic(err)
	}

	agentServer, err := server.NewServer(conf)
	if err != nil {
		panic(err)
	}
	agentServer.Run()
}
