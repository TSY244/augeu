// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewAugeuAPI creates a new Augeu instance
func NewAugeuAPI(spec *loads.Document) *AugeuAPI {
	return &AugeuAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		GetGetClientsHandler: GetGetClientsHandlerFunc(func(params GetGetClientsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetGetClients has not yet been implemented")
		}),
		GetVersionHandler: GetVersionHandlerFunc(func(params GetVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation GetVersion has not yet been implemented")
		}),
		PostGetApplicationEventHandler: PostGetApplicationEventHandlerFunc(func(params PostGetApplicationEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostGetApplicationEvent has not yet been implemented")
		}),
		PostGetClientIDHandler: PostGetClientIDHandlerFunc(func(params PostGetClientIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PostGetClientID has not yet been implemented")
		}),
		PostGetLoginEventHandler: PostGetLoginEventHandlerFunc(func(params PostGetLoginEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostGetLoginEvent has not yet been implemented")
		}),
		PostGetPowershellEventHandler: PostGetPowershellEventHandlerFunc(func(params PostGetPowershellEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostGetPowershellEvent has not yet been implemented")
		}),
		PostGetProcessEventHandler: PostGetProcessEventHandlerFunc(func(params PostGetProcessEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostGetProcessEvent has not yet been implemented")
		}),
		PostGetRdpEventHandler: PostGetRdpEventHandlerFunc(func(params PostGetRdpEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostGetRdpEvent has not yet been implemented")
		}),
		PostGetSecurityEventHandler: PostGetSecurityEventHandlerFunc(func(params PostGetSecurityEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostGetSecurityEvent has not yet been implemented")
		}),
		PostGetServiceEventHandler: PostGetServiceEventHandlerFunc(func(params PostGetServiceEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostGetServiceEvent has not yet been implemented")
		}),
		PostGetSystemEventHandler: PostGetSystemEventHandlerFunc(func(params PostGetSystemEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostGetSystemEvent has not yet been implemented")
		}),
		PostLoginHandler: PostLoginHandlerFunc(func(params PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation PostLogin has not yet been implemented")
		}),
		PostRegisterHandler: PostRegisterHandlerFunc(func(params PostRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation PostRegister has not yet been implemented")
		}),
		PostUpdataApplicationEventHandler: PostUpdataApplicationEventHandlerFunc(func(params PostUpdataApplicationEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUpdataApplicationEvent has not yet been implemented")
		}),
		PostUpdataPowershellEventHandler: PostUpdataPowershellEventHandlerFunc(func(params PostUpdataPowershellEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUpdataPowershellEvent has not yet been implemented")
		}),
		PostUpdataProcessEventHandler: PostUpdataProcessEventHandlerFunc(func(params PostUpdataProcessEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUpdataProcessEvent has not yet been implemented")
		}),
		PostUpdataRdpEventHandler: PostUpdataRdpEventHandlerFunc(func(params PostUpdataRdpEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUpdataRdpEvent has not yet been implemented")
		}),
		PostUpdataSecurityEventHandler: PostUpdataSecurityEventHandlerFunc(func(params PostUpdataSecurityEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUpdataSecurityEvent has not yet been implemented")
		}),
		PostUpdataServiceEventHandler: PostUpdataServiceEventHandlerFunc(func(params PostUpdataServiceEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUpdataServiceEvent has not yet been implemented")
		}),
		PostUpdataSystemEventHandler: PostUpdataSystemEventHandlerFunc(func(params PostUpdataSystemEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUpdataSystemEvent has not yet been implemented")
		}),
		PostUploadLoginEventHandler: PostUploadLoginEventHandlerFunc(func(params PostUploadLoginEventParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUploadLoginEvent has not yet been implemented")
		}),
	}
}

