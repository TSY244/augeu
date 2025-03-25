// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PostVolunteerUploadIDFrontMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var PostVolunteerUploadIDFrontMaxParseMemory int64 = 32 << 20

// NewPostVolunteerUploadIDFrontParams creates a new PostVolunteerUploadIDFrontParams object
//
// There are no default values defined in the spec.
func NewPostVolunteerUploadIDFrontParams() PostVolunteerUploadIDFrontParams {

	return PostVolunteerUploadIDFrontParams{}
}

// PostVolunteerUploadIDFrontParams contains all the bound params for the post volunteer upload ID front operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostVolunteerUploadIDFront
type PostVolunteerUploadIDFrontParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*上传身份证正面
	  Required: true
	  In: formData
	*/
	File io.ReadCloser
	/*志愿者id
	  Required: true
	  In: formData
	*/
	VolunteerID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostVolunteerUploadIDFrontParams() beforehand.
func (o *PostVolunteerUploadIDFrontParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(PostVolunteerUploadIDFrontMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else if err := o.bindFile(file, fileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.File = &runtime.File{Data: file, Header: fileHeader}
	}

	fdVolunteerID, fdhkVolunteerID, _ := fds.GetOK("volunteerId")
	if err := o.bindVolunteerID(fdVolunteerID, fdhkVolunteerID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFile binds file parameter File.
//
// The only supported validations on files are MinLength and MaxLength
func (o *PostVolunteerUploadIDFrontParams) bindFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindVolunteerID binds and validates parameter VolunteerID from formData.
func (o *PostVolunteerUploadIDFrontParams) bindVolunteerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("volunteerId", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("volunteerId", "formData", raw); err != nil {
		return err
	}
	o.VolunteerID = raw

	return nil
}
