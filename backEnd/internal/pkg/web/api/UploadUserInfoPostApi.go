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
	"augeu/public/util/convert"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) UploadUserInfoPostApiHandlerFunc() operations.PostUploadUserInfoHandlerFunc {
	return func(params operations.PostUploadUserInfoParams) middleware.Responder {
		apiName := "UploadUserInfoPostApi"
		if resp := middleware2.CheckAgentRole(params.HTTPRequest, apiManager.s); resp != nil {
			return resp
		}
		data := params.Data

		if data == nil {
			return operations.NewPostUploadUserInfoBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations.PostUploadUserInfoBadRequestCode)),
				Message: convert.StrPtr("param is nil"),
			})
		}

		for _, userInfo := range data {
			if userInfo.UUID == nil || userInfo.Name == nil ||
				userInfo.Description == nil || userInfo.IsFocus == nil ||
				userInfo.LocalAccount == nil {
				return operations.NewPostUploadUserInfoBadRequest().WithPayload(&models.BadRequestError{
					Code:    convert.Int64P(int64(operations.PostUploadUserInfoBadRequestCode)),
					Message: convert.StrPtr("param is nil"),
				})
			}
		}
		users := convert2.ArrayCopy(data, convert2.ModelUser2DbUser)
		if err := HostInfo.InsertUserBatch(apiManager.s.RootCtx, apiManager.s.DBM.DB, users); err != nil {
			logger.Errorf("UploadUserInfoPostApiHandlerFunc -> HostInfo.InsertUserBatch -> %v", err)
			return operations.NewPostUploadUserInfoInternalServerError().WithPayload(&models.ActionFailure{
				From:    utils.StrP(apiName),
				Reason:  utils.StrP(web.InternalError),
				Success: web.Fail,
			})
		}
		return operations.NewPostUploadUserInfoOK().WithPayload(&models.SuccessResponse{
			Code:    convert.Int64P(200),
			Message: convert.StrPtr("Success"),
			Success: web.Success,
		})
	}
}
