package main

import (
	config2 "augeu/server/internal/pkg/config"
	"augeu/server/internal/pkg/logger"
	"augeu/server/internal/pkg/server"
	"augeu/server/internal/pkg/service"
	"context"
	"log"
	"os"
	"os/signal"
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

	config := server.Config{
		ListenPort:       config2.CoreConfig.HttpServer.HttpPort,
		EnableStaticAuth: config2.CoreConfig.HttpServer.EnableStaticAuth,
		DBHost:           config2.CoreConfig.DbConfig.Host,
		DBPort:           config2.CoreConfig.DbConfig.Port,
		DBUser:           config2.CoreConfig.DbConfig.Username,
		DBPassword:       config2.CoreConfig.DbConfig.Password,
		DBName:           config2.CoreConfig.DbConfig.Dbname,
	}
	bPServer, err := server.NewServer(&config)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	cancelCtxSilence = append(cancelCtxSilence, bPServer.Cancel)

	err = service.StartApi(bPServer)
	if err != nil {
		log.Fatalf("Failed to start service: %v", err)
		return
	}

	select {}
}
