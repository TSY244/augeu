package api

import (
	"augeu/backEnd/internal/pkg/server"
	"augeu/public/pkg/swaggerCore/restapi/operations"
	"log"
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
		// 处理预检请求
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			w.WriteHeader(http.StatusOK)
			return
		}

		// 设置 CORS 头信息
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		// 记录请求信息
		rw := &responseWriter{ResponseWriter: w}
		handler.ServeHTTP(rw, r)

		end := time.Now()
		duration := end.Sub(start)

		log.Printf("[%v] [%v] %v len:%v cost:%v response:[%v] body_len:[%v]",
			r.RemoteAddr, r.Method, r.RequestURI, r.ContentLength, duration.String(),
			rw.statusCode, rw.bodyLen,
		)
	}
}

// 自定义响应写入器，用于记录响应状态码和响应体长度
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	bodyLen    int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.bodyLen += n
	return n, err
}

func (apiManager *ApiManager) InitApi(swapi *operations.AugeuAPI) {
	swapi.GetVersionHandler = apiManager.GetVersionApiHandlerFunc()
	swapi.PostGetClientIDHandler = apiManager.GetClientIdPostApiHandlerFunc()
	swapi.PostLoginHandler = apiManager.LoginPostApiHandlerFunc()
	swapi.PostRegisterHandler = apiManager.RegisterPostApiHandlerFunc()
	swapi.PostUploadLoginEventHandler = apiManager.UploadLoginEventApiHandlerFunc()
	swapi.GetGetClientsHandler = apiManager.GetClientsGetHandlerFunc()
	swapi.PostGetLoginEventHandler = apiManager.GetLoginEventGetApi()
	swapi.PostUploadRdpEventHandler = apiManager.UploadRdpEventPostApiHandlerFunc()
	swapi.PostUploadRdpEventHandler = apiManager.UploadRdpEventPostApiHandlerFunc()
}
