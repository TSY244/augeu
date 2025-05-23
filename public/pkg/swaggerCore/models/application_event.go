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

// ApplicationEvent application event
//
// swagger:model ApplicationEvent
type ApplicationEvent struct {

	// create time
	// Format: date-time
	CreateTime strfmt.DateTime `json:"CreateTime,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// event ID
	EventID string `json:"EventID,omitempty"`

	// ID
	ID int64 `json:"ID,omitempty"`

	// level display name
	LevelDisplayName string `json:"LevelDisplayName,omitempty"`

	// UUID
	UUID string `json:"UUID,omitempty"`
}

// Validate validates this application event
func (m *ApplicationEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationEvent) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreateTime", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this application event based on context it is used
func (m *ApplicationEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationEvent) UnmarshalBinary(b []byte) error {
	var res ApplicationEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
