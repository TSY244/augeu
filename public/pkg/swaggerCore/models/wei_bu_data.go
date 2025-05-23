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

// WeiBuData wei bu data
//
// swagger:model WeiBuData
type WeiBuData struct {

	// 恶意软件家族
	// Required: true
	MalwareFamily *string `json:"MalwareFamily"`

	// 多引擎检测结果
	// Required: true
	MultiEngines *string `json:"MultiEngines"`

	// network
	// Required: true
	Network *WeiBuNetwork `json:"Network"`

	// 文件报告页网址
	// Required: true
	Permalink *string `json:"Permalink"`

	// signature
	// Required: true
	Signature *WeiBuSignature `json:"Signature"`

	// strings
	// Required: true
	Strings *Strings `json:"Strings"`

	// 文件提交时间
	// Required: true
	SubmitTime *string `json:"SubmitTime"`

	// 标签
	// Required: true
	Tag *WeiBuTag `json:"Tag"`

	// 威胁等级
	// Required: true
	ThreatLevel *string `json:"ThreatLevel"`

	// 威胁分值
	// Required: true
	ThreatScore *int64 `json:"ThreatScore"`

	// 威胁类型
	// Required: true
	ThreatType *string `json:"ThreatType"`
}

// Validate validates this wei bu data
func (m *WeiBuData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMalwareFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMultiEngines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermalink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmitTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreatLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreatScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreatType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WeiBuData) validateMalwareFamily(formats strfmt.Registry) error {

	if err := validate.Required("MalwareFamily", "body", m.MalwareFamily); err != nil {
		return err
	}

	return nil
}

func (m *WeiBuData) validateMultiEngines(formats strfmt.Registry) error {

	if err := validate.Required("MultiEngines", "body", m.MultiEngines); err != nil {
		return err
	}

	return nil
}

func (m *WeiBuData) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Required("Network", "body", m.Network); err != nil {
		return err
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Network")
			}
			return err
		}
	}

	return nil
}

func (m *WeiBuData) validatePermalink(formats strfmt.Registry) error {

	if err := validate.Required("Permalink", "body", m.Permalink); err != nil {
		return err
	}

	return nil
}

func (m *WeiBuData) validateSignature(formats strfmt.Registry) error {

	if err := validate.Required("Signature", "body", m.Signature); err != nil {
		return err
	}

	if m.Signature != nil {
		if err := m.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Signature")
			}
			return err
		}
	}

	return nil
}

func (m *WeiBuData) validateStrings(formats strfmt.Registry) error {

	if err := validate.Required("Strings", "body", m.Strings); err != nil {
		return err
	}

	if m.Strings != nil {
		if err := m.Strings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Strings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Strings")
			}
			return err
		}
	}

	return nil
}

func (m *WeiBuData) validateSubmitTime(formats strfmt.Registry) error {

	if err := validate.Required("SubmitTime", "body", m.SubmitTime); err != nil {
		return err
	}

	return nil
}

func (m *WeiBuData) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("Tag", "body", m.Tag); err != nil {
		return err
	}

	if m.Tag != nil {
		if err := m.Tag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Tag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Tag")
			}
			return err
		}
	}

	return nil
}

func (m *WeiBuData) validateThreatLevel(formats strfmt.Registry) error {

	if err := validate.Required("ThreatLevel", "body", m.ThreatLevel); err != nil {
		return err
	}

	return nil
}

func (m *WeiBuData) validateThreatScore(formats strfmt.Registry) error {

	if err := validate.Required("ThreatScore", "body", m.ThreatScore); err != nil {
		return err
	}

	return nil
}

func (m *WeiBuData) validateThreatType(formats strfmt.Registry) error {

	if err := validate.Required("ThreatType", "body", m.ThreatType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this wei bu data based on the context it is used
func (m *WeiBuData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStrings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WeiBuData) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.Network != nil {

		if err := m.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Network")
			}
			return err
		}
	}

	return nil
}

func (m *WeiBuData) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if m.Signature != nil {

		if err := m.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Signature")
			}
			return err
		}
	}

	return nil
}

func (m *WeiBuData) contextValidateStrings(ctx context.Context, formats strfmt.Registry) error {

	if m.Strings != nil {

		if err := m.Strings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Strings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Strings")
			}
			return err
		}
	}

	return nil
}

func (m *WeiBuData) contextValidateTag(ctx context.Context, formats strfmt.Registry) error {

	if m.Tag != nil {

		if err := m.Tag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Tag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Tag")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WeiBuData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WeiBuData) UnmarshalBinary(b []byte) error {
	var res WeiBuData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
