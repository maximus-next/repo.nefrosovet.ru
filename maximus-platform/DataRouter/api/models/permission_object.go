// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PermissionObject Permission_object
// swagger:model Permission_object
type PermissionObject struct {

	// publish
	Publish []string `json:"publish"`

	// subscribe
	Subscribe []string `json:"subscribe"`
}

// Validate validates this permission object
func (m *PermissionObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionObject) UnmarshalBinary(b []byte) error {
	var res PermissionObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
