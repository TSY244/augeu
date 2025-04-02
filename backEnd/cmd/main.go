package main

import (
	"augeu/backEnd/internal/pkg/DBMnager/TokenTable"
	config2 "augeu/backEnd/internal/pkg/config"
	"augeu/backEnd/internal/pkg/server"
	"augeu/backEnd/internal/pkg/service"
	"augeu/backEnd/internal/utils/utils"
	"augeu/public/pkg/logger"
	utils2 "augeu/public/util/utils"
	"context"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	sigExitOnce      = new(sync.Once)
	cancelCtxSilence = make([]context.CancelFunc, 0)
)

func init() {
	go sigExitOnce.Do(func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		defer signal.Stop(c)

		for {
			select {
			case <-c:
				for _, ctx := range cancelCtxSilence {
					ctx()
				}
				time.Sleep(500 * time.Millisecond)
				logger.Warn("Server exit by signal")
				os.Exit(1)
				return
			}
		}
	})
}

func main() {

	// 初始化配置
	config2.Init()

	config := server.Config{
		ListenPort:       config2.CoreConfig.HttpServer.HttpPort,
		EnableStaticAuth: config2.CoreConfig.HttpServer.EnableStaticAuth,
		DBHost:           config2.CoreConfig.DbConfig.Host,
		DBPort:           config2.CoreConfig.DbConfig.Port,
		DBUser:           config2.CoreConfig.DbConfig.Username,
		DBPassword:       config2.CoreConfig.DbConfig.Password,
		DBName:           config2.CoreConfig.DbConfig.Dbname,
		CoreConfig:       &config2.CoreConfig,
	}
	DwServer, err := server.NewServer(&config)
	if err != nil {
		logger.Fatalf("Failed to create server: %v", err)
	}

	// 启动websocket服务
	DwServer.RunWebsocket()

	cancelCtxSilence = append(cancelCtxSilence, DwServer.Cancel)

	err = service.StartApi(DwServer)
	if err != nil {
		logger.Fatalf("Failed to start service: %v", err)
		return
	}

	// get tempToken
	var token string
	if tempToken, isHave, err := DwServer.CheckToken(); err != nil {
		logger.Fatalf("Failed to get tempToken: %v", err)
		return
	} else if isHave {
		token = tempToken
	} else {
		token = utils.GenerateToken()
	}
	//DwServer.Secrete = token
	logger.Infof("token: %s", token)

	ips, err := utils2.GetIps()
	if err != nil {
		logger.Fatalf("Failed to get ips: %v", err)
		return
	}
	// 拼接ips
	ip := strings.Join(*ips, ",")

	if err := TokenTable.GeneratorToken(ip, token, DwServer.DBM.DB); err != nil {
		logger.Fatalf("Failed to generator tempToken: %v", err)
		return
	}

	select {}
}
