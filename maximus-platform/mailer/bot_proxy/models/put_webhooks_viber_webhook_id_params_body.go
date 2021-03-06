// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PutWebhooksViberWebhookIDParamsBody put webhooks viber webhook Id params body
// swagger:model putWebhooksViberWebhookIdParamsBody
type PutWebhooksViberWebhookIDParamsBody struct {
	ViberParams

	PutWebhooksViberWebhookIDParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PutWebhooksViberWebhookIDParamsBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ViberParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ViberParams = aO0

	// AO1
	var aO1 PutWebhooksViberWebhookIDParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PutWebhooksViberWebhookIDParamsBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PutWebhooksViberWebhookIDParamsBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ViberParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PutWebhooksViberWebhookIDParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put webhooks viber webhook Id params body
func (m *PutWebhooksViberWebhookIDParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ViberParams
	if err := m.ViberParams.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PutWebhooksViberWebhookIDParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PutWebhooksViberWebhookIDParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutWebhooksViberWebhookIDParamsBody) UnmarshalBinary(b []byte) error {
	var res PutWebhooksViberWebhookIDParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutWebhooksViberWebhookIDParamsBodyAllOf1 put webhooks viber webhook ID params body all of1
// swagger:model PutWebhooksViberWebhookIDParamsBodyAllOf1
type PutWebhooksViberWebhookIDParamsBodyAllOf1 interface{}
