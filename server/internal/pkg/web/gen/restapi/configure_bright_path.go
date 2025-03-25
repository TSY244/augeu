// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"augeu/server/internal/pkg/web/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name BrightPath --spec ../../../../../swagger.yaml --principal models.Principle

func configureFlags(api *operations.BrightPathAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BrightPathAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.PostUploadAvatarMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.PostVolunteerUploadIDBackMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.PostVolunteerUploadIDFrontMaxParseMemory = 32 << 20

	if api.GetAdminViewRealTimeVolunteerHelpInfoHandler == nil {
		api.GetAdminViewRealTimeVolunteerHelpInfoHandler = operations.GetAdminViewRealTimeVolunteerHelpInfoHandlerFunc(func(params operations.GetAdminViewRealTimeVolunteerHelpInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetAdminViewRealTimeVolunteerHelpInfo has not yet been implemented")
		})
	}
	if api.GetVersionHandler == nil {
		api.GetVersionHandler = operations.GetVersionHandlerFunc(func(params operations.GetVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetVersion has not yet been implemented")
		})
	}
	if api.GetVolunteerViewOwnMissionHistoryHandler == nil {
		api.GetVolunteerViewOwnMissionHistoryHandler = operations.GetVolunteerViewOwnMissionHistoryHandlerFunc(func(params operations.GetVolunteerViewOwnMissionHistoryParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetVolunteerViewOwnMissionHistory has not yet been implemented")
		})
	}
	if api.PostAdminApproveVolunteerRegistrationHandler == nil {
		api.PostAdminApproveVolunteerRegistrationHandler = operations.PostAdminApproveVolunteerRegistrationHandlerFunc(func(params operations.PostAdminApproveVolunteerRegistrationParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostAdminApproveVolunteerRegistration has not yet been implemented")
		})
	}
	if api.PostAdminDeleteVolunteerHandler == nil {
		api.PostAdminDeleteVolunteerHandler = operations.PostAdminDeleteVolunteerHandlerFunc(func(params operations.PostAdminDeleteVolunteerParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostAdminDeleteVolunteer has not yet been implemented")
		})
	}
	if api.PostAdminResetPasswordHandler == nil {
		api.PostAdminResetPasswordHandler = operations.PostAdminResetPasswordHandlerFunc(func(params operations.PostAdminResetPasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostAdminResetPassword has not yet been implemented")
		})
	}
	if api.PostAdminResetVolunteerAvatarToDefaultHandler == nil {
		api.PostAdminResetVolunteerAvatarToDefaultHandler = operations.PostAdminResetVolunteerAvatarToDefaultHandlerFunc(func(params operations.PostAdminResetVolunteerAvatarToDefaultParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostAdminResetVolunteerAvatarToDefault has not yet been implemented")
		})
	}
	if api.PostAdminUpdateVolunteerInfoHandler == nil {
		api.PostAdminUpdateVolunteerInfoHandler = operations.PostAdminUpdateVolunteerInfoHandlerFunc(func(params operations.PostAdminUpdateVolunteerInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostAdminUpdateVolunteerInfo has not yet been implemented")
		})
	}
	if api.PostLoginHandler == nil {
		api.PostLoginHandler = operations.PostLoginHandlerFunc(func(params operations.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostLogin has not yet been implemented")
		})
	}
	if api.PostLogoutHandler == nil {
		api.PostLogoutHandler = operations.PostLogoutHandlerFunc(func(params operations.PostLogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostLogout has not yet been implemented")
		})
	}
	if api.PostResetPasswordHandler == nil {
		api.PostResetPasswordHandler = operations.PostResetPasswordHandlerFunc(func(params operations.PostResetPasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostResetPassword has not yet been implemented")
		})
	}
	if api.PostUpdateAvatarHandler == nil {
		api.PostUpdateAvatarHandler = operations.PostUpdateAvatarHandlerFunc(func(params operations.PostUpdateAvatarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUpdateAvatar has not yet been implemented")
		})
	}
	if api.PostUploadAvatarHandler == nil {
		api.PostUploadAvatarHandler = operations.PostUploadAvatarHandlerFunc(func(params operations.PostUploadAvatarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUploadAvatar has not yet been implemented")
		})
	}
	if api.PostVolunteerAcceptMissionHandler == nil {
		api.PostVolunteerAcceptMissionHandler = operations.PostVolunteerAcceptMissionHandlerFunc(func(params operations.PostVolunteerAcceptMissionParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostVolunteerAcceptMission has not yet been implemented")
		})
	}
	if api.PostVolunteerConnectGuideCaneCameraHandler == nil {
		api.PostVolunteerConnectGuideCaneCameraHandler = operations.PostVolunteerConnectGuideCaneCameraHandlerFunc(func(params operations.PostVolunteerConnectGuideCaneCameraParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostVolunteerConnectGuideCaneCamera has not yet been implemented")
		})
	}
	if api.PostVolunteerRegisterHandler == nil {
		api.PostVolunteerRegisterHandler = operations.PostVolunteerRegisterHandlerFunc(func(params operations.PostVolunteerRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostVolunteerRegister has not yet been implemented")
		})
	}
	if api.PostVolunteerUploadIDBackHandler == nil {
		api.PostVolunteerUploadIDBackHandler = operations.PostVolunteerUploadIDBackHandlerFunc(func(params operations.PostVolunteerUploadIDBackParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostVolunteerUploadIDBack has not yet been implemented")
		})
	}
	if api.PostVolunteerUploadIDFrontHandler == nil {
		api.PostVolunteerUploadIDFrontHandler = operations.PostVolunteerUploadIDFrontHandlerFunc(func(params operations.PostVolunteerUploadIDFrontParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostVolunteerUploadIDFront has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
