// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// EventCollectionOKCode is the HTTP code returned for type EventCollectionOK
const EventCollectionOKCode int = 200

/*EventCollectionOK Коллекция событий

swagger:response eventCollectionOK
*/
type EventCollectionOK struct {

	/*
	  In: Body
	*/
	Payload *EventCollectionOKBody `json:"body,omitempty"`
}

// NewEventCollectionOK creates EventCollectionOK with default headers values
func NewEventCollectionOK() *EventCollectionOK {

	return &EventCollectionOK{}
}

// WithPayload adds the payload to the event collection o k response
func (o *EventCollectionOK) WithPayload(payload *EventCollectionOKBody) *EventCollectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the event collection o k response
func (o *EventCollectionOK) SetPayload(payload *EventCollectionOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EventCollectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EventCollectionNotFoundCode is the HTTP code returned for type EventCollectionNotFound
const EventCollectionNotFoundCode int = 404

/*EventCollectionNotFound Not found

swagger:response eventCollectionNotFound
*/
type EventCollectionNotFound struct {

	/*
	  In: Body
	*/
	Payload *EventCollectionNotFoundBody `json:"body,omitempty"`
}

// NewEventCollectionNotFound creates EventCollectionNotFound with default headers values
func NewEventCollectionNotFound() *EventCollectionNotFound {

	return &EventCollectionNotFound{}
}

// WithPayload adds the payload to the event collection not found response
func (o *EventCollectionNotFound) WithPayload(payload *EventCollectionNotFoundBody) *EventCollectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the event collection not found response
func (o *EventCollectionNotFound) SetPayload(payload *EventCollectionNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EventCollectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EventCollectionMethodNotAllowedCode is the HTTP code returned for type EventCollectionMethodNotAllowed
const EventCollectionMethodNotAllowedCode int = 405

/*EventCollectionMethodNotAllowed Invalid Method

swagger:response eventCollectionMethodNotAllowed
*/
type EventCollectionMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *EventCollectionMethodNotAllowedBody `json:"body,omitempty"`
}

// NewEventCollectionMethodNotAllowed creates EventCollectionMethodNotAllowed with default headers values
func NewEventCollectionMethodNotAllowed() *EventCollectionMethodNotAllowed {

	return &EventCollectionMethodNotAllowed{}
}

// WithPayload adds the payload to the event collection method not allowed response
func (o *EventCollectionMethodNotAllowed) WithPayload(payload *EventCollectionMethodNotAllowedBody) *EventCollectionMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the event collection method not allowed response
func (o *EventCollectionMethodNotAllowed) SetPayload(payload *EventCollectionMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EventCollectionMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EventCollectionInternalServerErrorCode is the HTTP code returned for type EventCollectionInternalServerError
const EventCollectionInternalServerErrorCode int = 500

/*EventCollectionInternalServerError Internal sersver error

swagger:response eventCollectionInternalServerError
*/
type EventCollectionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *EventCollectionInternalServerErrorBody `json:"body,omitempty"`
}

// NewEventCollectionInternalServerError creates EventCollectionInternalServerError with default headers values
func NewEventCollectionInternalServerError() *EventCollectionInternalServerError {

	return &EventCollectionInternalServerError{}
}

// WithPayload adds the payload to the event collection internal server error response
func (o *EventCollectionInternalServerError) WithPayload(payload *EventCollectionInternalServerErrorBody) *EventCollectionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the event collection internal server error response
func (o *EventCollectionInternalServerError) SetPayload(payload *EventCollectionInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EventCollectionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
