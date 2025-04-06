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

// ClientInfo client info
//
// swagger:model ClientInfo
type ClientInfo struct {

	// 系统信息
	// Required: true
	SystemInfo *SystemInfo `json:"SystemInfo"`

	// IP 地址列表
	// Required: true
	IP []string `json:"ip"`

	// 唯一标识符
	// Required: true
	UUID *string `json:"uuid"`
}

// Validate validates this client info
func (m *ClientInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSystemInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientInfo) validateSystemInfo(formats strfmt.Registry) error {

	if err := validate.Required("SystemInfo", "body", m.SystemInfo); err != nil {
		return err
	}

	if m.SystemInfo != nil {
		if err := m.SystemInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SystemInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SystemInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ClientInfo) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *ClientInfo) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this client info based on the context it is used
func (m *ClientInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSystemInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientInfo) contextValidateSystemInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemInfo != nil {

		if err := m.SystemInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SystemInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SystemInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientInfo) UnmarshalBinary(b []byte) error {
	var res ClientInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