/*AugeuAPI Augeu api */
type AugeuAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// GetGetClientsHandler sets the operation handler for the get get clients operation
	GetGetClientsHandler GetGetClientsHandler
	// GetVersionHandler sets the operation handler for the get version operation
	GetVersionHandler GetVersionHandler
	// PostGetApplicationEventHandler sets the operation handler for the post get application event operation
	PostGetApplicationEventHandler PostGetApplicationEventHandler
	// PostGetClientIDHandler sets the operation handler for the post get client ID operation
	PostGetClientIDHandler PostGetClientIDHandler
	// PostGetLoginEventHandler sets the operation handler for the post get login event operation
	PostGetLoginEventHandler PostGetLoginEventHandler
	// PostGetPowershellEventHandler sets the operation handler for the post get powershell event operation
	PostGetPowershellEventHandler PostGetPowershellEventHandler
	// PostGetProcessEventHandler sets the operation handler for the post get process event operation
	PostGetProcessEventHandler PostGetProcessEventHandler
	// PostGetRdpEventHandler sets the operation handler for the post get rdp event operation
	PostGetRdpEventHandler PostGetRdpEventHandler
	// PostGetSecurityEventHandler sets the operation handler for the post get security event operation
	PostGetSecurityEventHandler PostGetSecurityEventHandler
	// PostGetServiceEventHandler sets the operation handler for the post get service event operation
	PostGetServiceEventHandler PostGetServiceEventHandler
	// PostGetSystemEventHandler sets the operation handler for the post get system event operation
	PostGetSystemEventHandler PostGetSystemEventHandler
	// PostLoginHandler sets the operation handler for the post login operation
	PostLoginHandler PostLoginHandler
	// PostRegisterHandler sets the operation handler for the post register operation
	PostRegisterHandler PostRegisterHandler
	// PostUpdataApplicationEventHandler sets the operation handler for the post updata application event operation
	PostUpdataApplicationEventHandler PostUpdataApplicationEventHandler
	// PostUpdataPowershellEventHandler sets the operation handler for the post updata powershell event operation
	PostUpdataPowershellEventHandler PostUpdataPowershellEventHandler
	// PostUpdataProcessEventHandler sets the operation handler for the post updata process event operation
	PostUpdataProcessEventHandler PostUpdataProcessEventHandler
	// PostUpdataRdpEventHandler sets the operation handler for the post updata rdp event operation
	PostUpdataRdpEventHandler PostUpdataRdpEventHandler
	// PostUpdataSecurityEventHandler sets the operation handler for the post updata security event operation
	PostUpdataSecurityEventHandler PostUpdataSecurityEventHandler
	// PostUpdataServiceEventHandler sets the operation handler for the post updata service event operation
	PostUpdataServiceEventHandler PostUpdataServiceEventHandler
	// PostUpdataSystemEventHandler sets the operation handler for the post updata system event operation
	PostUpdataSystemEventHandler PostUpdataSystemEventHandler
	// PostUploadLoginEventHandler sets the operation handler for the post upload login event operation
	PostUploadLoginEventHandler PostUploadLoginEventHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *AugeuAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *AugeuAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *AugeuAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AugeuAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AugeuAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AugeuAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AugeuAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AugeuAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AugeuAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AugeuAPI
