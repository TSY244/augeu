package api

import (
	"augeu/backEnd/internal/pkg/DBMnager/TokenTable"
	"augeu/backEnd/internal/utils/consts/web"
	"augeu/backEnd/internal/utils/utils"
	"augeu/public/pkg/augeuJwt"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/snowNumbers"
	models2 "augeu/public/pkg/swaggerCore/models"
	operations2 "augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) GetClientIdPostApiHandlerFunc() operations2.PostGetClientIDHandlerFunc {
	return func(getClientIdParams operations2.PostGetClientIDParams) middleware.Responder {
		apiName := "GetClientIdPostApi"
		secrete := getClientIdParams.Data.Secret
		token, err := TokenTable.GetToken(apiManager.s.DBM.DB)
		if err != nil {
			logger.Errorf("GetClientIdPostApiHandlerFunc -> TokenTable.GetToken -> %v", err)
			return operations2.NewPostGetClientIDInternalServerError().WithPayload(&models2.ActionFailure{
				From:    convert.StrPtr(apiName),
				Reason:  convert.StrPtr("检测token 出错"),
				Success: web.Fail,
			})
		}
		if *secrete != token {
			logger.Error("secrete != token")
			return operations2.NewPostGetClientIDUnauthorized().WithPayload(&models2.UnauthorizedError{
				Code:    convert.Int64P(401),
				Message: convert.StrPtr("token 错误"),
			})
		}
		//uuid := getClientIdParams.Data.UUID
		clientId, err := snowNumbers.GetAnStrID()
		if err != nil {
			logger.Errorf("GetClientIdPostApiHandlerFunc -> snowNumbers.GetAnStrID -> %v", err)
			return operations2.NewPostGetClientIDInternalServerError().WithPayload(&models2.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}

		info := augeuJwt.Info{
			ClientId: clientId,
			Uuid:     *getClientIdParams.Data.UUID,
		}

		strJwt, err := augeuJwt.NewJwt(info)
		if err != nil {
			logger.Errorf("GetClientIdPostApiHandlerFunc -> augeuJwt.NewJwt -> %v", err)
			return operations2.NewPostGetClientIDInternalServerError().WithPayload(&models2.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}

		return operations2.NewPostGetClientIDOK().WithPayload(&models2.GetClientIDResponse{
			Success:  web.Success,
			ClientID: utils.StrP(clientId),
			Jwt:      strJwt,
		})
	}
}
