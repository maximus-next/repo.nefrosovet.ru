// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewEventCollectionParams creates a new EventCollectionParams object
// no default values defined in spec.
func NewEventCollectionParams() EventCollectionParams {

	return EventCollectionParams{}
}

// EventCollectionParams contains all the bound params for the event collection operation
// typically these are obtained from a http.Request
//
// swagger:parameters eventCollection
type EventCollectionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Канал назначения
	  In: query
	*/
	DstTopic *string
	/*Лимит выдачи
	  In: query
	*/
	Limit *int64
	/*Шаг выдачи
	  In: query
	*/
	Offset *int64
	/*Флаг ответа
	  In: query
	*/
	ReplyID *strfmt.UUID
	/*Идентификатор маршрута
	  In: query
	*/
	RouteID *strfmt.UUID
	/*Канал источника
	  In: query
	*/
	SrcTopic *string
	/*Идентификатор транзакции
	  In: query
	*/
	TransactionID *strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewEventCollectionParams() beforehand.
func (o *EventCollectionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDstTopic, qhkDstTopic, _ := qs.GetOK("dstTopic")
	if err := o.bindDstTopic(qDstTopic, qhkDstTopic, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qReplyID, qhkReplyID, _ := qs.GetOK("replyID")
	if err := o.bindReplyID(qReplyID, qhkReplyID, route.Formats); err != nil {
		res = append(res, err)
	}

	qRouteID, qhkRouteID, _ := qs.GetOK("routeID")
	if err := o.bindRouteID(qRouteID, qhkRouteID, route.Formats); err != nil {
		res = append(res, err)
	}

	qSrcTopic, qhkSrcTopic, _ := qs.GetOK("srcTopic")
	if err := o.bindSrcTopic(qSrcTopic, qhkSrcTopic, route.Formats); err != nil {
		res = append(res, err)
	}

	qTransactionID, qhkTransactionID, _ := qs.GetOK("transactionID")
	if err := o.bindTransactionID(qTransactionID, qhkTransactionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDstTopic binds and validates parameter DstTopic from query.
func (o *EventCollectionParams) bindDstTopic(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.DstTopic = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *EventCollectionParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *EventCollectionParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	return nil
}

// bindReplyID binds and validates parameter ReplyID from query.
func (o *EventCollectionParams) bindReplyID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("replyID", "query", "strfmt.UUID", raw)
	}
	o.ReplyID = (value.(*strfmt.UUID))

	if err := o.validateReplyID(formats); err != nil {
		return err
	}

	return nil
}

// validateReplyID carries on validations for parameter ReplyID
func (o *EventCollectionParams) validateReplyID(formats strfmt.Registry) error {

	if err := validate.FormatOf("replyID", "query", "uuid", o.ReplyID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindRouteID binds and validates parameter RouteID from query.
func (o *EventCollectionParams) bindRouteID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("routeID", "query", "strfmt.UUID", raw)
	}
	o.RouteID = (value.(*strfmt.UUID))

	if err := o.validateRouteID(formats); err != nil {
		return err
	}

	return nil
}

// validateRouteID carries on validations for parameter RouteID
func (o *EventCollectionParams) validateRouteID(formats strfmt.Registry) error {

	if err := validate.FormatOf("routeID", "query", "uuid", o.RouteID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindSrcTopic binds and validates parameter SrcTopic from query.
func (o *EventCollectionParams) bindSrcTopic(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.SrcTopic = &raw

	return nil
}

// bindTransactionID binds and validates parameter TransactionID from query.
func (o *EventCollectionParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("transactionID", "query", "strfmt.UUID", raw)
	}
	o.TransactionID = (value.(*strfmt.UUID))

	if err := o.validateTransactionID(formats); err != nil {
		return err
	}

	return nil
}

// validateTransactionID carries on validations for parameter TransactionID
func (o *EventCollectionParams) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.FormatOf("transactionID", "query", "uuid", o.TransactionID.String(), formats); err != nil {
		return err
	}
	return nil
}
