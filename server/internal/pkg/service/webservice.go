package service

import (
	"augeu/public/pkg/logger"
	"augeu/server/internal/pkg/server"
	"augeu/server/internal/pkg/web/api"
	"augeu/server/internal/pkg/web/gen/restapi"
	"augeu/server/internal/pkg/web/gen/restapi/operations"
	"context"
	"errors"
	"github.com/go-openapi/loads"
)

func StartApi(bPServer *server.Server) error {
	if bPServer == nil {
		return errors.New("server is nil")
	}
	swaSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}
	apiBase := operations.NewBrightPathAPI(swaSpec)
	apiBase.Logger = logger.Infof

	apiManager := api.NewApiManager(bPServer)
	apiManager.InitApi(apiBase)

	webServer := restapi.NewServer(apiBase)
	webServer.Port = bPServer.Config.ListenPort
	webServer.ConfigureAPI()
	webServer.SetHandler(
		apiManager.HookHttpMiddleware(webServer.GetHandler()),
	)
	webServer.EnabledListeners = []string{"http"}
	go func() {
		if err := webServer.Serve(); err != nil {
			logger.Errorf("serve failed: %v", err)
		}
	}()

	go serveDaemonWithCtx(webServer, bPServer.RootCtx)
	return nil
}

func serveDaemonWithCtx(webServer *restapi.Server, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("web service stop")
			err := webServer.Shutdown()
			if err != nil {
				logger.Errorf("web service stop failed: %v", err)
				return
			}
			return
		default:
			continue
		}
	}
}
