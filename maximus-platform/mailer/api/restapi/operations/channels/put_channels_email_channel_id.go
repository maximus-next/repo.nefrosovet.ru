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

// PutChannelsEmailChannelIDHandlerFunc turns a function with the right signature into a put channels email channel ID handler
type PutChannelsEmailChannelIDHandlerFunc func(PutChannelsEmailChannelIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutChannelsEmailChannelIDHandlerFunc) Handle(params PutChannelsEmailChannelIDParams) middleware.Responder {
	return fn(params)
}

// PutChannelsEmailChannelIDHandler interface for that can handle valid put channels email channel ID params
type PutChannelsEmailChannelIDHandler interface {
	Handle(PutChannelsEmailChannelIDParams) middleware.Responder
}

// NewPutChannelsEmailChannelID creates a new http.Handler for the put channels email channel ID operation
func NewPutChannelsEmailChannelID(ctx *middleware.Context, handler PutChannelsEmailChannelIDHandler) *PutChannelsEmailChannelID {
	return &PutChannelsEmailChannelID{Context: ctx, Handler: handler}
}

/*PutChannelsEmailChannelID swagger:route PUT /channels/email/{channelID} Channels putChannelsEmailChannelId

Изменение e-mail канала

*/
type PutChannelsEmailChannelID struct {
	Context *middleware.Context
	Handler PutChannelsEmailChannelIDHandler
}

