package middleware

import (
	"augeu/backEnd/internal/pkg/DBMnager/HostInfo"
	"augeu/backEnd/internal/pkg/server"
	"augeu/public/pkg/augeuJwt"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	operations2 "augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

// parseJwtInfo 提取公共JWT解析逻辑
func parseJwtInfo(r *http.Request) (*augeuJwt.Info, error) {
	tokenStr := r.Header.Get(HeadKey)
	if tokenStr == "" {
		return nil, fmt.Errorf("missing %s header", HeadKey)
	}

	jwtInfo, err := augeuJwt.ParseJwt(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("parse jwt failed: %w", err)
	}
	return &jwtInfo, nil
}

// GetRole 从请求中获取用户角色
func GetRole(r *http.Request) (int, error) {
	jwtInfo, err := parseJwtInfo(r)
	if err != nil {
		return ErrorRole, err
	}
	return jwtInfo.Role, nil
}

// GetUserName 从请求中获取用户名
func GetUserName(r *http.Request) (string, error) {
	jwtInfo, err := parseJwtInfo(r)
	if err != nil {
		return "", err
	}
	return jwtInfo.UserInfo.Name, nil
}

// GetClientId 从请求中获取客户端ID
func GetClientId(r *http.Request) (string, error) {
	jwtInfo, err := parseJwtInfo(r)
	if err != nil {
		return "", err
	}
	return jwtInfo.ClientInfo.ClientId, nil
}

func GetClientUuid(r *http.Request) (string, error) {
	jwtInfo, err := parseJwtInfo(r)
	if err != nil {
		return "", err
	}
	return jwtInfo.ClientInfo.Uuid, nil
}

func GetInfo(r *http.Request) (*augeuJwt.Info, error) {
	jwtInfo, err := parseJwtInfo(r)
	if err != nil {
		return nil, err
	}
	return jwtInfo, nil
}

func CheckAgentRole(r *http.Request, s *server.Server) middleware.Responder {
	info, err := GetInfo(r)
	if err != nil {
		logger.Errorf("UploadLoginEventApiHandlerFunc -> %v", err)
		return operations2.NewPostGetClientIDUnauthorized().WithPayload(&models.UnauthorizedError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventForbiddenCode)),
			Message: convert.StrPtr("jwt is invalid"),
		})
	}
	if info.Role != RoleAgent {
		logger.Errorf("UploadLoginEventApiHandlerFunc -> info.Role != middleware2.RoleAgent")
		return operations2.NewPostGetClientIDUnauthorized().WithPayload(&models.UnauthorizedError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventForbiddenCode)),
			Message: convert.StrPtr("jwt is invalid"),
		})
	}
	clientId := info.ClientInfo.ClientId
	if clientId == "" {
		logger.Errorf("UploadLoginEventApiHandlerFunc -> clientid is empty")
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("clientid is empty"),
		})
	}
	if ret := s.CheckClientId(clientId); !ret {
		// 说明没有这个clientId 或者已经断开
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("clientid is invalid"),
		})
	}

	uuid := info.ClientInfo.Uuid
	if uuid == "" {
		logger.Errorf("UploadLoginEventApiHandlerFunc -> uuid is empty")
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("uuid is empty"),
		})
	}
	// 检测 uuid
	sgentInfo, err := HostInfo.GetAgentInfoByClientId(s.DBM.DB, clientId)
	if err != nil {
		logger.Errorf("UploadLoginEventApiHandlerFunc -> GetAgentInfoByClientId err:%v", err)
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("uuid is invalid"),
		})
	}
	if sgentInfo.UUID != uuid {
		logger.Errorf("UploadLoginEventApiHandlerFunc -> uuid is invalid")
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("uuid is invalid"),
		})
	}

	return nil
}
