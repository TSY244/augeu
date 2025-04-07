package api

import (
	"augeu/backEnd/internal/pkg/DBMnager/Log"
	middleware2 "augeu/backEnd/internal/pkg/web/middleware"
	"augeu/backEnd/internal/utils/consts/web"
	convert2 "augeu/backEnd/internal/utils/convert"
	"augeu/backEnd/internal/utils/utils"
	"augeu/public/pkg/swaggerCore/models"
	"augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) GetLoginEventGetApi() operations.PostGetLoginEventHandlerFunc {
	return func(params operations.PostGetLoginEventParams) middleware.Responder {
		if resp := middleware2.CheckUserRole(params.HTTPRequest, apiManager.s); resp != nil {
			return resp
		}

		tempData := params.Body

		if tempData == nil {
			return operations.NewPostGetLoginEventBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations.PostGetLoginEventBadRequestCode)),
				Message: convert.StrPtr("body is empty"),
			})
		}
		if tempData.PageAndSize == nil {
			return operations.NewPostGetLoginEventBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations.PostGetLoginEventBadRequestCode)),
				Message: convert.StrPtr("pageAndSize is empty"),
			})
		}
		if tempData.PageAndSize.Page == nil || tempData.PageAndSize.Size == nil {
			return operations.NewPostGetLoginEventBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations.PostGetLoginEventBadRequestCode)),
				Message: convert.StrPtr("page or size is empty"),
			})
		}
		page := *tempData.PageAndSize.Page
		size := *tempData.PageAndSize.Size

		queryData := &Log.QueryLoginEventParams{
			ClientId: tempData.ClientID,
			SourceIp: tempData.IP,
			UUID:     tempData.UUID,
			Page:     page,
			Size:     size,
		}
		loginEvents, total, err := Log.QueryLoginEvent(apiManager.s.RootCtx, apiManager.s.DBM.DB, queryData)
		if err != nil {
			return operations.NewPostGetLoginEventInternalServerError().WithPayload(&models.ActionFailure{
				From:    utils.StrP(InternalServerError),
				Reason:  utils.StrP(err.Error()),
				Success: web.Fail,
			})
		}
		loginEventModels := convert2.ArrayCopy(loginEvents, convert2.DbLoginEvent2modelLogEvent)
		return operations.NewPostGetLoginEventOK().WithPayload(&models.GetLoginEventResponse{
			Data: loginEventModels,
			Page: &models.PageMeta{
				Page:  convert.Int64P(page),
				Size:  convert.Int64P(size),
				Total: total,
			},
			Success: true,
		})
	}
}