func (o *PutChannelsEmailChannelID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutChannelsEmailChannelIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutChannelsEmailChannelIDBadRequestBody put channels email channel ID bad request body
// swagger:model PutChannelsEmailChannelIDBadRequestBody
type PutChannelsEmailChannelIDBadRequestBody struct {
	models.Error400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutChannelsEmailChannelIDBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PutChannelsEmailChannelIDBadRequestBodyAO0
	var putChannelsEmailChannelIDBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &putChannelsEmailChannelIDBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = putChannelsEmailChannelIDBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutChannelsEmailChannelIDBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	putChannelsEmailChannelIDBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putChannelsEmailChannelIDBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put channels email channel ID bad request body
func (o *PutChannelsEmailChannelIDBadRequestBody) Validate(formats strfmt.Registry) error {
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
func (o *PutChannelsEmailChannelIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutChannelsEmailChannelIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutChannelsEmailChannelIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutChannelsEmailChannelIDBody put channels email channel ID body
// swagger:model PutChannelsEmailChannelIDBody
type PutChannelsEmailChannelIDBody struct {
	models.ChannelParamsMail

	PutChannelsEmailChannelIDParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutChannelsEmailChannelIDBody) UnmarshalJSON(raw []byte) error {
	// PutChannelsEmailChannelIDParamsBodyAO0
	var putChannelsEmailChannelIDParamsBodyAO0 models.ChannelParamsMail
	if err := swag.ReadJSON(raw, &putChannelsEmailChannelIDParamsBodyAO0); err != nil {
		return err
	}
	o.ChannelParamsMail = putChannelsEmailChannelIDParamsBodyAO0

	// PutChannelsEmailChannelIDParamsBodyAO1
	var putChannelsEmailChannelIDParamsBodyAO1 PutChannelsEmailChannelIDParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &putChannelsEmailChannelIDParamsBodyAO1); err != nil {
		return err
	}
	o.PutChannelsEmailChannelIDParamsBodyAllOf1 = putChannelsEmailChannelIDParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutChannelsEmailChannelIDBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putChannelsEmailChannelIDParamsBodyAO0, err := swag.WriteJSON(o.ChannelParamsMail)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putChannelsEmailChannelIDParamsBodyAO0)

	putChannelsEmailChannelIDParamsBodyAO1, err := swag.WriteJSON(o.PutChannelsEmailChannelIDParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putChannelsEmailChannelIDParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put channels email channel ID body
func (o *PutChannelsEmailChannelIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ChannelParamsMail
	if err := o.ChannelParamsMail.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PutChannelsEmailChannelIDParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PutChannelsEmailChannelIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutChannelsEmailChannelIDBody) UnmarshalBinary(b []byte) error {
	var res PutChannelsEmailChannelIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutChannelsEmailChannelIDForbiddenBody put channels email channel ID forbidden body
// swagger:model PutChannelsEmailChannelIDForbiddenBody
type PutChannelsEmailChannelIDForbiddenBody struct {
	models.Error403Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutChannelsEmailChannelIDForbiddenBody) UnmarshalJSON(raw []byte) error {
	// PutChannelsEmailChannelIDForbiddenBodyAO0
	var putChannelsEmailChannelIDForbiddenBodyAO0 models.Error403Data
	if err := swag.ReadJSON(raw, &putChannelsEmailChannelIDForbiddenBodyAO0); err != nil {
		return err
	}
	o.Error403Data = putChannelsEmailChannelIDForbiddenBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutChannelsEmailChannelIDForbiddenBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	putChannelsEmailChannelIDForbiddenBodyAO0, err := swag.WriteJSON(o.Error403Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putChannelsEmailChannelIDForbiddenBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put channels email channel ID forbidden body
func (o *PutChannelsEmailChannelIDForbiddenBody) Validate(formats strfmt.Registry) error {
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
func (o *PutChannelsEmailChannelIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutChannelsEmailChannelIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res PutChannelsEmailChannelIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutChannelsEmailChannelIDMethodNotAllowedBody put channels email channel ID method not allowed body
// swagger:model PutChannelsEmailChannelIDMethodNotAllowedBody
type PutChannelsEmailChannelIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutChannelsEmailChannelIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PutChannelsEmailChannelIDMethodNotAllowedBodyAO0
	var putChannelsEmailChannelIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &putChannelsEmailChannelIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = putChannelsEmailChannelIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutChannelsEmailChannelIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	putChannelsEmailChannelIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putChannelsEmailChannelIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put channels email channel ID method not allowed body
func (o *PutChannelsEmailChannelIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PutChannelsEmailChannelIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutChannelsEmailChannelIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PutChannelsEmailChannelIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutChannelsEmailChannelIDNotFoundBody put channels email channel ID not found body
// swagger:model PutChannelsEmailChannelIDNotFoundBody
type PutChannelsEmailChannelIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutChannelsEmailChannelIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// PutChannelsEmailChannelIDNotFoundBodyAO0
	var putChannelsEmailChannelIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &putChannelsEmailChannelIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = putChannelsEmailChannelIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutChannelsEmailChannelIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	putChannelsEmailChannelIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putChannelsEmailChannelIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put channels email channel ID not found body
func (o *PutChannelsEmailChannelIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *PutChannelsEmailChannelIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutChannelsEmailChannelIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutChannelsEmailChannelIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutChannelsEmailChannelIDOKBody put channels email channel ID o k body
// swagger:model PutChannelsEmailChannelIDOKBody
type PutChannelsEmailChannelIDOKBody struct {
	models.SuccessData

	// data
	Data []*DataItemEmail `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutChannelsEmailChannelIDOKBody) UnmarshalJSON(raw []byte) error {
	// PutChannelsEmailChannelIDOKBodyAO0
	var putChannelsEmailChannelIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &putChannelsEmailChannelIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = putChannelsEmailChannelIDOKBodyAO0

	// PutChannelsEmailChannelIDOKBodyAO1
	var dataPutChannelsEmailChannelIDOKBodyAO1 struct {
		Data []*DataItemEmail `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPutChannelsEmailChannelIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutChannelsEmailChannelIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutChannelsEmailChannelIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putChannelsEmailChannelIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putChannelsEmailChannelIDOKBodyAO0)

	var dataPutChannelsEmailChannelIDOKBodyAO1 struct {
		Data []*DataItemEmail `json:"data"`
	}

	dataPutChannelsEmailChannelIDOKBodyAO1.Data = o.Data

	jsonDataPutChannelsEmailChannelIDOKBodyAO1, errPutChannelsEmailChannelIDOKBodyAO1 := swag.WriteJSON(dataPutChannelsEmailChannelIDOKBodyAO1)
	if errPutChannelsEmailChannelIDOKBodyAO1 != nil {
		return nil, errPutChannelsEmailChannelIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutChannelsEmailChannelIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put channels email channel ID o k body
func (o *PutChannelsEmailChannelIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutChannelsEmailChannelIDOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("putChannelsEmailChannelIdOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutChannelsEmailChannelIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutChannelsEmailChannelIDOKBody) UnmarshalBinary(b []byte) error {
	var res PutChannelsEmailChannelIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutChannelsEmailChannelIDParamsBodyAllOf1 put channels email channel ID params body all of1
// swagger:model PutChannelsEmailChannelIDParamsBodyAllOf1
type PutChannelsEmailChannelIDParamsBodyAllOf1 interface{}