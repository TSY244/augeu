package middleware

import (
	"augeu/backEnd/internal/pkg/DBMnager/HostInfo"
	"augeu/backEnd/internal/pkg/DBMnager/UserInfo"
	"augeu/backEnd/internal/pkg/server"
	"augeu/backEnd/internal/utils/consts/web"
	"augeu/public/pkg/augeuJwt"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	operations2 "augeu/public/pkg/swaggerCore/restapi/operations"
	"augeu/public/util/convert"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"strings"
)

// parseJwtInfo 提取公共JWT解析逻辑
func parseJwtInfo(r *http.Request) (*augeuJwt.Info, error) {
	tokenStr := r.Header.Get(HeadKey)
	if tokenStr == "" {
		return nil, fmt.Errorf("missing %s header", HeadKey)
	}
	// 先去除 Bearer
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

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
	_, err = HostInfo.GetAgentInfoByUuid(s.DBM.DB, uuid)
	if err != nil {
		logger.Errorf("UploadLoginEventApiHandlerFunc -> GetAgentInfoByUuid err:%v", err)
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("uuid is invalid"),
		})
	}

	return nil
}

func CheckUserRole(r *http.Request, s *server.Server) middleware.Responder {
	info, err := GetInfo(r)
	if err != nil {
		logger.Errorf("CheckUserRole -> GetRole err:%v", err)
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("get role is error"),
		})
	}

	if info.Role != RoleUser {
		logger.Errorf("CheckUserRole -> role:%v", info.Role)
	}
	userName, err := GetUserName(r)
	if err != nil {
		logger.Errorf("CheckUserRole -> GetUserName err:%v", err)
		return operations2.NewPostLoginInternalServerError().WithPayload(&models.ActionFailure{
			From:    convert.StrPtr("check jwt"),
			Reason:  convert.StrPtr("get user name is error"),
			Success: web.Fail,
		})
	}
	user, err := UserInfo.GetUserByName(s.DBM.DB, userName)
	if err != nil {
		logger.Errorf("CheckUserRole -> GetUserByName err:%v", err)
		return operations2.NewPostLoginInternalServerError().WithPayload(&models.ActionFailure{
			From:    convert.StrPtr("check jwt"),
			Reason:  convert.StrPtr("get user name is error"),
			Success: web.Fail,
		})
	}
	if user.UserName != info.UserInfo.Name {
		logger.Errorf("CheckUserRole -> userName is invalid")
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("userName is invalid"),
		})
	}
	return nil
}

// 存在绕过风险，不意见使用
func CheckJwt(r *http.Request, s *server.Server) middleware.Responder {
	info, err := GetInfo(r)
	if err != nil {
		logger.Errorf("CheckJwt -> GetInfo error: %v", err)
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("jwt is invalid"),
		})
	}
	if info.Role == RoleUser {
		return CheckUserRole(r, s)
	}
	return CheckAgentRole(r, s)

}

// CheckAgentJwt 检查agent的jwt，不用检查clientId
//
// 注意：
//   - 这个检测函数只给getRule 使用
func CheckAgentJwt(r *http.Request, s *server.Server) middleware.Responder {
	info, err := GetInfo(r)
	if err != nil {
		logger.Errorf("CheckJwt -> GetInfo error: %v", err)
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("jwt is invalid"),
		})
	}
	if info.Role != RoleAgent {
		logger.Errorf("CheckAgentJwt -> info.Role != middleware2.RoleAgent")
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("jwt is invalid"),
		})
	}
	uuid := info.ClientInfo.Uuid
	if uuid == "" {
		logger.Errorf("CheckAgentJwt -> uuid is empty")
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("uuid is empty"),
		})
	}
	agent, err := HostInfo.GetAgentInfoByUuid(s.DBM.DB, uuid)
	if err != nil {
		logger.Errorf("CheckAgentJwt -> GetAgentInfoByUuid err:%v", err)
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("uuid is invalid"),
		})
	}
	if agent.ClientID != info.ClientInfo.ClientId {
		logger.Errorf("CheckAgentJwt -> agent.ClientID != info.ClientInfo.ClientId")
		return operations2.NewPostUploadLoginEventBadRequest().WithPayload(&models.BadRequestError{
			Code:    convert.Int64P(int64(operations2.PostUploadLoginEventBadRequestCode)),
			Message: convert.StrPtr("uuid is invalid"),
		})
	}
	return nil
}
