// Code generated by go-swagger; DO NOT EDIT.

package role

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

// PostRolesRoleIDUsersUserIDHandlerFunc turns a function with the right signature into a post roles role ID users user ID handler
type PostRolesRoleIDUsersUserIDHandlerFunc func(PostRolesRoleIDUsersUserIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRolesRoleIDUsersUserIDHandlerFunc) Handle(params PostRolesRoleIDUsersUserIDParams) middleware.Responder {
	return fn(params)
}

// PostRolesRoleIDUsersUserIDHandler interface for that can handle valid post roles role ID users user ID params
type PostRolesRoleIDUsersUserIDHandler interface {
	Handle(PostRolesRoleIDUsersUserIDParams) middleware.Responder
}

// NewPostRolesRoleIDUsersUserID creates a new http.Handler for the post roles role ID users user ID operation
func NewPostRolesRoleIDUsersUserID(ctx *middleware.Context, handler PostRolesRoleIDUsersUserIDHandler) *PostRolesRoleIDUsersUserID {
	return &PostRolesRoleIDUsersUserID{Context: ctx, Handler: handler}
}

/*PostRolesRoleIDUsersUserID swagger:route POST /roles/{roleID}/users/{userID} Role postRolesRoleIdUsersUserId

Добавление пользователя в роль

*/
type PostRolesRoleIDUsersUserID struct {
	Context *middleware.Context
	Handler PostRolesRoleIDUsersUserIDHandler
}

func (o *PostRolesRoleIDUsersUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostRolesRoleIDUsersUserIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostRolesRoleIDUsersUserIDBadRequestBody post roles role ID users user ID bad request body
// swagger:model PostRolesRoleIDUsersUserIDBadRequestBody
type PostRolesRoleIDUsersUserIDBadRequestBody struct {
	models.Error400Data

	// errors
	// Required: true
	Errors *PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors `json:"errors"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRolesRoleIDUsersUserIDBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostRolesRoleIDUsersUserIDBadRequestBodyAO0
	var postRolesRoleIDUsersUserIDBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &postRolesRoleIDUsersUserIDBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = postRolesRoleIDUsersUserIDBadRequestBodyAO0

	// PostRolesRoleIDUsersUserIDBadRequestBodyAO1
	var dataPostRolesRoleIDUsersUserIDBadRequestBodyAO1 struct {
		Errors *PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}
	if err := swag.ReadJSON(raw, &dataPostRolesRoleIDUsersUserIDBadRequestBodyAO1); err != nil {
		return err
	}

	o.Errors = dataPostRolesRoleIDUsersUserIDBadRequestBodyAO1.Errors

	o.Message = dataPostRolesRoleIDUsersUserIDBadRequestBodyAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRolesRoleIDUsersUserIDBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postRolesRoleIDUsersUserIDBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRolesRoleIDUsersUserIDBadRequestBodyAO0)

	var dataPostRolesRoleIDUsersUserIDBadRequestBodyAO1 struct {
		Errors *PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}

	dataPostRolesRoleIDUsersUserIDBadRequestBodyAO1.Errors = o.Errors

	dataPostRolesRoleIDUsersUserIDBadRequestBodyAO1.Message = o.Message

	jsonDataPostRolesRoleIDUsersUserIDBadRequestBodyAO1, errPostRolesRoleIDUsersUserIDBadRequestBodyAO1 := swag.WriteJSON(dataPostRolesRoleIDUsersUserIDBadRequestBodyAO1)
	if errPostRolesRoleIDUsersUserIDBadRequestBodyAO1 != nil {
		return nil, errPostRolesRoleIDUsersUserIDBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataPostRolesRoleIDUsersUserIDBadRequestBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post roles role ID users user ID bad request body
func (o *PostRolesRoleIDUsersUserIDBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *PostRolesRoleIDUsersUserIDBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("postRolesRoleIdUsersUserIdBadRequest"+"."+"errors", "body", o.Errors); err != nil {
		return err
	}

	if o.Errors != nil {
		if err := o.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postRolesRoleIdUsersUserIdBadRequest" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

func (o *PostRolesRoleIDUsersUserIDBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("postRolesRoleIdUsersUserIdBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostRolesRoleIDUsersUserIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRolesRoleIDUsersUserIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostRolesRoleIDUsersUserIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors post roles role ID users user ID bad request body a o1 errors
// swagger:model PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors
type PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors struct {

	// core
	Core string `json:"core,omitempty"`

	// json
	JSON string `json:"json,omitempty"`

	// validation
	Validation interface{} `json:"validation,omitempty"`
}

// Validate validates this post roles role ID users user ID bad request body a o1 errors
func (o *PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors) UnmarshalBinary(b []byte) error {
	var res PostRolesRoleIDUsersUserIDBadRequestBodyAO1Errors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRolesRoleIDUsersUserIDInternalServerErrorBody post roles role ID users user ID internal server error body
// swagger:model PostRolesRoleIDUsersUserIDInternalServerErrorBody
type PostRolesRoleIDUsersUserIDInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRolesRoleIDUsersUserIDInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PostRolesRoleIDUsersUserIDInternalServerErrorBodyAO0
	var postRolesRoleIDUsersUserIDInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &postRolesRoleIDUsersUserIDInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = postRolesRoleIDUsersUserIDInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRolesRoleIDUsersUserIDInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postRolesRoleIDUsersUserIDInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRolesRoleIDUsersUserIDInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post roles role ID users user ID internal server error body
func (o *PostRolesRoleIDUsersUserIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *PostRolesRoleIDUsersUserIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRolesRoleIDUsersUserIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostRolesRoleIDUsersUserIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRolesRoleIDUsersUserIDMethodNotAllowedBody post roles role ID users user ID method not allowed body
// swagger:model PostRolesRoleIDUsersUserIDMethodNotAllowedBody
type PostRolesRoleIDUsersUserIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRolesRoleIDUsersUserIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0
	var postRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRolesRoleIDUsersUserIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post roles role ID users user ID method not allowed body
func (o *PostRolesRoleIDUsersUserIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostRolesRoleIDUsersUserIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRolesRoleIDUsersUserIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostRolesRoleIDUsersUserIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRolesRoleIDUsersUserIDNotFoundBody post roles role ID users user ID not found body
// swagger:model PostRolesRoleIDUsersUserIDNotFoundBody
type PostRolesRoleIDUsersUserIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRolesRoleIDUsersUserIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// PostRolesRoleIDUsersUserIDNotFoundBodyAO0
	var postRolesRoleIDUsersUserIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &postRolesRoleIDUsersUserIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = postRolesRoleIDUsersUserIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRolesRoleIDUsersUserIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postRolesRoleIDUsersUserIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRolesRoleIDUsersUserIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post roles role ID users user ID not found body
func (o *PostRolesRoleIDUsersUserIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *PostRolesRoleIDUsersUserIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRolesRoleIDUsersUserIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostRolesRoleIDUsersUserIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRolesRoleIDUsersUserIDOKBody post roles role ID users user ID o k body
// swagger:model PostRolesRoleIDUsersUserIDOKBody
type PostRolesRoleIDUsersUserIDOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []string `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRolesRoleIDUsersUserIDOKBody) UnmarshalJSON(raw []byte) error {
	// PostRolesRoleIDUsersUserIDOKBodyAO0
	var postRolesRoleIDUsersUserIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postRolesRoleIDUsersUserIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postRolesRoleIDUsersUserIDOKBodyAO0

	// PostRolesRoleIDUsersUserIDOKBodyAO1
	var dataPostRolesRoleIDUsersUserIDOKBodyAO1 struct {
		Data []string `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPostRolesRoleIDUsersUserIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostRolesRoleIDUsersUserIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRolesRoleIDUsersUserIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postRolesRoleIDUsersUserIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRolesRoleIDUsersUserIDOKBodyAO0)

	var dataPostRolesRoleIDUsersUserIDOKBodyAO1 struct {
		Data []string `json:"data"`
	}

	dataPostRolesRoleIDUsersUserIDOKBodyAO1.Data = o.Data

	jsonDataPostRolesRoleIDUsersUserIDOKBodyAO1, errPostRolesRoleIDUsersUserIDOKBodyAO1 := swag.WriteJSON(dataPostRolesRoleIDUsersUserIDOKBodyAO1)
	if errPostRolesRoleIDUsersUserIDOKBodyAO1 != nil {
		return nil, errPostRolesRoleIDUsersUserIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostRolesRoleIDUsersUserIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post roles role ID users user ID o k body
func (o *PostRolesRoleIDUsersUserIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostRolesRoleIDUsersUserIDOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("postRolesRoleIdUsersUserIdOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostRolesRoleIDUsersUserIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRolesRoleIDUsersUserIDOKBody) UnmarshalBinary(b []byte) error {
	var res PostRolesRoleIDUsersUserIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
