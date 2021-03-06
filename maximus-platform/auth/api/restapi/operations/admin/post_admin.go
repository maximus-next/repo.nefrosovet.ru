// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// PostAdminHandlerFunc turns a function with the right signature into a post admin handler
type PostAdminHandlerFunc func(PostAdminParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAdminHandlerFunc) Handle(params PostAdminParams) middleware.Responder {
	return fn(params)
}

// PostAdminHandler interface for that can handle valid post admin params
type PostAdminHandler interface {
	Handle(PostAdminParams) middleware.Responder
}

// NewPostAdmin creates a new http.Handler for the post admin operation
func NewPostAdmin(ctx *middleware.Context, handler PostAdminHandler) *PostAdmin {
	return &PostAdmin{Context: ctx, Handler: handler}
}

/*PostAdmin swagger:route POST /admin Admin postAdmin

Редактирование админской учетной записи

*/
type PostAdmin struct {
	Context *middleware.Context
	Handler PostAdminHandler
}

func (o *PostAdmin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAdminParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAdminBadRequestBody post admin bad request body
// swagger:model PostAdminBadRequestBody
type PostAdminBadRequestBody struct {
	models.Error400Data

	// errors
	// Required: true
	Errors *PostAdminBadRequestBodyAO1Errors `json:"errors"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostAdminBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostAdminBadRequestBodyAO0
	var postAdminBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &postAdminBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = postAdminBadRequestBodyAO0

	// PostAdminBadRequestBodyAO1
	var dataPostAdminBadRequestBodyAO1 struct {
		Errors *PostAdminBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}
	if err := swag.ReadJSON(raw, &dataPostAdminBadRequestBodyAO1); err != nil {
		return err
	}

	o.Errors = dataPostAdminBadRequestBodyAO1.Errors

	o.Message = dataPostAdminBadRequestBodyAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostAdminBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postAdminBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postAdminBadRequestBodyAO0)

	var dataPostAdminBadRequestBodyAO1 struct {
		Errors *PostAdminBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}

	dataPostAdminBadRequestBodyAO1.Errors = o.Errors

	dataPostAdminBadRequestBodyAO1.Message = o.Message

	jsonDataPostAdminBadRequestBodyAO1, errPostAdminBadRequestBodyAO1 := swag.WriteJSON(dataPostAdminBadRequestBodyAO1)
	if errPostAdminBadRequestBodyAO1 != nil {
		return nil, errPostAdminBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataPostAdminBadRequestBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post admin bad request body
func (o *PostAdminBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *PostAdminBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("postAdminBadRequest"+"."+"errors", "body", o.Errors); err != nil {
		return err
	}

	if o.Errors != nil {
		if err := o.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postAdminBadRequest" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

func (o *PostAdminBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("postAdminBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostAdminBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdminBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostAdminBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAdminBadRequestBodyAO1Errors post admin bad request body a o1 errors
// swagger:model PostAdminBadRequestBodyAO1Errors
type PostAdminBadRequestBodyAO1Errors struct {

	// core
	Core string `json:"core,omitempty"`

	// json
	JSON string `json:"json,omitempty"`

	// validation
	Validation interface{} `json:"validation,omitempty"`
}

// Validate validates this post admin bad request body a o1 errors
func (o *PostAdminBadRequestBodyAO1Errors) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAdminBadRequestBodyAO1Errors) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdminBadRequestBodyAO1Errors) UnmarshalBinary(b []byte) error {
	var res PostAdminBadRequestBodyAO1Errors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAdminBody post admin body
// swagger:model PostAdminBody
type PostAdminBody struct {
	models.AdminParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostAdminBody) UnmarshalJSON(raw []byte) error {
	// PostAdminParamsBodyAO0
	var postAdminParamsBodyAO0 models.AdminParams
	if err := swag.ReadJSON(raw, &postAdminParamsBodyAO0); err != nil {
		return err
	}
	o.AdminParams = postAdminParamsBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostAdminBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postAdminParamsBodyAO0, err := swag.WriteJSON(o.AdminParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postAdminParamsBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post admin body
func (o *PostAdminBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.AdminParams
	if err := o.AdminParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostAdminBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdminBody) UnmarshalBinary(b []byte) error {
	var res PostAdminBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAdminInternalServerErrorBody post admin internal server error body
// swagger:model PostAdminInternalServerErrorBody
type PostAdminInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostAdminInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PostAdminInternalServerErrorBodyAO0
	var postAdminInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &postAdminInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = postAdminInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostAdminInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postAdminInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postAdminInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post admin internal server error body
func (o *PostAdminInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *PostAdminInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdminInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostAdminInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAdminMethodNotAllowedBody post admin method not allowed body
// swagger:model PostAdminMethodNotAllowedBody
type PostAdminMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostAdminMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostAdminMethodNotAllowedBodyAO0
	var postAdminMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postAdminMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postAdminMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostAdminMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postAdminMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postAdminMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post admin method not allowed body
func (o *PostAdminMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostAdminMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdminMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostAdminMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAdminOKBody post admin o k body
// swagger:model PostAdminOKBody
type PostAdminOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []interface{} `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostAdminOKBody) UnmarshalJSON(raw []byte) error {
	// PostAdminOKBodyAO0
	var postAdminOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postAdminOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postAdminOKBodyAO0

	// PostAdminOKBodyAO1
	var dataPostAdminOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPostAdminOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostAdminOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostAdminOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postAdminOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postAdminOKBodyAO0)

	var dataPostAdminOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}

	dataPostAdminOKBodyAO1.Data = o.Data

	jsonDataPostAdminOKBodyAO1, errPostAdminOKBodyAO1 := swag.WriteJSON(dataPostAdminOKBodyAO1)
	if errPostAdminOKBodyAO1 != nil {
		return nil, errPostAdminOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostAdminOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post admin o k body
func (o *PostAdminOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostAdminOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("postAdminOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostAdminOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdminOKBody) UnmarshalBinary(b []byte) error {
	var res PostAdminOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
