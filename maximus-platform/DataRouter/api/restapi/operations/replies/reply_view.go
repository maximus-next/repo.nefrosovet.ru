// Code generated by go-swagger; DO NOT EDIT.

package replies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
)

// ReplyViewHandlerFunc turns a function with the right signature into a reply view handler
type ReplyViewHandlerFunc func(ReplyViewParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplyViewHandlerFunc) Handle(params ReplyViewParams) middleware.Responder {
	return fn(params)
}

// ReplyViewHandler interface for that can handle valid reply view params
type ReplyViewHandler interface {
	Handle(ReplyViewParams) middleware.Responder
}

// NewReplyView creates a new http.Handler for the reply view operation
func NewReplyView(ctx *middleware.Context, handler ReplyViewHandler) *ReplyView {
	return &ReplyView{Context: ctx, Handler: handler}
}

/*ReplyView swagger:route GET /replies/{replyID} Replies replyView

Информация о шаблоне ответа

*/
type ReplyView struct {
	Context *middleware.Context
	Handler ReplyViewHandler
}

func (o *ReplyView) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplyViewParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ReplyViewInternalServerErrorBody reply view internal server error body
// swagger:model ReplyViewInternalServerErrorBody
type ReplyViewInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ReplyViewInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// ReplyViewInternalServerErrorBodyAO0
	var replyViewInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &replyViewInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = replyViewInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ReplyViewInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	replyViewInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, replyViewInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reply view internal server error body
func (o *ReplyViewInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *ReplyViewInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplyViewInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ReplyViewInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ReplyViewMethodNotAllowedBody reply view method not allowed body
// swagger:model ReplyViewMethodNotAllowedBody
type ReplyViewMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ReplyViewMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// ReplyViewMethodNotAllowedBodyAO0
	var replyViewMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &replyViewMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = replyViewMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ReplyViewMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	replyViewMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, replyViewMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reply view method not allowed body
func (o *ReplyViewMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *ReplyViewMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplyViewMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res ReplyViewMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ReplyViewNotFoundBody reply view not found body
// swagger:model ReplyViewNotFoundBody
type ReplyViewNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ReplyViewNotFoundBody) UnmarshalJSON(raw []byte) error {
	// ReplyViewNotFoundBodyAO0
	var replyViewNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &replyViewNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = replyViewNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ReplyViewNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	replyViewNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, replyViewNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reply view not found body
func (o *ReplyViewNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *ReplyViewNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplyViewNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ReplyViewNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ReplyViewOKBody reply view o k body
// swagger:model ReplyViewOKBody
type ReplyViewOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ReplyViewOKBody) UnmarshalJSON(raw []byte) error {
	// ReplyViewOKBodyAO0
	var replyViewOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &replyViewOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = replyViewOKBodyAO0

	// ReplyViewOKBodyAO1
	var dataReplyViewOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataReplyViewOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataReplyViewOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ReplyViewOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	replyViewOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, replyViewOKBodyAO0)

	var dataReplyViewOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataReplyViewOKBodyAO1.Data = o.Data

	jsonDataReplyViewOKBodyAO1, errReplyViewOKBodyAO1 := swag.WriteJSON(dataReplyViewOKBodyAO1)
	if errReplyViewOKBodyAO1 != nil {
		return nil, errReplyViewOKBodyAO1
	}
	_parts = append(_parts, jsonDataReplyViewOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reply view o k body
func (o *ReplyViewOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ReplyViewOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("replyViewOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ReplyViewOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplyViewOKBody) UnmarshalBinary(b []byte) error {
	var res ReplyViewOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}