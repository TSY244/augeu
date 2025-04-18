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

// RegisterRequest register request
//
// swagger:model RegisterRequest
type RegisterRequest struct {

	// 密码
	// Required: true
	PassWord *string `json:"passWord"`

	// server 提供的token
	// Required: true
	Secrete *string `json:"secrete"`

	// 用户名
	// Required: true
	UserName *string `json:"userName"`
}

// Validate validates this register request
func (m *RegisterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassWord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecrete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisterRequest) validatePassWord(formats strfmt.Registry) error {

	if err := validate.Required("passWord", "body", m.PassWord); err != nil {
		return err
	}

	return nil
}

func (m *RegisterRequest) validateSecrete(formats strfmt.Registry) error {

	if err := validate.Required("secrete", "body", m.Secrete); err != nil {
		return err
	}

	return nil
}

func (m *RegisterRequest) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this register request based on context it is used
func (m *RegisterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegisterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterRequest) UnmarshalBinary(b []byte) error {
	var res RegisterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
