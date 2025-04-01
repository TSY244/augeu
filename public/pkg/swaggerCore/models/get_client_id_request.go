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

// GetClientIDRequest get client Id request
//
// swagger:model GetClientIdRequest
type GetClientIDRequest struct {

	// IP 地址列表
	// Required: true
	IP []string `json:"ip"`

	// 密钥
	// Required: true
	Secret *string `json:"secret"`

	// 系统信息
	// Required: true
	SystemInfo *SystemInfo `json:"system_info"`

	// 唯一标识符
	// Required: true
	UUID *string `json:"uuid"`
}

// Validate validates this get client Id request
func (m *GetClientIDRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemInfo(formats); err != nil {
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

func (m *GetClientIDRequest) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *GetClientIDRequest) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("secret", "body", m.Secret); err != nil {
		return err
	}

	return nil
}

func (m *GetClientIDRequest) validateSystemInfo(formats strfmt.Registry) error {

	if err := validate.Required("system_info", "body", m.SystemInfo); err != nil {
		return err
	}

	if m.SystemInfo != nil {
		if err := m.SystemInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system_info")
			}
			return err
		}
	}

	return nil
}

func (m *GetClientIDRequest) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get client Id request based on the context it is used
func (m *GetClientIDRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSystemInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClientIDRequest) contextValidateSystemInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemInfo != nil {

		if err := m.SystemInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetClientIDRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientIDRequest) UnmarshalBinary(b []byte) error {
	var res GetClientIDRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
