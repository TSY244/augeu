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

// EventRDPLogon event r d p logon
//
// swagger:model EventRDPLogon
type EventRDPLogon struct {

	// address
	Address string `json:"Address,omitempty"`

	// create at
	// Format: date-time
	CreateAt strfmt.DateTime `json:"CreateAt,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// domain
	Domain string `json:"Domain,omitempty"`

	// event ID
	EventID string `json:"EventID,omitempty"`

	// ID
	ID int64 `json:"ID,omitempty"`

	// login name
	LoginName string `json:"LoginName,omitempty"`

	// UUID
	UUID string `json:"UUID,omitempty"`
}

// Validate validates this event r d p logon
func (m *EventRDPLogon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventRDPLogon) validateCreateAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateAt) { // not required
		return nil
	}

	if err := validate.FormatOf("CreateAt", "body", "date-time", m.CreateAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event r d p logon based on context it is used
func (m *EventRDPLogon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventRDPLogon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventRDPLogon) UnmarshalBinary(b []byte) error {
	var res EventRDPLogon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
