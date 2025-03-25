// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"augeu/server/internal/pkg/web/gen/models"
)

// GetVersionOKCode is the HTTP code returned for type GetVersionOK
const GetVersionOKCode int = 200

/*
GetVersionOK 返回 BrightPath Api 版本号

swagger:response getVersionOK
*/
type GetVersionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Version `json:"body,omitempty"`
}

// NewGetVersionOK creates GetVersionOK with default headers values
func NewGetVersionOK() *GetVersionOK {

	return &GetVersionOK{}
}

// WithPayload adds the payload to the get version o k response
func (o *GetVersionOK) WithPayload(payload *models.Version) *GetVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get version o k response
func (o *GetVersionOK) SetPayload(payload *models.Version) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
