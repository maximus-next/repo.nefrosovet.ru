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

// DeleteRolesRoleIDUsersUserIDHandlerFunc turns a function with the right signature into a delete roles role ID users user ID handler
type DeleteRolesRoleIDUsersUserIDHandlerFunc func(DeleteRolesRoleIDUsersUserIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRolesRoleIDUsersUserIDHandlerFunc) Handle(params DeleteRolesRoleIDUsersUserIDParams) middleware.Responder {
	return fn(params)
}

// DeleteRolesRoleIDUsersUserIDHandler interface for that can handle valid delete roles role ID users user ID params
type DeleteRolesRoleIDUsersUserIDHandler interface {
	Handle(DeleteRolesRoleIDUsersUserIDParams) middleware.Responder
}

// NewDeleteRolesRoleIDUsersUserID creates a new http.Handler for the delete roles role ID users user ID operation
func NewDeleteRolesRoleIDUsersUserID(ctx *middleware.Context, handler DeleteRolesRoleIDUsersUserIDHandler) *DeleteRolesRoleIDUsersUserID {
	return &DeleteRolesRoleIDUsersUserID{Context: ctx, Handler: handler}
}

/*DeleteRolesRoleIDUsersUserID swagger:route DELETE /roles/{roleID}/users/{userID} Role deleteRolesRoleIdUsersUserId

Удаление пользователя из роли

*/
type DeleteRolesRoleIDUsersUserID struct {
	Context *middleware.Context
	Handler DeleteRolesRoleIDUsersUserIDHandler
}

func (o *DeleteRolesRoleIDUsersUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRolesRoleIDUsersUserIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteRolesRoleIDUsersUserIDBadRequestBody delete roles role ID users user ID bad request body
// swagger:model DeleteRolesRoleIDUsersUserIDBadRequestBody
type DeleteRolesRoleIDUsersUserIDBadRequestBody struct {
	models.Error400Data

	// errors
	// Required: true
	Errors *DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors `json:"errors"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteRolesRoleIDUsersUserIDBadRequestBody) UnmarshalJSON(raw []byte) error {
	// DeleteRolesRoleIDUsersUserIDBadRequestBodyAO0
	var deleteRolesRoleIDUsersUserIDBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &deleteRolesRoleIDUsersUserIDBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = deleteRolesRoleIDUsersUserIDBadRequestBodyAO0

	// DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1
	var dataDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1 struct {
		Errors *DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}
	if err := swag.ReadJSON(raw, &dataDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1); err != nil {
		return err
	}

	o.Errors = dataDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1.Errors

	o.Message = dataDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteRolesRoleIDUsersUserIDBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	deleteRolesRoleIDUsersUserIDBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteRolesRoleIDUsersUserIDBadRequestBodyAO0)

	var dataDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1 struct {
		Errors *DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}

	dataDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1.Errors = o.Errors

	dataDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1.Message = o.Message

	jsonDataDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1, errDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1 := swag.WriteJSON(dataDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1)
	if errDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1 != nil {
		return nil, errDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataDeleteRolesRoleIDUsersUserIDBadRequestBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete roles role ID users user ID bad request body
func (o *DeleteRolesRoleIDUsersUserIDBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteRolesRoleIDUsersUserIDBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("deleteRolesRoleIdUsersUserIdBadRequest"+"."+"errors", "body", o.Errors); err != nil {
		return err
	}

	if o.Errors != nil {
		if err := o.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteRolesRoleIdUsersUserIdBadRequest" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

func (o *DeleteRolesRoleIDUsersUserIDBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteRolesRoleIdUsersUserIdBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteRolesRoleIDUsersUserIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteRolesRoleIDUsersUserIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteRolesRoleIDUsersUserIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors delete roles role ID users user ID bad request body a o1 errors
// swagger:model DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors
type DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors struct {

	// core
	Core string `json:"core,omitempty"`

	// json
	JSON string `json:"json,omitempty"`

	// validation
	Validation interface{} `json:"validation,omitempty"`
}

// Validate validates this delete roles role ID users user ID bad request body a o1 errors
func (o *DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors) UnmarshalBinary(b []byte) error {
	var res DeleteRolesRoleIDUsersUserIDBadRequestBodyAO1Errors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteRolesRoleIDUsersUserIDInternalServerErrorBody delete roles role ID users user ID internal server error body
// swagger:model DeleteRolesRoleIDUsersUserIDInternalServerErrorBody
type DeleteRolesRoleIDUsersUserIDInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteRolesRoleIDUsersUserIDInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// DeleteRolesRoleIDUsersUserIDInternalServerErrorBodyAO0
	var deleteRolesRoleIDUsersUserIDInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &deleteRolesRoleIDUsersUserIDInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = deleteRolesRoleIDUsersUserIDInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteRolesRoleIDUsersUserIDInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteRolesRoleIDUsersUserIDInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteRolesRoleIDUsersUserIDInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete roles role ID users user ID internal server error body
func (o *DeleteRolesRoleIDUsersUserIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *DeleteRolesRoleIDUsersUserIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteRolesRoleIDUsersUserIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeleteRolesRoleIDUsersUserIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteRolesRoleIDUsersUserIDMethodNotAllowedBody delete roles role ID users user ID method not allowed body
// swagger:model DeleteRolesRoleIDUsersUserIDMethodNotAllowedBody
type DeleteRolesRoleIDUsersUserIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteRolesRoleIDUsersUserIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// DeleteRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0
	var deleteRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &deleteRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = deleteRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteRolesRoleIDUsersUserIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteRolesRoleIDUsersUserIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete roles role ID users user ID method not allowed body
func (o *DeleteRolesRoleIDUsersUserIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *DeleteRolesRoleIDUsersUserIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteRolesRoleIDUsersUserIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res DeleteRolesRoleIDUsersUserIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteRolesRoleIDUsersUserIDNotFoundBody delete roles role ID users user ID not found body
// swagger:model DeleteRolesRoleIDUsersUserIDNotFoundBody
type DeleteRolesRoleIDUsersUserIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteRolesRoleIDUsersUserIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// DeleteRolesRoleIDUsersUserIDNotFoundBodyAO0
	var deleteRolesRoleIDUsersUserIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &deleteRolesRoleIDUsersUserIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = deleteRolesRoleIDUsersUserIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteRolesRoleIDUsersUserIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteRolesRoleIDUsersUserIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteRolesRoleIDUsersUserIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete roles role ID users user ID not found body
func (o *DeleteRolesRoleIDUsersUserIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *DeleteRolesRoleIDUsersUserIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteRolesRoleIDUsersUserIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteRolesRoleIDUsersUserIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteRolesRoleIDUsersUserIDOKBody delete roles role ID users user ID o k body
// swagger:model DeleteRolesRoleIDUsersUserIDOKBody
type DeleteRolesRoleIDUsersUserIDOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []interface{} `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteRolesRoleIDUsersUserIDOKBody) UnmarshalJSON(raw []byte) error {
	// DeleteRolesRoleIDUsersUserIDOKBodyAO0
	var deleteRolesRoleIDUsersUserIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &deleteRolesRoleIDUsersUserIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = deleteRolesRoleIDUsersUserIDOKBodyAO0

	// DeleteRolesRoleIDUsersUserIDOKBodyAO1
	var dataDeleteRolesRoleIDUsersUserIDOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataDeleteRolesRoleIDUsersUserIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataDeleteRolesRoleIDUsersUserIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteRolesRoleIDUsersUserIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	deleteRolesRoleIDUsersUserIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteRolesRoleIDUsersUserIDOKBodyAO0)

	var dataDeleteRolesRoleIDUsersUserIDOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}

	dataDeleteRolesRoleIDUsersUserIDOKBodyAO1.Data = o.Data

	jsonDataDeleteRolesRoleIDUsersUserIDOKBodyAO1, errDeleteRolesRoleIDUsersUserIDOKBodyAO1 := swag.WriteJSON(dataDeleteRolesRoleIDUsersUserIDOKBodyAO1)
	if errDeleteRolesRoleIDUsersUserIDOKBodyAO1 != nil {
		return nil, errDeleteRolesRoleIDUsersUserIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataDeleteRolesRoleIDUsersUserIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete roles role ID users user ID o k body
func (o *DeleteRolesRoleIDUsersUserIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteRolesRoleIDUsersUserIDOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("deleteRolesRoleIdUsersUserIdOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteRolesRoleIDUsersUserIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteRolesRoleIDUsersUserIDOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteRolesRoleIDUsersUserIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
