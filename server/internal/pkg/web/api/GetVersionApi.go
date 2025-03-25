package api

import (
	"augeu/server/internal/pkg/web/gen/models"
	"augeu/server/internal/pkg/web/gen/restapi/operations"
	"augeu/server/internal/utils"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) GetVersionApiHandlerFunc() operations.GetVersionHandlerFunc {
	return func(getVersionParams operations.GetVersionParams) middleware.Responder {
		return operations.NewGetVersionOK().WithPayload(&models.Version{
			Version: utils.StrP(Version),
		})
	}
}
