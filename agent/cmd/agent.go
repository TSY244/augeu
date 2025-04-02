//go:build windows
// +build windows

package main

import (
	"augeu/agent/internal/pkg/config"
	"augeu/agent/internal/pkg/server"
	"flag"
	//config2 "augeu/public/pkg/config"
)

var (
	secret        = flag.String("s", "", "server 提供的 secret")
	remoteAddr    = flag.String("r", "", "server api 地址，请注意加上/api/v1，举个栗子：http://127.0.0.1/api/v1")
	websocketAddr = flag.String("ws", "", "server 提供的websocket 地址，注意使用ws://开头")
)

func main() {
	flag.Parse()
	if *secret == "" || *remoteAddr == "" || *websocketAddr == "" {
		flag.Usage()
		panic("Please provide all parameters")
	}

	conf := config.Config{
		RemoteAddr:    *remoteAddr,
		Secret:        *secret,
		WebsocketAddr: *websocketAddr,
	}

	agentServer, err := server.NewServer(&conf)
	if err != nil {
		panic(err)
	}
	agentServer.Run()
}
