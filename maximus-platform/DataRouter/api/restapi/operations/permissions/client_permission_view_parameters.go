// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewClientPermissionViewParams creates a new ClientPermissionViewParams object
// no default values defined in spec.
func NewClientPermissionViewParams() ClientPermissionViewParams {

	return ClientPermissionViewParams{}
}

// ClientPermissionViewParams contains all the bound params for the client permission view operation
// typically these are obtained from a http.Request
//
// swagger:parameters clientPermissionView
type ClientPermissionViewParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор клиента
	  Required: true
	  In: path
	*/
	ClientID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewClientPermissionViewParams() beforehand.
func (o *ClientPermissionViewParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rClientID, rhkClientID, _ := route.Params.GetOK("clientID")
	if err := o.bindClientID(rClientID, rhkClientID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClientID binds and validates parameter ClientID from path.
func (o *ClientPermissionViewParams) bindClientID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("clientID", "path", "strfmt.UUID", raw)
	}
	o.ClientID = *(value.(*strfmt.UUID))

	if err := o.validateClientID(formats); err != nil {
		return err
	}

	return nil
}

// validateClientID carries on validations for parameter ClientID
func (o *ClientPermissionViewParams) validateClientID(formats strfmt.Registry) error {

	if err := validate.FormatOf("clientID", "path", "uuid", o.ClientID.String(), formats); err != nil {
		return err
	}
	return nil
}