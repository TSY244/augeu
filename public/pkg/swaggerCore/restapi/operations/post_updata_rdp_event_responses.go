// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostUpdataRdpEventOKCode is the HTTP code returned for type PostUpdataRdpEventOK
const PostUpdataRdpEventOKCode int = 200

/*
PostUpdataRdpEventOK 上传成功

swagger:response postUpdataRdpEventOK
*/
type PostUpdataRdpEventOK struct {
}

// NewPostUpdataRdpEventOK creates PostUpdataRdpEventOK with default headers values
func NewPostUpdataRdpEventOK() *PostUpdataRdpEventOK {

	return &PostUpdataRdpEventOK{}
}

// WriteResponse to the client
func (o *PostUpdataRdpEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
