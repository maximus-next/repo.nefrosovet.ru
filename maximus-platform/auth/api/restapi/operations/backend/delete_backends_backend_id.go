// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "repo.nefrosovet.ru/maximus-platform/auth/api/models"
)

// DeleteBackendsBackendIDHandlerFunc turns a function with the right signature into a delete backends backend ID handler
type DeleteBackendsBackendIDHandlerFunc func(DeleteBackendsBackendIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBackendsBackendIDHandlerFunc) Handle(params DeleteBackendsBackendIDParams) middleware.Responder {
	return fn(params)
}

// DeleteBackendsBackendIDHandler interface for that can handle valid delete backends backend ID params
type DeleteBackendsBackendIDHandler interface {
	Handle(DeleteBackendsBackendIDParams) middleware.Responder
}

// NewDeleteBackendsBackendID creates a new http.Handler for the delete backends backend ID operation
func NewDeleteBackendsBackendID(ctx *middleware.Context, handler DeleteBackendsBackendIDHandler) *DeleteBackendsBackendID {
	return &DeleteBackendsBackendID{Context: ctx, Handler: handler}
}

/*DeleteBackendsBackendID swagger:route DELETE /backends/{backendID} Backend deleteBackendsBackendId

Удаление бэкенда

*/
type DeleteBackendsBackendID struct {
	Context *middleware.Context
	Handler DeleteBackendsBackendIDHandler
}

func (o *DeleteBackendsBackendID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteBackendsBackendIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteBackendsBackendIDInternalServerErrorBody delete backends backend ID internal server error body
// swagger:model DeleteBackendsBackendIDInternalServerErrorBody
type DeleteBackendsBackendIDInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteBackendsBackendIDInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// DeleteBackendsBackendIDInternalServerErrorBodyAO0
	var deleteBackendsBackendIDInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &deleteBackendsBackendIDInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = deleteBackendsBackendIDInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteBackendsBackendIDInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteBackendsBackendIDInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteBackendsBackendIDInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete backends backend ID internal server error body
func (o *DeleteBackendsBackendIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error500Data
	if err := o.Error500Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteBackendsBackendIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteBackendsBackendIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeleteBackendsBackendIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteBackendsBackendIDMethodNotAllowedBody delete backends backend ID method not allowed body
// swagger:model DeleteBackendsBackendIDMethodNotAllowedBody
type DeleteBackendsBackendIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteBackendsBackendIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// DeleteBackendsBackendIDMethodNotAllowedBodyAO0
	var deleteBackendsBackendIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &deleteBackendsBackendIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = deleteBackendsBackendIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteBackendsBackendIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteBackendsBackendIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteBackendsBackendIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete backends backend ID method not allowed body
func (o *DeleteBackendsBackendIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error405Data
	if err := o.Error405Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteBackendsBackendIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteBackendsBackendIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res DeleteBackendsBackendIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteBackendsBackendIDNotFoundBody delete backends backend ID not found body
// swagger:model DeleteBackendsBackendIDNotFoundBody
type DeleteBackendsBackendIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteBackendsBackendIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// DeleteBackendsBackendIDNotFoundBodyAO0
	var deleteBackendsBackendIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &deleteBackendsBackendIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = deleteBackendsBackendIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteBackendsBackendIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteBackendsBackendIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteBackendsBackendIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete backends backend ID not found body
func (o *DeleteBackendsBackendIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error404Data
	if err := o.Error404Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteBackendsBackendIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteBackendsBackendIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteBackendsBackendIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteBackendsBackendIDOKBody delete backends backend ID o k body
// swagger:model DeleteBackendsBackendIDOKBody
type DeleteBackendsBackendIDOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []interface{} `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteBackendsBackendIDOKBody) UnmarshalJSON(raw []byte) error {
	// DeleteBackendsBackendIDOKBodyAO0
	var deleteBackendsBackendIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &deleteBackendsBackendIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = deleteBackendsBackendIDOKBodyAO0

	// DeleteBackendsBackendIDOKBodyAO1
	var dataDeleteBackendsBackendIDOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataDeleteBackendsBackendIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataDeleteBackendsBackendIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteBackendsBackendIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	deleteBackendsBackendIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteBackendsBackendIDOKBodyAO0)

	var dataDeleteBackendsBackendIDOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}

	dataDeleteBackendsBackendIDOKBodyAO1.Data = o.Data

	jsonDataDeleteBackendsBackendIDOKBodyAO1, errDeleteBackendsBackendIDOKBodyAO1 := swag.WriteJSON(dataDeleteBackendsBackendIDOKBodyAO1)
	if errDeleteBackendsBackendIDOKBodyAO1 != nil {
		return nil, errDeleteBackendsBackendIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataDeleteBackendsBackendIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete backends backend ID o k body
func (o *DeleteBackendsBackendIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SuccessData
	if err := o.SuccessData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteBackendsBackendIDOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("deleteBackendsBackendIdOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteBackendsBackendIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteBackendsBackendIDOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteBackendsBackendIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}