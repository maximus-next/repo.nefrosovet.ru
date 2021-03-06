// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostWebhooksViberParamsBody post webhooks viber params body
// swagger:model postWebhooksViberParamsBody
type PostWebhooksViberParamsBody struct {
	ViberParams

	PostWebhooksViberParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostWebhooksViberParamsBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ViberParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ViberParams = aO0

	// AO1
	var aO1 PostWebhooksViberParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PostWebhooksViberParamsBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostWebhooksViberParamsBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ViberParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PostWebhooksViberParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post webhooks viber params body
func (m *PostWebhooksViberParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ViberParams
	if err := m.ViberParams.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PostWebhooksViberParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostWebhooksViberParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostWebhooksViberParamsBody) UnmarshalBinary(b []byte) error {
	var res PostWebhooksViberParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostWebhooksViberParamsBodyAllOf1 post webhooks viber params body all of1
// swagger:model PostWebhooksViberParamsBodyAllOf1
type PostWebhooksViberParamsBodyAllOf1 interface{}
