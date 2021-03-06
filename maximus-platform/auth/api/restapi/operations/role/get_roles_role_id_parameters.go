// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRolesRoleIDParams creates a new GetRolesRoleIDParams object
// no default values defined in spec.
func NewGetRolesRoleIDParams() GetRolesRoleIDParams {

	return GetRolesRoleIDParams{}
}

// GetRolesRoleIDParams contains all the bound params for the get roles role ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetRolesRoleID
type GetRolesRoleIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор роли
	  Required: true
	  In: path
	*/
	RoleID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRolesRoleIDParams() beforehand.
func (o *GetRolesRoleIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rRoleID, rhkRoleID, _ := route.Params.GetOK("roleID")
	if err := o.bindRoleID(rRoleID, rhkRoleID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindRoleID binds and validates parameter RoleID from path.
func (o *GetRolesRoleIDParams) bindRoleID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.RoleID = raw

	return nil
}
