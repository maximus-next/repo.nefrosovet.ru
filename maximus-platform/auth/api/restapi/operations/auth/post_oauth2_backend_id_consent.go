// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "repo.nefrosovet.ru/maximus-platform/auth/api/models"
)

// PostOauth2BackendIDConsentHandlerFunc turns a function with the right signature into a post oauth2 backend ID consent handler
type PostOauth2BackendIDConsentHandlerFunc func(PostOauth2BackendIDConsentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOauth2BackendIDConsentHandlerFunc) Handle(params PostOauth2BackendIDConsentParams) middleware.Responder {
	return fn(params)
}

// PostOauth2BackendIDConsentHandler interface for that can handle valid post oauth2 backend ID consent params
type PostOauth2BackendIDConsentHandler interface {
	Handle(PostOauth2BackendIDConsentParams) middleware.Responder
}

// NewPostOauth2BackendIDConsent creates a new http.Handler for the post oauth2 backend ID consent operation
func NewPostOauth2BackendIDConsent(ctx *middleware.Context, handler PostOauth2BackendIDConsentHandler) *PostOauth2BackendIDConsent {
	return &PostOauth2BackendIDConsent{Context: ctx, Handler: handler}
}

/*PostOauth2BackendIDConsent swagger:route POST /oauth2/{backendID}/consent Auth postOauth2BackendIdConsent

Обмен authorizationCode на access_token и refresh_token

*/
type PostOauth2BackendIDConsent struct {
	Context *middleware.Context
	Handler PostOauth2BackendIDConsentHandler
}

