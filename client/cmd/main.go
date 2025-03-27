package cmd

import (
	"augeu/client/internal/pkg/config"
	"augeu/client/internal/pkg/server"
	//config2 "augeu/public/pkg/config"
)

func main() {
	conf := config.Config{
		Websocket: config.WebsocketConf{
			RemoteAddr: "",
			RemotePort: 0,
		},
	}

	server, err := server.NewServer(&conf)
	if err != nil {
		panic(err)
	}

}
