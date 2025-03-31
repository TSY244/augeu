package api

import (
	"augeu/backEnd/internal/pkg/web/gen/models"
	"augeu/backEnd/internal/pkg/web/gen/restapi/operations"
	"augeu/backEnd/internal/utils/consts/web"
	"augeu/backEnd/internal/utils/utils"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/snowNumbers"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) GetClientIdPostApiHandlerFunc() operations.PostGetClientIDHandlerFunc {
	return func(getClientIdParams operations.PostGetClientIDParams) middleware.Responder {
		apiName := "GetClientIdPostApi"
		clientId, err := snowNumbers.GetAnStrID()
		if err != nil {
			logger.Errorf("GetClientIdPostApiHandlerFunc -> snowNumbers.GetAnStrID -> %v", err)
			return operations.NewPostGetClientIDInternalServerError().WithPayload(&models.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}
		return operations.NewPostGetClientIDOK().WithPayload(&models.GetClientIDResponse{
			Success:  web.Success,
			ClientID: utils.StrP(clientId),
		})
	}
}
