package api

import (
	"augeu/backEnd/internal/pkg/web/gen/models"
	"augeu/backEnd/internal/pkg/web/gen/restapi/operations"
	"augeu/backEnd/internal/utils/utils"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) GetVersionApiHandlerFunc() operations.GetVersionHandlerFunc {
	return func(getVersionParams operations.GetVersionParams) middleware.Responder {
		return operations.NewGetVersionOK().WithPayload(&models.Version{
			Version: utils.StrP(Version),
		})
	}
}
