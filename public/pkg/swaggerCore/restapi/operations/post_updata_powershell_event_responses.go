// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostUpdataPowershellEventOKCode is the HTTP code returned for type PostUpdataPowershellEventOK
const PostUpdataPowershellEventOKCode int = 200

/*
PostUpdataPowershellEventOK 上传成功

swagger:response postUpdataPowershellEventOK
*/
type PostUpdataPowershellEventOK struct {
}

// NewPostUpdataPowershellEventOK creates PostUpdataPowershellEventOK with default headers values
func NewPostUpdataPowershellEventOK() *PostUpdataPowershellEventOK {

	return &PostUpdataPowershellEventOK{}
}

// WriteResponse to the client
func (o *PostUpdataPowershellEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
