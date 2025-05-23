// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WeiBuTag wei bu tag
//
// swagger:model WeiBuTag
type WeiBuTag struct {

	// s
	// Required: true
	S []string `json:"s"`

	// x
	// Required: true
	X []string `json:"x"`
}

// Validate validates this wei bu tag
func (m *WeiBuTag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateX(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WeiBuTag) validateS(formats strfmt.Registry) error {

	if err := validate.Required("s", "body", m.S); err != nil {
		return err
	}

	return nil
}

func (m *WeiBuTag) validateX(formats strfmt.Registry) error {

	if err := validate.Required("x", "body", m.X); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wei bu tag based on context it is used
func (m *WeiBuTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WeiBuTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WeiBuTag) UnmarshalBinary(b []byte) error {
	var res WeiBuTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