func (o *PostOauth2BackendIDConsent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOauth2BackendIDConsentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostOauth2BackendIDConsentBadRequestBody post oauth2 backend ID consent bad request body
// swagger:model PostOauth2BackendIDConsentBadRequestBody
type PostOauth2BackendIDConsentBadRequestBody struct {
	models.Error400Data

	// errors
	// Required: true
	Errors *PostOauth2BackendIDConsentBadRequestBodyAO1Errors `json:"errors"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostOauth2BackendIDConsentBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostOauth2BackendIDConsentBadRequestBodyAO0
	var postOauth2BackendIDConsentBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &postOauth2BackendIDConsentBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = postOauth2BackendIDConsentBadRequestBodyAO0

	// PostOauth2BackendIDConsentBadRequestBodyAO1
	var dataPostOauth2BackendIDConsentBadRequestBodyAO1 struct {
		Errors *PostOauth2BackendIDConsentBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}
	if err := swag.ReadJSON(raw, &dataPostOauth2BackendIDConsentBadRequestBodyAO1); err != nil {
		return err
	}

	o.Errors = dataPostOauth2BackendIDConsentBadRequestBodyAO1.Errors

	o.Message = dataPostOauth2BackendIDConsentBadRequestBodyAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostOauth2BackendIDConsentBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postOauth2BackendIDConsentBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postOauth2BackendIDConsentBadRequestBodyAO0)

	var dataPostOauth2BackendIDConsentBadRequestBodyAO1 struct {
		Errors *PostOauth2BackendIDConsentBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}

	dataPostOauth2BackendIDConsentBadRequestBodyAO1.Errors = o.Errors

	dataPostOauth2BackendIDConsentBadRequestBodyAO1.Message = o.Message

	jsonDataPostOauth2BackendIDConsentBadRequestBodyAO1, errPostOauth2BackendIDConsentBadRequestBodyAO1 := swag.WriteJSON(dataPostOauth2BackendIDConsentBadRequestBodyAO1)
	if errPostOauth2BackendIDConsentBadRequestBodyAO1 != nil {
		return nil, errPostOauth2BackendIDConsentBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataPostOauth2BackendIDConsentBadRequestBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post oauth2 backend ID consent bad request body
func (o *PostOauth2BackendIDConsentBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *PostOauth2BackendIDConsentBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("postOauth2BackendIdConsentBadRequest"+"."+"errors", "body", o.Errors); err != nil {
		return err
	}

	if o.Errors != nil {
		if err := o.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOauth2BackendIdConsentBadRequest" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

func (o *PostOauth2BackendIDConsentBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("postOauth2BackendIdConsentBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOauth2BackendIDConsentBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOauth2BackendIDConsentBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostOauth2BackendIDConsentBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOauth2BackendIDConsentBadRequestBodyAO1Errors post oauth2 backend ID consent bad request body a o1 errors
// swagger:model PostOauth2BackendIDConsentBadRequestBodyAO1Errors
type PostOauth2BackendIDConsentBadRequestBodyAO1Errors struct {

	// core
	Core string `json:"core,omitempty"`

	// json
	JSON string `json:"json,omitempty"`

	// validation
	Validation interface{} `json:"validation,omitempty"`
}

// Validate validates this post oauth2 backend ID consent bad request body a o1 errors
func (o *PostOauth2BackendIDConsentBadRequestBodyAO1Errors) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOauth2BackendIDConsentBadRequestBodyAO1Errors) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOauth2BackendIDConsentBadRequestBodyAO1Errors) UnmarshalBinary(b []byte) error {
	var res PostOauth2BackendIDConsentBadRequestBodyAO1Errors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOauth2BackendIDConsentBody post oauth2 backend ID consent body
// swagger:model PostOauth2BackendIDConsentBody
type PostOauth2BackendIDConsentBody struct {
	models.OAuth2ConsentParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostOauth2BackendIDConsentBody) UnmarshalJSON(raw []byte) error {
	// PostOauth2BackendIDConsentParamsBodyAO0
	var postOauth2BackendIDConsentParamsBodyAO0 models.OAuth2ConsentParams
	if err := swag.ReadJSON(raw, &postOauth2BackendIDConsentParamsBodyAO0); err != nil {
		return err
	}
	o.OAuth2ConsentParams = postOauth2BackendIDConsentParamsBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostOauth2BackendIDConsentBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postOauth2BackendIDConsentParamsBodyAO0, err := swag.WriteJSON(o.OAuth2ConsentParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postOauth2BackendIDConsentParamsBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post oauth2 backend ID consent body
func (o *PostOauth2BackendIDConsentBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.OAuth2ConsentParams
	if err := o.OAuth2ConsentParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostOauth2BackendIDConsentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOauth2BackendIDConsentBody) UnmarshalBinary(b []byte) error {
	var res PostOauth2BackendIDConsentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOauth2BackendIDConsentInternalServerErrorBody post oauth2 backend ID consent internal server error body
// swagger:model PostOauth2BackendIDConsentInternalServerErrorBody
type PostOauth2BackendIDConsentInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostOauth2BackendIDConsentInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PostOauth2BackendIDConsentInternalServerErrorBodyAO0
	var postOauth2BackendIDConsentInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &postOauth2BackendIDConsentInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = postOauth2BackendIDConsentInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostOauth2BackendIDConsentInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postOauth2BackendIDConsentInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postOauth2BackendIDConsentInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post oauth2 backend ID consent internal server error body
func (o *PostOauth2BackendIDConsentInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *PostOauth2BackendIDConsentInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOauth2BackendIDConsentInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostOauth2BackendIDConsentInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOauth2BackendIDConsentOKBody post oauth2 backend ID consent o k body
// swagger:model PostOauth2BackendIDConsentOKBody
type PostOauth2BackendIDConsentOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostOauth2BackendIDConsentOKBody) UnmarshalJSON(raw []byte) error {
	// PostOauth2BackendIDConsentOKBodyAO0
	var postOauth2BackendIDConsentOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postOauth2BackendIDConsentOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postOauth2BackendIDConsentOKBodyAO0

	// PostOauth2BackendIDConsentOKBodyAO1
	var dataPostOauth2BackendIDConsentOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPostOauth2BackendIDConsentOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostOauth2BackendIDConsentOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostOauth2BackendIDConsentOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postOauth2BackendIDConsentOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postOauth2BackendIDConsentOKBodyAO0)

	var dataPostOauth2BackendIDConsentOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataPostOauth2BackendIDConsentOKBodyAO1.Data = o.Data

	jsonDataPostOauth2BackendIDConsentOKBodyAO1, errPostOauth2BackendIDConsentOKBodyAO1 := swag.WriteJSON(dataPostOauth2BackendIDConsentOKBodyAO1)
	if errPostOauth2BackendIDConsentOKBodyAO1 != nil {
		return nil, errPostOauth2BackendIDConsentOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostOauth2BackendIDConsentOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post oauth2 backend ID consent o k body
func (o *PostOauth2BackendIDConsentOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostOauth2BackendIDConsentOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOauth2BackendIdConsentOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOauth2BackendIDConsentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOauth2BackendIDConsentOKBody) UnmarshalBinary(b []byte) error {
	var res PostOauth2BackendIDConsentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}