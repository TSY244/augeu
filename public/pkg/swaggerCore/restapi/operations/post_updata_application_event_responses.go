// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"augeu/public/pkg/swaggerCore/models"
)

// PostUpdataApplicationEventOKCode is the HTTP code returned for type PostUpdataApplicationEventOK
const PostUpdataApplicationEventOKCode int = 200

/*
PostUpdataApplicationEventOK 上传成功

swagger:response postUpdataApplicationEventOK
*/
type PostUpdataApplicationEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewPostUpdataApplicationEventOK creates PostUpdataApplicationEventOK with default headers values
func NewPostUpdataApplicationEventOK() *PostUpdataApplicationEventOK {

	return &PostUpdataApplicationEventOK{}
}

// WithPayload adds the payload to the post updata application event o k response
func (o *PostUpdataApplicationEventOK) WithPayload(payload *models.SuccessResponse) *PostUpdataApplicationEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post updata application event o k response
func (o *PostUpdataApplicationEventOK) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUpdataApplicationEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUpdataApplicationEventBadRequestCode is the HTTP code returned for type PostUpdataApplicationEventBadRequest
const PostUpdataApplicationEventBadRequestCode int = 400

/*
PostUpdataApplicationEventBadRequest 输入参数错误

swagger:response postUpdataApplicationEventBadRequest
*/
type PostUpdataApplicationEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewPostUpdataApplicationEventBadRequest creates PostUpdataApplicationEventBadRequest with default headers values
func NewPostUpdataApplicationEventBadRequest() *PostUpdataApplicationEventBadRequest {

	return &PostUpdataApplicationEventBadRequest{}
}

// WithPayload adds the payload to the post updata application event bad request response
func (o *PostUpdataApplicationEventBadRequest) WithPayload(payload *models.BadRequestError) *PostUpdataApplicationEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post updata application event bad request response
func (o *PostUpdataApplicationEventBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUpdataApplicationEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUpdataApplicationEventForbiddenCode is the HTTP code returned for type PostUpdataApplicationEventForbidden
const PostUpdataApplicationEventForbiddenCode int = 403

/*
PostUpdataApplicationEventForbidden 没有权限

swagger:response postUpdataApplicationEventForbidden
*/
type PostUpdataApplicationEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.UnauthorizedError `json:"body,omitempty"`
}

// NewPostUpdataApplicationEventForbidden creates PostUpdataApplicationEventForbidden with default headers values
func NewPostUpdataApplicationEventForbidden() *PostUpdataApplicationEventForbidden {

	return &PostUpdataApplicationEventForbidden{}
}

// WithPayload adds the payload to the post updata application event forbidden response
func (o *PostUpdataApplicationEventForbidden) WithPayload(payload *models.UnauthorizedError) *PostUpdataApplicationEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post updata application event forbidden response
func (o *PostUpdataApplicationEventForbidden) SetPayload(payload *models.UnauthorizedError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUpdataApplicationEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUpdataApplicationEventInternalServerErrorCode is the HTTP code returned for type PostUpdataApplicationEventInternalServerError
const PostUpdataApplicationEventInternalServerErrorCode int = 500

/*
PostUpdataApplicationEventInternalServerError 内部错误

swagger:response postUpdataApplicationEventInternalServerError
*/
type PostUpdataApplicationEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ActionFailure `json:"body,omitempty"`
}

// NewPostUpdataApplicationEventInternalServerError creates PostUpdataApplicationEventInternalServerError with default headers values
func NewPostUpdataApplicationEventInternalServerError() *PostUpdataApplicationEventInternalServerError {

	return &PostUpdataApplicationEventInternalServerError{}
}

// WithPayload adds the payload to the post updata application event internal server error response
func (o *PostUpdataApplicationEventInternalServerError) WithPayload(payload *models.ActionFailure) *PostUpdataApplicationEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post updata application event internal server error response
func (o *PostUpdataApplicationEventInternalServerError) SetPayload(payload *models.ActionFailure) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUpdataApplicationEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
