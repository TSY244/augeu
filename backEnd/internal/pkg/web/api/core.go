package api

import (
	"augeu/backEnd/internal/pkg/server"
	"augeu/backEnd/internal/pkg/web/gen/restapi/operations"
	"augeu/public/pkg/logger"
	"net/http"
	"time"
)

type ApiManager struct {
	s *server.Server
}

var Version = "v0.0.1"

func NewApiManager(s *server.Server) *ApiManager {
	return &ApiManager{
		s: s,
	}
}

func (apiManager *ApiManager) HookHttpMiddleware(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			end := time.Now()
			duration := end.Sub(start)

			if r.Response != nil {
				logger.Infof("[%v] [%v] %v len:%v cost:%v response:[%v] body_len:[%v]",
					r.RemoteAddr, r.Method, r.RequestURI, r.ContentLength, duration.String(),
					r.Response.StatusCode, r.Response.ContentLength,
				)
			} else {
				logger.Info("[%v] [%v] %v len:%v cost:%v",
					r.RemoteAddr, r.Method, r.RequestURI, r.ContentLength, duration.String())
			}
		}()
		// cors header
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		handler.ServeHTTP(w, r)
	}
}

func (apiManager *ApiManager) InitApi(swapi *operations.AugeuAPI) {
	//todo add api handler
	swapi.GetVersionHandler = apiManager.GetVersionApiHandlerFunc()
	swapi.PostGetClientIDHandler = apiManager.GetClientIdPostApiHandlerFunc()

}
