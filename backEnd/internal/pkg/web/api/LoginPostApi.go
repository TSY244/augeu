package api

import (
	"augeu/backEnd/internal/pkg/DBMnager/TokenTable"
	"augeu/backEnd/internal/pkg/DBMnager/UserInfo"
	"augeu/backEnd/internal/utils/consts/web"
	"augeu/backEnd/internal/utils/utils"
	"augeu/public/pkg/augeuJwt"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	operations2 "augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"errors"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) LoginPostApiHandlerFunc() operations2.PostLoginHandlerFunc {
	return func(postLoginParams operations2.PostLoginParams) middleware.Responder {
		apiName := "LoginPostApi"
		credentials := postLoginParams.Credentials
		if credentials == nil {
			logger.Error("credentials == nil")
			return operations2.NewPostLoginInternalServerError().WithPayload(&models.ActionFailure{
				From:    convert.StrPtr(apiName),
				Reason:  convert.StrPtr("参数错误"),
				Success: web.Fail,
			})
		}
		secrete := credentials.Secrete
		if secrete == nil {
			logger.Error("secrete == nil")
			return operations2.NewPostLoginInternalServerError().WithPayload(&models.ActionFailure{
				From:    convert.StrPtr(apiName),
				Reason:  convert.StrPtr("参数错误"),
				Success: web.Fail,
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
		if credentials.Name == nil || credentials.Password == nil {
			logger.Error("credentials.Name == nil || credentials.Password == nil")
			return operations2.NewPostLoginInternalServerError().WithPayload(&models.ActionFailure{
				From:    convert.StrPtr(apiName),
				Reason:  convert.StrPtr("参数错误"),
				Success: web.Fail,
			})
		}
		name := *credentials.Name
		password := *credentials.Password
		if err := UserInfo.CheckUser(apiManager.s.DBM.DB, name, password); err != nil {
			reason := "用户名或密码错误"
			if !errors.Is(err, UserInfo.ErrUserNameOrPassword) {
				logger.Errorf("LoginPostApiHandlerFunc -> UserInfo.CheckUser -> %v", err)
				reason = "内部错误"
			}

			return operations2.NewPostLoginInternalServerError().WithPayload(&models.ActionFailure{
				From:    convert.StrPtr(apiName),
				Reason:  convert.StrPtr(reason),
				Success: web.Fail,
			})
		}
		jwt, err := augeuJwt.NewJwt(augeuJwt.Info{
			Role: 0,
			UserInfo: augeuJwt.UserInfo{
				Name: name,
			},
		})
		if err != nil {
			logger.Errorf("LoginPostApiHandlerFunc -> augeuJwt.NewJwt -> %v", err)
			return operations2.NewPostLoginInternalServerError().WithPayload(&models.ActionFailure{
				From:    convert.StrPtr(apiName),
				Reason:  convert.StrPtr("产生jwt 出错"),
				Success: web.Fail,
			})
		}
		return operations2.NewPostLoginOK().WithPayload(&models.LoginResponse{
			Jwt:     utils.StrP(jwt),
			Success: web.Success,
		})
	}
}
