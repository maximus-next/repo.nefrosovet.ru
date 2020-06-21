// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetWebhooksOKBody get webhooks o k body
// swagger:model getWebhooksOKBody
type GetWebhooksOKBody struct {
	SuccessData

	GetWebhooksOKBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetWebhooksOKBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SuccessData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SuccessData = aO0

	// AO1
	var aO1 GetWebhooksOKBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.GetWebhooksOKBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetWebhooksOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.GetWebhooksOKBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get webhooks o k body
func (m *GetWebhooksOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SuccessData
	if err := m.SuccessData.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with GetWebhooksOKBodyAllOf1
	if err := m.GetWebhooksOKBodyAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetWebhooksOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWebhooksOKBody) UnmarshalBinary(b []byte) error {
	var res GetWebhooksOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}