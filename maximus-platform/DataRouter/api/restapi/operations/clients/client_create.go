// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
)

// ClientCreateHandlerFunc turns a function with the right signature into a client create handler
type ClientCreateHandlerFunc func(ClientCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ClientCreateHandlerFunc) Handle(params ClientCreateParams) middleware.Responder {
	return fn(params)
}

// ClientCreateHandler interface for that can handle valid client create params
type ClientCreateHandler interface {
	Handle(ClientCreateParams) middleware.Responder
}

// NewClientCreate creates a new http.Handler for the client create operation
func NewClientCreate(ctx *middleware.Context, handler ClientCreateHandler) *ClientCreate {
	return &ClientCreate{Context: ctx, Handler: handler}
}

/*ClientCreate swagger:route POST /clients Clients clientCreate

Создание клиента

*/
type ClientCreate struct {
	Context *middleware.Context
	Handler ClientCreateHandler
}

func (o *ClientCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewClientCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ClientCreateBadRequestBody client create bad request body
// swagger:model ClientCreateBadRequestBody
type ClientCreateBadRequestBody struct {
	models.Error400Data

	ClientCreateBadRequestBodyAllOf1

	// errors
	Errors *ClientCreateBadRequestBodyAO2Errors `json:"errors,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ClientCreateBadRequestBody) UnmarshalJSON(raw []byte) error {
	// ClientCreateBadRequestBodyAO0
	var clientCreateBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &clientCreateBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = clientCreateBadRequestBodyAO0

	// ClientCreateBadRequestBodyAO1
	var clientCreateBadRequestBodyAO1 ClientCreateBadRequestBodyAllOf1
	if err := swag.ReadJSON(raw, &clientCreateBadRequestBodyAO1); err != nil {
		return err
	}
	o.ClientCreateBadRequestBodyAllOf1 = clientCreateBadRequestBodyAO1

	// ClientCreateBadRequestBodyAO2
	var dataClientCreateBadRequestBodyAO2 struct {
		Errors *ClientCreateBadRequestBodyAO2Errors `json:"errors,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataClientCreateBadRequestBodyAO2); err != nil {
		return err
	}

	o.Errors = dataClientCreateBadRequestBodyAO2.Errors

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ClientCreateBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	clientCreateBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, clientCreateBadRequestBodyAO0)

	clientCreateBadRequestBodyAO1, err := swag.WriteJSON(o.ClientCreateBadRequestBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, clientCreateBadRequestBodyAO1)

	var dataClientCreateBadRequestBodyAO2 struct {
		Errors *ClientCreateBadRequestBodyAO2Errors `json:"errors,omitempty"`
	}

	dataClientCreateBadRequestBodyAO2.Errors = o.Errors

	jsonDataClientCreateBadRequestBodyAO2, errClientCreateBadRequestBodyAO2 := swag.WriteJSON(dataClientCreateBadRequestBodyAO2)
	if errClientCreateBadRequestBodyAO2 != nil {
		return nil, errClientCreateBadRequestBodyAO2
	}
	_parts = append(_parts, jsonDataClientCreateBadRequestBodyAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this client create bad request body
func (o *ClientCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error400Data
	if err := o.Error400Data.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ClientCreateBadRequestBodyAllOf1

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClientCreateBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.Errors) { // not required
		return nil
	}

	if o.Errors != nil {
		if err := o.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientCreateBadRequest" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClientCreateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClientCreateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ClientCreateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ClientCreateBadRequestBodyAO2Errors client create bad request body a o2 errors
// swagger:model ClientCreateBadRequestBodyAO2Errors
type ClientCreateBadRequestBodyAO2Errors struct {

	// validation
	Validation *ClientCreateBadRequestBodyAO2ErrorsValidation `json:"validation,omitempty"`
}

// Validate validates this client create bad request body a o2 errors
func (o *ClientCreateBadRequestBodyAO2Errors) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValidation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClientCreateBadRequestBodyAO2Errors) validateValidation(formats strfmt.Registry) error {

	if swag.IsZero(o.Validation) { // not required
		return nil
	}

	if o.Validation != nil {
		if err := o.Validation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientCreateBadRequest" + "." + "errors" + "." + "validation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClientCreateBadRequestBodyAO2Errors) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClientCreateBadRequestBodyAO2Errors) UnmarshalBinary(b []byte) error {
	var res ClientCreateBadRequestBodyAO2Errors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ClientCreateBadRequestBodyAO2ErrorsValidation client create bad request body a o2 errors validation
// swagger:model ClientCreateBadRequestBodyAO2ErrorsValidation
type ClientCreateBadRequestBodyAO2ErrorsValidation struct {

	// ID
	// Enum: [string format unique]
	ID string `json:"ID,omitempty"`

	// password
	// Enum: [string required]
	Password string `json:"password,omitempty"`

	// ttl
	// Enum: [int]
	TTL string `json:"ttl,omitempty"`

	// username
	// Enum: [string required]
	Username string `json:"username,omitempty"`
}

// Validate validates this client create bad request body a o2 errors validation
func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clientCreateBadRequestBodyAO2ErrorsValidationTypeIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","format","unique"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientCreateBadRequestBodyAO2ErrorsValidationTypeIDPropEnum = append(clientCreateBadRequestBodyAO2ErrorsValidationTypeIDPropEnum, v)
	}
}

