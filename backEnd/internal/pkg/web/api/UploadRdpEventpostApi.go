package api

import (
	middleware2 "augeu/backEnd/internal/pkg/web/middleware"
	"augeu/public/pkg/swaggerCore/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) UploadRdpEventApiHandlerFunc() operations.PostGetRdpEventHandlerFunc {
	return func(params operations.PostGetRdpEventParams) middleware.Responder {
		if resp := middleware2.CheckAgentRole(params.HTTPRequest, apiManager.s); resp != nil {
			return resp
		}
		data := params.Body

	}
}
