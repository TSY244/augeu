package api

import (
	"augeu/backEnd/internal/pkg/DBMnager/TokenTable"
	"augeu/backEnd/internal/pkg/DBMnager/UserInfo"
	"augeu/backEnd/internal/utils/consts/web"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	operations2 "augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"github.com/go-openapi/runtime/middleware"
	"strings"
)

func (apiManager *ApiManager) RegisterPostApi() operations2.PostRegisterHandlerFunc {
	return func(postRegisterParams operations2.PostRegisterParams) middleware.Responder {
		// 检测是否是本地的请求
		remoteAddr := postRegisterParams.HTTPRequest.RemoteAddr
		if !strings.Contains(remoteAddr, "127.0.0.1") {
			return operations2.NewPostRegisterBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations2.PostRegisterBadRequestCode)),
				Message: convert.StrPtr("非法请求"),
			})
		}

		apiName := "RegisterPostApi"
		name := postRegisterParams.Data.UserName
		password := postRegisterParams.Data.PassWord
		secrete := postRegisterParams.Data.Secrete
		if name == nil || password == nil || secrete == nil {
			return operations2.NewPostRegisterBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations2.PostRegisterBadRequestCode)),
				Message: convert.StrPtr("传入的参数为空"),
			})
		}

		token, err := TokenTable.GetToken(apiManager.s.DBM.DB)
		if err != nil {
			logger.Errorf("GetClientIdPostApiHandlerFunc -> TokenTable.GetToken -> %v", err)
			return operations2.NewPostGetClientIDInternalServerError().WithPayload(&models.ActionFailure{
				From:    convert.StrPtr(apiName),
				Reason:  convert.StrPtr("检测token 出错"),
				Success: web.Fail,
			})
		}
		if *secrete != token {
			logger.Error("secrete != token")
			return operations2.NewPostGetClientIDUnauthorized().WithPayload(&models.UnauthorizedError{
				Code:    convert.Int64P(int64(operations2.PostLoginUnauthorizedCode)),
				Message: convert.StrPtr("token 错误"),
			})
		}

		if err := UserInfo.AddUser(*name, *password, apiManager.s.DBM.DB); err != nil {
			logger.Errorf("RegisterPostApiHandlerFunc -> UserInfo.AddUser -> %v", err)
			return operations2.NewPostRegisterInternalServerError().WithPayload(&models.ActionFailure{
				From:    convert.StrPtr(apiName),
				Reason:  convert.StrPtr("添加用户失败"),
				Success: web.Fail,
			})
		}
		return operations2.NewPostRegisterOK().WithPayload(&models.RegisterResponse{
			Message: convert.StrPtr("添加用户成功"),
			Success: web.Success,
		})
	}
}
