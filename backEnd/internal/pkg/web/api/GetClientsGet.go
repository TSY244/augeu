package api

import (
	"augeu/backEnd/internal/pkg/DBMnager/HostInfo"
	middleware2 "augeu/backEnd/internal/pkg/web/middleware"
	"augeu/backEnd/internal/utils/consts/web"
	convert2 "augeu/backEnd/internal/utils/convert"
	"augeu/backEnd/internal/utils/utils"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	"augeu/public/pkg/swaggerCore/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) GetClientsGetHandlerFunc() operations.GetGetClientsHandlerFunc {
	return func(params operations.GetGetClientsParams) middleware.Responder {
		if resp := middleware2.CheckJwt(params.HTTPRequest, apiManager.s); resp != nil {
			return resp
		}
		clientIds, err := apiManager.s.GetAllClientId()
		if err != nil {
			logger.Errorf("GetClientsGetHandlerFunc -> GetAllClientId error: %v", err)
			return operations.NewGetGetClientsInternalServerError().WithPayload(&models.ActionFailure{
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}
		clients, err := HostInfo.GetAgentsByClientIds(apiManager.s.DBM.DB, clientIds)
		if err != nil {
			logger.Errorf("GetClientsGetHandlerFunc -> GetAgentsByClientIds error: %v", err)
			return operations.NewGetGetClientsInternalServerError().WithPayload(&models.ActionFailure{
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}
		clientInfo := convert2.ArrayCopy(clients, convert2.DbHostinfo2moduleHostinfo)
		return operations.NewGetGetClientsOK().WithPayload(&models.GetClientsResponse{
			Clients: clientInfo,
		})
	}
}
