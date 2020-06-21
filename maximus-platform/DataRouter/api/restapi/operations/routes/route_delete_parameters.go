// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRouteDeleteParams creates a new RouteDeleteParams object
// no default values defined in spec.
func NewRouteDeleteParams() RouteDeleteParams {

	return RouteDeleteParams{}
}

// RouteDeleteParams contains all the bound params for the route delete operation
// typically these are obtained from a http.Request
//
// swagger:parameters routeDelete
type RouteDeleteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор маршрута
	  Required: true
	  In: path
	*/
	RouteID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRouteDeleteParams() beforehand.
func (o *RouteDeleteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rRouteID, rhkRouteID, _ := route.Params.GetOK("routeID")
	if err := o.bindRouteID(rRouteID, rhkRouteID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindRouteID binds and validates parameter RouteID from path.
func (o *RouteDeleteParams) bindRouteID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("routeID", "path", "strfmt.UUID", raw)
	}
	o.RouteID = *(value.(*strfmt.UUID))

	if err := o.validateRouteID(formats); err != nil {
		return err
	}

	return nil
}

// validateRouteID carries on validations for parameter RouteID
func (o *RouteDeleteParams) validateRouteID(formats strfmt.Registry) error {

	if err := validate.FormatOf("routeID", "path", "uuid", o.RouteID.String(), formats); err != nil {
		return err
	}
	return nil
}