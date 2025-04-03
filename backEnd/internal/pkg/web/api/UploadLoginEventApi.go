package api

import (
	"augeu/backEnd/internal/pkg/DBMnager/Log"
	middleware2 "augeu/backEnd/internal/pkg/web/middleware"
	convert2 "augeu/backEnd/internal/utils/convert"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	operations2 "augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) UploadLoginEventApiHandlerFunc() operations2.PostUploadLoginEventHandlerFunc {
	return func(params operations2.PostUploadLoginEventParams) middleware.Responder {
		// check jwt
		if resp := middleware2.CheckAgentRole(params.HTTPRequest, apiManager.s); resp != nil {
			return resp
		}
		uuid, err := middleware2.GetClientUuid(params.HTTPRequest)
		if err != nil {
			logger.Errorf("UploadLoginEventApiHandlerFunc -> GetClientUuid error: %v", err)
			return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
				Message: convert.StrPtr("GetClientUuid error"),
			})
		}

		event := params.EventLog
		if event == nil {
			logger.Errorf("UploadLoginEventApiHandlerFunc -> event is nil")
			return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
				Message: convert.StrPtr("event is nil"),
			})
		}

		dbData := convert2.ArrayCopy(event, convert2.LoginEvent2Db)
		instertData := make([]*Log.LoginEvent, 0)
		for _, data := range dbData {
			if data.UUID != uuid {
				logger.Errorf("UploadLoginEventApiHandlerFunc -> uuid is invalid, data.UUID: %v, uuid: %v", data.UUID, uuid)
				continue
			}
			instertData = append(instertData, &data)
		}
		if err := Log.InsertLoginEventBatch(apiManager.s.RootCtx, apiManager.s.DBM.DB, instertData); err != nil {
			logger.Errorf("UploadLoginEventApiHandlerFunc -> InsertLoginEventBatch error: %v", err)
			return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
				Message: convert.StrPtr("InsertLoginEventBatch error"),
			})
		}

		return operations2.NewPostUploadLoginEventOK().WithPayload(&models.UploadLoginEventResponse{
			Message: convert.StrPtr("add insert success"),
			Success: true,
		})
	}
}
