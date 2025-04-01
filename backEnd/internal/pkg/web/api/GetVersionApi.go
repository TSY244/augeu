package api

import (
	"augeu/backEnd/internal/utils/utils"
	"augeu/public/pkg/swaggerCore/models"
	operations2 "augeu/public/pkg/swaggerCore/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) GetVersionApiHandlerFunc() operations2.GetVersionHandlerFunc {
	return func(getVersionParams operations2.GetVersionParams) middleware.Responder {
		return operations2.NewGetVersionOK().WithPayload(&models.Version{
			Version: utils.StrP(Version),
		})
	}
}
