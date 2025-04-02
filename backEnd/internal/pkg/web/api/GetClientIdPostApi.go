package api

import (
	"augeu/backEnd/internal/pkg/DBMnager"
	"augeu/backEnd/internal/pkg/DBMnager/HostInfo"
	"augeu/backEnd/internal/pkg/DBMnager/TokenTable"
	"augeu/backEnd/internal/utils/consts/web"
	convert2 "augeu/backEnd/internal/utils/convert"
	"augeu/backEnd/internal/utils/utils"
	"augeu/public/pkg/augeuJwt"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/snowNumbers"
	models2 "augeu/public/pkg/swaggerCore/models"
	operations2 "augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"errors"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) GetClientIdPostApiHandlerFunc() operations2.PostGetClientIDHandlerFunc {
	return func(getClientIdParams operations2.PostGetClientIDParams) middleware.Responder {
		apiName := "GetClientIdPostApi"
		data := getClientIdParams.Data
		if data == nil {
			logger.Error("data.Secret == nil")
			return operations2.NewPostGetClientIDUnauthorized().WithPayload(&models2.UnauthorizedError{
				Code:    convert.Int64P(int64(operations2.PostGetClientIDInternalServerErrorCode)),
				Message: convert.StrPtr("传入的参数为空"),
			})
		}
		secrete := data.Secret
		if secrete == nil {
			logger.Error("secrete == nil")
			return operations2.NewPostGetClientIDUnauthorized().WithPayload(&models2.UnauthorizedError{
				Code:    convert.Int64P(int64(operations2.PostGetClientIDInternalServerErrorCode)),
				Message: convert.StrPtr("secrete 为空"),
			})
		}

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
				Code:    convert.Int64P(int64(operations2.PostGetClientIDUnauthorizedCode)),
				Message: convert.StrPtr("token 错误"),
			})
		}

		clientData := convert2.GetClientIDRequest2Db(*getClientIdParams.Data)
		if err := HostInfo.InsertHostInfo(apiManager.s.DBM.DB, &clientData); err != nil {
			if !errors.Is(err, DBMnager.ErrDuplicateEntry) {
				logger.Errorf("GetClientIdPostApiHandlerFunc -> HostInfo.InsertHostInfo -> %v", err)
				return operations2.NewPostGetClientIDInternalServerError().WithPayload(&models2.ActionFailure{
					From:    utils.StrP(apiName),
					Reason:  utils.StrP(web.InternalError),
					Success: web.Fail,
				})
			}
			logger.Warnf("GetClientIdPostApiHandlerFunc -> HostInfo.InsertHostInfo -> %v", err)
		}

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
