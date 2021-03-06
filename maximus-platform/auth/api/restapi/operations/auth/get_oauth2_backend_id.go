// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// GetOauth2BackendIDHandlerFunc turns a function with the right signature into a get oauth2 backend ID handler
type GetOauth2BackendIDHandlerFunc func(GetOauth2BackendIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOauth2BackendIDHandlerFunc) Handle(params GetOauth2BackendIDParams) middleware.Responder {
	return fn(params)
}

// GetOauth2BackendIDHandler interface for that can handle valid get oauth2 backend ID params
type GetOauth2BackendIDHandler interface {
	Handle(GetOauth2BackendIDParams) middleware.Responder
}

// NewGetOauth2BackendID creates a new http.Handler for the get oauth2 backend ID operation
func NewGetOauth2BackendID(ctx *middleware.Context, handler GetOauth2BackendIDHandler) *GetOauth2BackendID {
	return &GetOauth2BackendID{Context: ctx, Handler: handler}
}

/*GetOauth2BackendID swagger:route GET /oauth2/{backendID} Auth getOauth2BackendId

Получение URI для аутентификации по OAuth2

*/
type GetOauth2BackendID struct {
	Context *middleware.Context
	Handler GetOauth2BackendIDHandler
}

func (o *GetOauth2BackendID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetOauth2BackendIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetOauth2BackendIDBadRequestBody get oauth2 backend ID bad request body
// swagger:model GetOauth2BackendIDBadRequestBody
type GetOauth2BackendIDBadRequestBody struct {
	models.Error400Data

	// errors
	// Required: true
	Errors *GetOauth2BackendIDBadRequestBodyAO1Errors `json:"errors"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetOauth2BackendIDBadRequestBody) UnmarshalJSON(raw []byte) error {
	// GetOauth2BackendIDBadRequestBodyAO0
	var getOauth2BackendIDBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &getOauth2BackendIDBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = getOauth2BackendIDBadRequestBodyAO0

	// GetOauth2BackendIDBadRequestBodyAO1
	var dataGetOauth2BackendIDBadRequestBodyAO1 struct {
		Errors *GetOauth2BackendIDBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}
	if err := swag.ReadJSON(raw, &dataGetOauth2BackendIDBadRequestBodyAO1); err != nil {
		return err
	}

	o.Errors = dataGetOauth2BackendIDBadRequestBodyAO1.Errors

	o.Message = dataGetOauth2BackendIDBadRequestBodyAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetOauth2BackendIDBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getOauth2BackendIDBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getOauth2BackendIDBadRequestBodyAO0)

	var dataGetOauth2BackendIDBadRequestBodyAO1 struct {
		Errors *GetOauth2BackendIDBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}

	dataGetOauth2BackendIDBadRequestBodyAO1.Errors = o.Errors

	dataGetOauth2BackendIDBadRequestBodyAO1.Message = o.Message

	jsonDataGetOauth2BackendIDBadRequestBodyAO1, errGetOauth2BackendIDBadRequestBodyAO1 := swag.WriteJSON(dataGetOauth2BackendIDBadRequestBodyAO1)
	if errGetOauth2BackendIDBadRequestBodyAO1 != nil {
		return nil, errGetOauth2BackendIDBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataGetOauth2BackendIDBadRequestBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get oauth2 backend ID bad request body
func (o *GetOauth2BackendIDBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error400Data
	if err := o.Error400Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOauth2BackendIDBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("getOauth2BackendIdBadRequest"+"."+"errors", "body", o.Errors); err != nil {
		return err
	}

	if o.Errors != nil {
		if err := o.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOauth2BackendIdBadRequest" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

func (o *GetOauth2BackendIDBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getOauth2BackendIdBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOauth2BackendIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOauth2BackendIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetOauth2BackendIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetOauth2BackendIDBadRequestBodyAO1Errors get oauth2 backend ID bad request body a o1 errors
// swagger:model GetOauth2BackendIDBadRequestBodyAO1Errors
type GetOauth2BackendIDBadRequestBodyAO1Errors struct {

	// core
	Core string `json:"core,omitempty"`

	// json
	JSON string `json:"json,omitempty"`

	// validation
	Validation interface{} `json:"validation,omitempty"`
}

// Validate validates this get oauth2 backend ID bad request body a o1 errors
func (o *GetOauth2BackendIDBadRequestBodyAO1Errors) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOauth2BackendIDBadRequestBodyAO1Errors) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOauth2BackendIDBadRequestBodyAO1Errors) UnmarshalBinary(b []byte) error {
	var res GetOauth2BackendIDBadRequestBodyAO1Errors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetOauth2BackendIDInternalServerErrorBody get oauth2 backend ID internal server error body
// swagger:model GetOauth2BackendIDInternalServerErrorBody
type GetOauth2BackendIDInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetOauth2BackendIDInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// GetOauth2BackendIDInternalServerErrorBodyAO0
	var getOauth2BackendIDInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &getOauth2BackendIDInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = getOauth2BackendIDInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetOauth2BackendIDInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getOauth2BackendIDInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getOauth2BackendIDInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get oauth2 backend ID internal server error body
func (o *GetOauth2BackendIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *GetOauth2BackendIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOauth2BackendIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetOauth2BackendIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetOauth2BackendIDOKBody get oauth2 backend ID o k body
// swagger:model GetOauth2BackendIDOKBody
type GetOauth2BackendIDOKBody struct {

	// Путь для аутентификации клиента
	// Required: true
	OAuth2Path *string `json:"OAuth2Path"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetOauth2BackendIDOKBody) UnmarshalJSON(raw []byte) error {
	// GetOauth2BackendIDOKBodyAO0
	var dataGetOauth2BackendIDOKBodyAO0 struct {
		OAuth2Path *string `json:"OAuth2Path"`
	}
	if err := swag.ReadJSON(raw, &dataGetOauth2BackendIDOKBodyAO0); err != nil {
		return err
	}

	o.OAuth2Path = dataGetOauth2BackendIDOKBodyAO0.OAuth2Path

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetOauth2BackendIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataGetOauth2BackendIDOKBodyAO0 struct {
		OAuth2Path *string `json:"OAuth2Path"`
	}

	dataGetOauth2BackendIDOKBodyAO0.OAuth2Path = o.OAuth2Path

	jsonDataGetOauth2BackendIDOKBodyAO0, errGetOauth2BackendIDOKBodyAO0 := swag.WriteJSON(dataGetOauth2BackendIDOKBodyAO0)
	if errGetOauth2BackendIDOKBodyAO0 != nil {
		return nil, errGetOauth2BackendIDOKBodyAO0
	}
	_parts = append(_parts, jsonDataGetOauth2BackendIDOKBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get oauth2 backend ID o k body
func (o *GetOauth2BackendIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOAuth2Path(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOauth2BackendIDOKBody) validateOAuth2Path(formats strfmt.Registry) error {

	if err := validate.Required("getOauth2BackendIdOK"+"."+"OAuth2Path", "body", o.OAuth2Path); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOauth2BackendIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOauth2BackendIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetOauth2BackendIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
