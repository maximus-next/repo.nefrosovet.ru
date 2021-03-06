// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostBackendsBackendIDGroupsParams creates a new PostBackendsBackendIDGroupsParams object
// no default values defined in spec.
func NewPostBackendsBackendIDGroupsParams() PostBackendsBackendIDGroupsParams {

	return PostBackendsBackendIDGroupsParams{}
}

// PostBackendsBackendIDGroupsParams contains all the bound params for the post backends backend ID groups operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostBackendsBackendIDGroups
type PostBackendsBackendIDGroupsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор бэкенда
	  Required: true
	  In: path
	*/
	BackendID string
	/*
	  Required: true
	  In: body
	*/
	Body []*PostBackendsBackendIDGroupsParamsBodyItems0
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostBackendsBackendIDGroupsParams() beforehand.
func (o *PostBackendsBackendIDGroupsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBackendID, rhkBackendID, _ := route.Params.GetOK("backendID")
	if err := o.bindBackendID(rBackendID, rhkBackendID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*PostBackendsBackendIDGroupsParamsBodyItems0
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate array of body objects
			for i := range body {
				if body[i] == nil {
					continue
				}
				if err := body[i].Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}
			if len(res) == 0 {
				o.Body = body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBackendID binds and validates parameter BackendID from path.
func (o *PostBackendsBackendIDGroupsParams) bindBackendID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.BackendID = raw

	return nil
}