func (o *AugeuAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GetGetClientsHandler == nil {
		unregistered = append(unregistered, "GetGetClientsHandler")
	}
	if o.GetVersionHandler == nil {
		unregistered = append(unregistered, "GetVersionHandler")
	}
	if o.PostGetApplicationEventHandler == nil {
		unregistered = append(unregistered, "PostGetApplicationEventHandler")
	}
	if o.PostGetClientIDHandler == nil {
		unregistered = append(unregistered, "PostGetClientIDHandler")
	}
	if o.PostGetLoginEventHandler == nil {
		unregistered = append(unregistered, "PostGetLoginEventHandler")
	}
	if o.PostGetPowershellEventHandler == nil {
		unregistered = append(unregistered, "PostGetPowershellEventHandler")
	}
	if o.PostGetProcessEventHandler == nil {
		unregistered = append(unregistered, "PostGetProcessEventHandler")
	}
	if o.PostGetRdpEventHandler == nil {
		unregistered = append(unregistered, "PostGetRdpEventHandler")
	}
	if o.PostGetSecurityEventHandler == nil {
		unregistered = append(unregistered, "PostGetSecurityEventHandler")
	}
	if o.PostGetServiceEventHandler == nil {
		unregistered = append(unregistered, "PostGetServiceEventHandler")
	}
	if o.PostGetSystemEventHandler == nil {
		unregistered = append(unregistered, "PostGetSystemEventHandler")
	}
	if o.PostLoginHandler == nil {
		unregistered = append(unregistered, "PostLoginHandler")
	}
	if o.PostRegisterHandler == nil {
		unregistered = append(unregistered, "PostRegisterHandler")
	}
	if o.PostUpdataApplicationEventHandler == nil {
		unregistered = append(unregistered, "PostUpdataApplicationEventHandler")
	}
	if o.PostUpdataPowershellEventHandler == nil {
		unregistered = append(unregistered, "PostUpdataPowershellEventHandler")
	}
	if o.PostUpdataProcessEventHandler == nil {
		unregistered = append(unregistered, "PostUpdataProcessEventHandler")
	}
	if o.PostUpdataRdpEventHandler == nil {
		unregistered = append(unregistered, "PostUpdataRdpEventHandler")
	}
	if o.PostUpdataSecurityEventHandler == nil {
		unregistered = append(unregistered, "PostUpdataSecurityEventHandler")
	}
	if o.PostUpdataServiceEventHandler == nil {
		unregistered = append(unregistered, "PostUpdataServiceEventHandler")
	}
	if o.PostUpdataSystemEventHandler == nil {
		unregistered = append(unregistered, "PostUpdataSystemEventHandler")
	}
	if o.PostUploadLoginEventHandler == nil {
		unregistered = append(unregistered, "PostUploadLoginEventHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AugeuAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AugeuAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *AugeuAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *AugeuAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *AugeuAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AugeuAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the augeu API
func (o *AugeuAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AugeuAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/get/clients"] = NewGetGetClients(o.context, o.GetGetClientsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/version"] = NewGetVersion(o.context, o.GetVersionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/get/applicationEvent"] = NewPostGetApplicationEvent(o.context, o.PostGetApplicationEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/getClientId"] = NewPostGetClientID(o.context, o.PostGetClientIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/get/loginEvent"] = NewPostGetLoginEvent(o.context, o.PostGetLoginEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/get/powershellEvent"] = NewPostGetPowershellEvent(o.context, o.PostGetPowershellEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/get/processEvent"] = NewPostGetProcessEvent(o.context, o.PostGetProcessEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/get/rdpEvent"] = NewPostGetRdpEvent(o.context, o.PostGetRdpEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/get/securityEvent"] = NewPostGetSecurityEvent(o.context, o.PostGetSecurityEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/get/serviceEvent"] = NewPostGetServiceEvent(o.context, o.PostGetServiceEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/get/systemEvent"] = NewPostGetSystemEvent(o.context, o.PostGetSystemEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = NewPostLogin(o.context, o.PostLoginHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/register"] = NewPostRegister(o.context, o.PostRegisterHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/updata/applicationEvent"] = NewPostUpdataApplicationEvent(o.context, o.PostUpdataApplicationEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/updata/powershellEvent"] = NewPostUpdataPowershellEvent(o.context, o.PostUpdataPowershellEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/updata/processEvent"] = NewPostUpdataProcessEvent(o.context, o.PostUpdataProcessEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/updata/rdpEvent"] = NewPostUpdataRdpEvent(o.context, o.PostUpdataRdpEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/updata/securityEvent"] = NewPostUpdataSecurityEvent(o.context, o.PostUpdataSecurityEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/updata/serviceEvent"] = NewPostUpdataServiceEvent(o.context, o.PostUpdataServiceEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/updata/systemEvent"] = NewPostUpdataSystemEvent(o.context, o.PostUpdataSystemEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/upload/loginEvent"] = NewPostUploadLoginEvent(o.context, o.PostUploadLoginEventHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AugeuAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *AugeuAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *AugeuAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *AugeuAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *AugeuAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
