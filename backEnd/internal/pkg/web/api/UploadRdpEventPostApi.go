package api

import (
	"augeu/backEnd/internal/pkg/DBMnager/Log"
	middleware2 "augeu/backEnd/internal/pkg/web/middleware"
	"augeu/backEnd/internal/utils/consts/web"
	convert2 "augeu/backEnd/internal/utils/convert"
	"augeu/backEnd/internal/utils/utils"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	"augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) UploadRdpEventPostApiHandlerFunc() operations.PostUploadRdpEventHandlerFunc {
	return func(params operations.PostUploadRdpEventParams) middleware.Responder {
		apiName := "UploadRdpEventApi"
		if resp := middleware2.CheckAgentRole(params.HTTPRequest, apiManager.s); resp != nil {
			return resp
		}
		uuid, err := middleware2.GetClientUuid(params.HTTPRequest)
		if err != nil {
			logger.Errorf("UploadRdpEventPostApiHandlerFunc -> GetClientUuid error: %v", err)
			return operations.NewPostGetRdpEventBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations.PostGetRdpEventBadRequestCode)),
				Message: convert.StrPtr("GetClientUuid error"),
			})
		}
		events := params.Data
		if events == nil {
			return operations.NewPostGetRdpEventBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations.PostGetRdpEventBadRequestCode)),
				Message: convert.StrPtr("events is nil"),
			})
		}
		ret := utils.CheckUuid(events, func(unit *models.RDPEventUnit) bool {
			if unit.Base == nil {
				return false
			}
			if unit.Base.UUID == nil {
				return false
			}
			if *unit.Base.UUID == uuid {
				return true
			}
			return false
		})
		if len(ret) == 0 {
			return operations.NewPostGetRdpEventBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations.PostGetRdpEventBadRequestCode)),
				Message: convert.StrPtr("uuid is not match"),
			})
		}
		dbEvent := convert2.ArrayCopy(events, convert2.ModelRdpEvent2DbRdpEvent)
		if err := Log.InsertRdpEventBatch(apiManager.s.RootCtx, apiManager.s.DBM.DB, dbEvent); err != nil {
			logger.Errorf("UploadRdpEventPostApiHandlerFunc -> HostInfo.InsertRdpEvent error: %v", err)
			return operations.NewPostGetRdpEventInternalServerError().WithPayload(&models.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}
		return operations.NewPostUploadEventLoginOK().WithPayload(&models.EventLogUploadResponse{
			Message: convert.StrPtr("Success"),
			Success: web.Success,
		})
	}
}
