package api

import (
	middleware2 "augeu/backEnd/internal/pkg/web/middleware"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	"augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"github.com/go-openapi/runtime/middleware"
)

func (apiManager *ApiManager) GetRulesGetHandlerFunc() operations.GetGetRulesHandlerFunc {
	return func(params operations.GetGetRulesParams) middleware.Responder {
		if resp := middleware2.CheckAgentRole(params.HTTPRequest, apiManager.s); resp != nil {
			return resp
		}
		rule, err := apiManager.s.GetRule()
		if err != nil {
			logger.Errorf("GetRulesGetApiHandlerFunc -> GetRule error: %v", err)
			return operations.NewGetGetRulesBadRequest().WithPayload(&models.BadRequestError{
				Code:    convert.Int64P(int64(operations.GetGetRulesBadRequestCode)),
				Message: convert.StrPtr("GetRule error"),
			})
		}
		return operations.NewGetGetRulesOK().WithPayload(&models.GetRulesResponse{
			Data:         convert.StrPtr(rule),
			ResponseCode: convert.Int64P(int64(operations.GetGetRulesOKCode)),
		})
	}
}
