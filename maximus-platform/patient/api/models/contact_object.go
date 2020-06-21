// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContactObject Contact_object
// swagger:model Contact_object
type ContactObject struct {

	// тип контактных данных
	// Enum: [EMAIL MOBILE]
	Type string `json:"type,omitempty"`

	// Контактные данные
	Value string `json:"value,omitempty"`

	// Статус подтверждения
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this contact object
func (m *ContactObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var contactObjectTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EMAIL","MOBILE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactObjectTypeTypePropEnum = append(contactObjectTypeTypePropEnum, v)
	}
}

const (

	// ContactObjectTypeEMAIL captures enum value "EMAIL"
	ContactObjectTypeEMAIL string = "EMAIL"

	// ContactObjectTypeMOBILE captures enum value "MOBILE"
	ContactObjectTypeMOBILE string = "MOBILE"
)

// prop value enum
func (m *ContactObject) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, contactObjectTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ContactObject) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactObject) UnmarshalBinary(b []byte) error {
	var res ContactObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}