const (

	// ClientCreateBadRequestBodyAO2ErrorsValidationIDString captures enum value "string"
	ClientCreateBadRequestBodyAO2ErrorsValidationIDString string = "string"

	// ClientCreateBadRequestBodyAO2ErrorsValidationIDFormat captures enum value "format"
	ClientCreateBadRequestBodyAO2ErrorsValidationIDFormat string = "format"

	// ClientCreateBadRequestBodyAO2ErrorsValidationIDUnique captures enum value "unique"
	ClientCreateBadRequestBodyAO2ErrorsValidationIDUnique string = "unique"
)

// prop value enum
func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) validateIDEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clientCreateBadRequestBodyAO2ErrorsValidationTypeIDPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	// value enum
	if err := o.validateIDEnum("clientCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"ID", "body", o.ID); err != nil {
		return err
	}

	return nil
}

var clientCreateBadRequestBodyAO2ErrorsValidationTypePasswordPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientCreateBadRequestBodyAO2ErrorsValidationTypePasswordPropEnum = append(clientCreateBadRequestBodyAO2ErrorsValidationTypePasswordPropEnum, v)
	}
}

const (

	// ClientCreateBadRequestBodyAO2ErrorsValidationPasswordString captures enum value "string"
	ClientCreateBadRequestBodyAO2ErrorsValidationPasswordString string = "string"

	// ClientCreateBadRequestBodyAO2ErrorsValidationPasswordRequired captures enum value "required"
	ClientCreateBadRequestBodyAO2ErrorsValidationPasswordRequired string = "required"
)

// prop value enum
func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) validatePasswordEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clientCreateBadRequestBodyAO2ErrorsValidationTypePasswordPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(o.Password) { // not required
		return nil
	}

	// value enum
	if err := o.validatePasswordEnum("clientCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

var clientCreateBadRequestBodyAO2ErrorsValidationTypeTTLPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["int"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientCreateBadRequestBodyAO2ErrorsValidationTypeTTLPropEnum = append(clientCreateBadRequestBodyAO2ErrorsValidationTypeTTLPropEnum, v)
	}
}

const (

	// ClientCreateBadRequestBodyAO2ErrorsValidationTTLInt captures enum value "int"
	ClientCreateBadRequestBodyAO2ErrorsValidationTTLInt string = "int"
)

// prop value enum
func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) validateTTLEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clientCreateBadRequestBodyAO2ErrorsValidationTypeTTLPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) validateTTL(formats strfmt.Registry) error {

	if swag.IsZero(o.TTL) { // not required
		return nil
	}

	// value enum
	if err := o.validateTTLEnum("clientCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"ttl", "body", o.TTL); err != nil {
		return err
	}

	return nil
}

var clientCreateBadRequestBodyAO2ErrorsValidationTypeUsernamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientCreateBadRequestBodyAO2ErrorsValidationTypeUsernamePropEnum = append(clientCreateBadRequestBodyAO2ErrorsValidationTypeUsernamePropEnum, v)
	}
}

const (

	// ClientCreateBadRequestBodyAO2ErrorsValidationUsernameString captures enum value "string"
	ClientCreateBadRequestBodyAO2ErrorsValidationUsernameString string = "string"

	// ClientCreateBadRequestBodyAO2ErrorsValidationUsernameRequired captures enum value "required"
	ClientCreateBadRequestBodyAO2ErrorsValidationUsernameRequired string = "required"
)

