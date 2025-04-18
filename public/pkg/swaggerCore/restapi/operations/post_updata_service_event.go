// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostUpdataServiceEventHandlerFunc turns a function with the right signature into a post updata service event handler
type PostUpdataServiceEventHandlerFunc func(PostUpdataServiceEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUpdataServiceEventHandlerFunc) Handle(params PostUpdataServiceEventParams) middleware.Responder {
	return fn(params)
}

// PostUpdataServiceEventHandler interface for that can handle valid post updata service event params
type PostUpdataServiceEventHandler interface {
	Handle(PostUpdataServiceEventParams) middleware.Responder
}

// NewPostUpdataServiceEvent creates a new http.Handler for the post updata service event operation
func NewPostUpdataServiceEvent(ctx *middleware.Context, handler PostUpdataServiceEventHandler) *PostUpdataServiceEvent {
	return &PostUpdataServiceEvent{Context: ctx, Handler: handler}
}

/*
	PostUpdataServiceEvent swagger:route POST /updata/serviceEvent postUpdataServiceEvent

上传服务事件
*/
type PostUpdataServiceEvent struct {
	Context *middleware.Context
	Handler PostUpdataServiceEventHandler
}

func (o *PostUpdataServiceEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostUpdataServiceEventParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
