// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "repo.nefrosovet.ru/maximus-platform/mailer/api/models"
)

// PostChannelsLocalSmsHandlerFunc turns a function with the right signature into a post channels local sms handler
type PostChannelsLocalSmsHandlerFunc func(PostChannelsLocalSmsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostChannelsLocalSmsHandlerFunc) Handle(params PostChannelsLocalSmsParams) middleware.Responder {
	return fn(params)
}

// PostChannelsLocalSmsHandler interface for that can handle valid post channels local sms params
type PostChannelsLocalSmsHandler interface {
	Handle(PostChannelsLocalSmsParams) middleware.Responder
}

// NewPostChannelsLocalSms creates a new http.Handler for the post channels local sms operation
func NewPostChannelsLocalSms(ctx *middleware.Context, handler PostChannelsLocalSmsHandler) *PostChannelsLocalSms {
	return &PostChannelsLocalSms{Context: ctx, Handler: handler}
}

/*PostChannelsLocalSms swagger:route POST /channels/local_sms Channels postChannelsLocalSms

Создание Local SMS канала

*/
type PostChannelsLocalSms struct {
	Context *middleware.Context
	Handler PostChannelsLocalSmsHandler
}

func (o *PostChannelsLocalSms) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostChannelsLocalSmsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostChannelsLocalSmsBadRequestBody post channels local sms bad request body
// swagger:model PostChannelsLocalSmsBadRequestBody
type PostChannelsLocalSmsBadRequestBody struct {
	models.Error400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostChannelsLocalSmsBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostChannelsLocalSmsBadRequestBodyAO0
	var postChannelsLocalSmsBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &postChannelsLocalSmsBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = postChannelsLocalSmsBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostChannelsLocalSmsBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postChannelsLocalSmsBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsLocalSmsBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post channels local sms bad request body
func (o *PostChannelsLocalSmsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error400Data
	if err := o.Error400Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostChannelsLocalSmsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChannelsLocalSmsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostChannelsLocalSmsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChannelsLocalSmsBody post channels local sms body
// swagger:model PostChannelsLocalSmsBody
type PostChannelsLocalSmsBody struct {
	models.ChannelParamsLocalSms

	PostChannelsLocalSmsParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostChannelsLocalSmsBody) UnmarshalJSON(raw []byte) error {
	// PostChannelsLocalSmsParamsBodyAO0
	var postChannelsLocalSmsParamsBodyAO0 models.ChannelParamsLocalSms
	if err := swag.ReadJSON(raw, &postChannelsLocalSmsParamsBodyAO0); err != nil {
		return err
	}
	o.ChannelParamsLocalSms = postChannelsLocalSmsParamsBodyAO0

	// PostChannelsLocalSmsParamsBodyAO1
	var postChannelsLocalSmsParamsBodyAO1 PostChannelsLocalSmsParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &postChannelsLocalSmsParamsBodyAO1); err != nil {
		return err
	}
	o.PostChannelsLocalSmsParamsBodyAllOf1 = postChannelsLocalSmsParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostChannelsLocalSmsBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postChannelsLocalSmsParamsBodyAO0, err := swag.WriteJSON(o.ChannelParamsLocalSms)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsLocalSmsParamsBodyAO0)

	postChannelsLocalSmsParamsBodyAO1, err := swag.WriteJSON(o.PostChannelsLocalSmsParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsLocalSmsParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post channels local sms body
func (o *PostChannelsLocalSmsBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ChannelParamsLocalSms
	if err := o.ChannelParamsLocalSms.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PostChannelsLocalSmsParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostChannelsLocalSmsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChannelsLocalSmsBody) UnmarshalBinary(b []byte) error {
	var res PostChannelsLocalSmsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChannelsLocalSmsForbiddenBody post channels local sms forbidden body
// swagger:model PostChannelsLocalSmsForbiddenBody
type PostChannelsLocalSmsForbiddenBody struct {
	models.Error403Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostChannelsLocalSmsForbiddenBody) UnmarshalJSON(raw []byte) error {
	// PostChannelsLocalSmsForbiddenBodyAO0
	var postChannelsLocalSmsForbiddenBodyAO0 models.Error403Data
	if err := swag.ReadJSON(raw, &postChannelsLocalSmsForbiddenBodyAO0); err != nil {
		return err
	}
	o.Error403Data = postChannelsLocalSmsForbiddenBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostChannelsLocalSmsForbiddenBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postChannelsLocalSmsForbiddenBodyAO0, err := swag.WriteJSON(o.Error403Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsLocalSmsForbiddenBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post channels local sms forbidden body
func (o *PostChannelsLocalSmsForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error403Data
	if err := o.Error403Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostChannelsLocalSmsForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChannelsLocalSmsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res PostChannelsLocalSmsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChannelsLocalSmsMethodNotAllowedBody post channels local sms method not allowed body
// swagger:model PostChannelsLocalSmsMethodNotAllowedBody
type PostChannelsLocalSmsMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostChannelsLocalSmsMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostChannelsLocalSmsMethodNotAllowedBodyAO0
	var postChannelsLocalSmsMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postChannelsLocalSmsMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postChannelsLocalSmsMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostChannelsLocalSmsMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postChannelsLocalSmsMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsLocalSmsMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post channels local sms method not allowed body
func (o *PostChannelsLocalSmsMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostChannelsLocalSmsMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChannelsLocalSmsMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostChannelsLocalSmsMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChannelsLocalSmsOKBody post channels local sms o k body
// swagger:model PostChannelsLocalSmsOKBody
type PostChannelsLocalSmsOKBody struct {
	models.SuccessData

	// data
	Data []*DataItemLocalSms `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostChannelsLocalSmsOKBody) UnmarshalJSON(raw []byte) error {
	// PostChannelsLocalSmsOKBodyAO0
	var postChannelsLocalSmsOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postChannelsLocalSmsOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postChannelsLocalSmsOKBodyAO0

	// PostChannelsLocalSmsOKBodyAO1
	var dataPostChannelsLocalSmsOKBodyAO1 struct {
		Data []*DataItemLocalSms `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPostChannelsLocalSmsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostChannelsLocalSmsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostChannelsLocalSmsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postChannelsLocalSmsOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsLocalSmsOKBodyAO0)

	var dataPostChannelsLocalSmsOKBodyAO1 struct {
		Data []*DataItemLocalSms `json:"data"`
	}

	dataPostChannelsLocalSmsOKBodyAO1.Data = o.Data

	jsonDataPostChannelsLocalSmsOKBodyAO1, errPostChannelsLocalSmsOKBodyAO1 := swag.WriteJSON(dataPostChannelsLocalSmsOKBodyAO1)
	if errPostChannelsLocalSmsOKBodyAO1 != nil {
		return nil, errPostChannelsLocalSmsOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostChannelsLocalSmsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post channels local sms o k body
func (o *PostChannelsLocalSmsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostChannelsLocalSmsOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("postChannelsLocalSmsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostChannelsLocalSmsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChannelsLocalSmsOKBody) UnmarshalBinary(b []byte) error {
	var res PostChannelsLocalSmsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChannelsLocalSmsParamsBodyAllOf1 post channels local sms params body all of1
// swagger:model PostChannelsLocalSmsParamsBodyAllOf1
type PostChannelsLocalSmsParamsBodyAllOf1 interface{}

// data-item-local-sms data item local sms
// swagger:model data-item-local-sms
type DataItemLocalSms struct {

	// Идентификатор канала
	ID string `json:"ID,omitempty"`

	// Тип канала
	Type string `json:"type,omitempty"`

	models.ChannelParamsLocalSms
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DataItemLocalSms) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID string `json:"ID,omitempty"`

		Type string `json:"type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.ID = dataAO0.ID

	o.Type = dataAO0.Type

	// AO1
	var aO1 models.ChannelParamsLocalSms
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	o.ChannelParamsLocalSms = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DataItemLocalSms) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID string `json:"ID,omitempty"`

		Type string `json:"type,omitempty"`
	}

	dataAO0.ID = o.ID

	dataAO0.Type = o.Type

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(o.ChannelParamsLocalSms)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this data item local sms
func (o *DataItemLocalSms) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ChannelParamsLocalSms
	if err := o.ChannelParamsLocalSms.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DataItemLocalSms) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DataItemLocalSms) UnmarshalBinary(b []byte) error {
	var res DataItemLocalSms
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