// prop value enum
func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) validateUsernameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clientCreateBadRequestBodyAO2ErrorsValidationTypeUsernamePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) validateUsername(formats strfmt.Registry) error {

	if swag.IsZero(o.Username) { // not required
		return nil
	}

	// value enum
	if err := o.validateUsernameEnum("clientCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClientCreateBadRequestBodyAO2ErrorsValidation) UnmarshalBinary(b []byte) error {
	var res ClientCreateBadRequestBodyAO2ErrorsValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ClientCreateBadRequestBodyAllOf1 client create bad request body all of1
// swagger:model ClientCreateBadRequestBodyAllOf1
type ClientCreateBadRequestBodyAllOf1 interface{}

// ClientCreateBody client create body
// swagger:model ClientCreateBody
type ClientCreateBody struct {

	// Идентификатор клиента
	// Format: uuid
	ID strfmt.UUID `json:"ID,omitempty"`

	models.ClientObject

	// Пароль клиента
	// Required: true
	// Min Length: 1
	Password *string `json:"password"`

	// Время жизни клиента в секундах.
	TTL int64 `json:"ttl,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ClientCreateBody) UnmarshalJSON(raw []byte) error {
	// ClientCreateParamsBodyAO0
	var dataClientCreateParamsBodyAO0 struct {
		ID strfmt.UUID `json:"ID,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataClientCreateParamsBodyAO0); err != nil {
		return err
	}

	o.ID = dataClientCreateParamsBodyAO0.ID

	// ClientCreateParamsBodyAO1
	var clientCreateParamsBodyAO1 models.ClientObject
	if err := swag.ReadJSON(raw, &clientCreateParamsBodyAO1); err != nil {
		return err
	}
	o.ClientObject = clientCreateParamsBodyAO1

	// ClientCreateParamsBodyAO2
	var dataClientCreateParamsBodyAO2 struct {
		Password *string `json:"password"`

		TTL int64 `json:"ttl,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataClientCreateParamsBodyAO2); err != nil {
		return err
	}

	o.Password = dataClientCreateParamsBodyAO2.Password

	o.TTL = dataClientCreateParamsBodyAO2.TTL

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ClientCreateBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	var dataClientCreateParamsBodyAO0 struct {
		ID strfmt.UUID `json:"ID,omitempty"`
	}

	dataClientCreateParamsBodyAO0.ID = o.ID

	jsonDataClientCreateParamsBodyAO0, errClientCreateParamsBodyAO0 := swag.WriteJSON(dataClientCreateParamsBodyAO0)
	if errClientCreateParamsBodyAO0 != nil {
		return nil, errClientCreateParamsBodyAO0
	}
	_parts = append(_parts, jsonDataClientCreateParamsBodyAO0)

	clientCreateParamsBodyAO1, err := swag.WriteJSON(o.ClientObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, clientCreateParamsBodyAO1)

	var dataClientCreateParamsBodyAO2 struct {
		Password *string `json:"password"`

		TTL int64 `json:"ttl,omitempty"`
	}

	dataClientCreateParamsBodyAO2.Password = o.Password

	dataClientCreateParamsBodyAO2.TTL = o.TTL

	jsonDataClientCreateParamsBodyAO2, errClientCreateParamsBodyAO2 := swag.WriteJSON(dataClientCreateParamsBodyAO2)
	if errClientCreateParamsBodyAO2 != nil {
		return nil, errClientCreateParamsBodyAO2
	}
	_parts = append(_parts, jsonDataClientCreateParamsBodyAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this client create body
func (o *ClientCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with models.ClientObject
	if err := o.ClientObject.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClientCreateBody) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"ID", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ClientCreateBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"password", "body", string(*o.Password), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClientCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClientCreateBody) UnmarshalBinary(b []byte) error {
	var res ClientCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ClientCreateInternalServerErrorBody client create internal server error body
// swagger:model ClientCreateInternalServerErrorBody
type ClientCreateInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ClientCreateInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// ClientCreateInternalServerErrorBodyAO0
	var clientCreateInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &clientCreateInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = clientCreateInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ClientCreateInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	clientCreateInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, clientCreateInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this client create internal server error body
func (o *ClientCreateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *ClientCreateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClientCreateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ClientCreateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ClientCreateMethodNotAllowedBody client create method not allowed body
// swagger:model ClientCreateMethodNotAllowedBody
type ClientCreateMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ClientCreateMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// ClientCreateMethodNotAllowedBodyAO0
	var clientCreateMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &clientCreateMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = clientCreateMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ClientCreateMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	clientCreateMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, clientCreateMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this client create method not allowed body
func (o *ClientCreateMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *ClientCreateMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClientCreateMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res ClientCreateMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ClientCreateOKBody client create o k body
// swagger:model ClientCreateOKBody
type ClientCreateOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ClientCreateOKBody) UnmarshalJSON(raw []byte) error {
	// ClientCreateOKBodyAO0
	var clientCreateOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &clientCreateOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = clientCreateOKBodyAO0

	// ClientCreateOKBodyAO1
	var dataClientCreateOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataClientCreateOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataClientCreateOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ClientCreateOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	clientCreateOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, clientCreateOKBodyAO0)

	var dataClientCreateOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataClientCreateOKBodyAO1.Data = o.Data

	jsonDataClientCreateOKBodyAO1, errClientCreateOKBodyAO1 := swag.WriteJSON(dataClientCreateOKBodyAO1)
	if errClientCreateOKBodyAO1 != nil {
		return nil, errClientCreateOKBodyAO1
	}
	_parts = append(_parts, jsonDataClientCreateOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this client create o k body
func (o *ClientCreateOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ClientCreateOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("clientCreateOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClientCreateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClientCreateOKBody) UnmarshalBinary(b []byte) error {
	var res ClientCreateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}