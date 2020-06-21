// Code generated by go-swagger; DO NOT EDIT.

package employee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewEmployeeViewParams creates a new EmployeeViewParams object
// no default values defined in spec.
func NewEmployeeViewParams() EmployeeViewParams {

	return EmployeeViewParams{}
}

// EmployeeViewParams contains all the bound params for the employee view operation
// typically these are obtained from a http.Request
//
// swagger:parameters EmployeeView
type EmployeeViewParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор сотрудника клиники
	  Required: true
	  In: path
	*/
	EmployeeID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewEmployeeViewParams() beforehand.
func (o *EmployeeViewParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rEmployeeID, rhkEmployeeID, _ := route.Params.GetOK("employeeID")
	if err := o.bindEmployeeID(rEmployeeID, rhkEmployeeID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEmployeeID binds and validates parameter EmployeeID from path.
func (o *EmployeeViewParams) bindEmployeeID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("employeeID", "path", "strfmt.UUID", raw)
	}
	o.EmployeeID = *(value.(*strfmt.UUID))

	if err := o.validateEmployeeID(formats); err != nil {
		return err
	}

	return nil
}

// validateEmployeeID carries on validations for parameter EmployeeID
func (o *EmployeeViewParams) validateEmployeeID(formats strfmt.Registry) error {

	if err := validate.FormatOf("employeeID", "path", "uuid", o.EmployeeID.String(), formats); err != nil {
		return err
	}
	return nil
}