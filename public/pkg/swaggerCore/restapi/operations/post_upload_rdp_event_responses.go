// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"augeu/public/pkg/swaggerCore/models"
)

// PostUploadRdpEventOKCode is the HTTP code returned for type PostUploadRdpEventOK
const PostUploadRdpEventOKCode int = 200

/*
PostUploadRdpEventOK 上传成功

swagger:response postUploadRdpEventOK
*/
type PostUploadRdpEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewPostUploadRdpEventOK creates PostUploadRdpEventOK with default headers values
func NewPostUploadRdpEventOK() *PostUploadRdpEventOK {

	return &PostUploadRdpEventOK{}
}

// WithPayload adds the payload to the post upload rdp event o k response
func (o *PostUploadRdpEventOK) WithPayload(payload *models.SuccessResponse) *PostUploadRdpEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post upload rdp event o k response
func (o *PostUploadRdpEventOK) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUploadRdpEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUploadRdpEventBadRequestCode is the HTTP code returned for type PostUploadRdpEventBadRequest
const PostUploadRdpEventBadRequestCode int = 400

/*
PostUploadRdpEventBadRequest 输入参数错误

swagger:response postUploadRdpEventBadRequest
*/
type PostUploadRdpEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequestError `json:"body,omitempty"`
}

// NewPostUploadRdpEventBadRequest creates PostUploadRdpEventBadRequest with default headers values
func NewPostUploadRdpEventBadRequest() *PostUploadRdpEventBadRequest {

	return &PostUploadRdpEventBadRequest{}
}

// WithPayload adds the payload to the post upload rdp event bad request response
func (o *PostUploadRdpEventBadRequest) WithPayload(payload *models.BadRequestError) *PostUploadRdpEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post upload rdp event bad request response
func (o *PostUploadRdpEventBadRequest) SetPayload(payload *models.BadRequestError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUploadRdpEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUploadRdpEventForbiddenCode is the HTTP code returned for type PostUploadRdpEventForbidden
const PostUploadRdpEventForbiddenCode int = 403

/*
PostUploadRdpEventForbidden 没有权限

swagger:response postUploadRdpEventForbidden
*/
type PostUploadRdpEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.UnauthorizedError `json:"body,omitempty"`
}

// NewPostUploadRdpEventForbidden creates PostUploadRdpEventForbidden with default headers values
func NewPostUploadRdpEventForbidden() *PostUploadRdpEventForbidden {

	return &PostUploadRdpEventForbidden{}
}

// WithPayload adds the payload to the post upload rdp event forbidden response
func (o *PostUploadRdpEventForbidden) WithPayload(payload *models.UnauthorizedError) *PostUploadRdpEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post upload rdp event forbidden response
func (o *PostUploadRdpEventForbidden) SetPayload(payload *models.UnauthorizedError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUploadRdpEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUploadRdpEventInternalServerErrorCode is the HTTP code returned for type PostUploadRdpEventInternalServerError
const PostUploadRdpEventInternalServerErrorCode int = 500

/*
PostUploadRdpEventInternalServerError 内部错误

swagger:response postUploadRdpEventInternalServerError
*/
type PostUploadRdpEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ActionFailure `json:"body,omitempty"`
}

// NewPostUploadRdpEventInternalServerError creates PostUploadRdpEventInternalServerError with default headers values
func NewPostUploadRdpEventInternalServerError() *PostUploadRdpEventInternalServerError {

	return &PostUploadRdpEventInternalServerError{}
}

// WithPayload adds the payload to the post upload rdp event internal server error response
func (o *PostUploadRdpEventInternalServerError) WithPayload(payload *models.ActionFailure) *PostUploadRdpEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post upload rdp event internal server error response
func (o *PostUploadRdpEventInternalServerError) SetPayload(payload *models.ActionFailure) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUploadRdpEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
