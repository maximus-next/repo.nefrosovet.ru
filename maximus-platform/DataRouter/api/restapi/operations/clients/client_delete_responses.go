// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ClientDeleteOKCode is the HTTP code returned for type ClientDeleteOK
const ClientDeleteOKCode int = 200

/*ClientDeleteOK Success

swagger:response clientDeleteOK
*/
type ClientDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *ClientDeleteOKBody `json:"body,omitempty"`
}

// NewClientDeleteOK creates ClientDeleteOK with default headers values
func NewClientDeleteOK() *ClientDeleteOK {

	return &ClientDeleteOK{}
}

// WithPayload adds the payload to the client delete o k response
func (o *ClientDeleteOK) WithPayload(payload *ClientDeleteOKBody) *ClientDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the client delete o k response
func (o *ClientDeleteOK) SetPayload(payload *ClientDeleteOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClientDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClientDeleteNotFoundCode is the HTTP code returned for type ClientDeleteNotFound
const ClientDeleteNotFoundCode int = 404

/*ClientDeleteNotFound Not found

swagger:response clientDeleteNotFound
*/
type ClientDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *ClientDeleteNotFoundBody `json:"body,omitempty"`
}

// NewClientDeleteNotFound creates ClientDeleteNotFound with default headers values
func NewClientDeleteNotFound() *ClientDeleteNotFound {

	return &ClientDeleteNotFound{}
}

// WithPayload adds the payload to the client delete not found response
func (o *ClientDeleteNotFound) WithPayload(payload *ClientDeleteNotFoundBody) *ClientDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the client delete not found response
func (o *ClientDeleteNotFound) SetPayload(payload *ClientDeleteNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClientDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClientDeleteMethodNotAllowedCode is the HTTP code returned for type ClientDeleteMethodNotAllowed
const ClientDeleteMethodNotAllowedCode int = 405

/*ClientDeleteMethodNotAllowed Invalid Method

swagger:response clientDeleteMethodNotAllowed
*/
type ClientDeleteMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *ClientDeleteMethodNotAllowedBody `json:"body,omitempty"`
}

// NewClientDeleteMethodNotAllowed creates ClientDeleteMethodNotAllowed with default headers values
func NewClientDeleteMethodNotAllowed() *ClientDeleteMethodNotAllowed {

	return &ClientDeleteMethodNotAllowed{}
}

// WithPayload adds the payload to the client delete method not allowed response
func (o *ClientDeleteMethodNotAllowed) WithPayload(payload *ClientDeleteMethodNotAllowedBody) *ClientDeleteMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the client delete method not allowed response
func (o *ClientDeleteMethodNotAllowed) SetPayload(payload *ClientDeleteMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClientDeleteMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClientDeleteInternalServerErrorCode is the HTTP code returned for type ClientDeleteInternalServerError
const ClientDeleteInternalServerErrorCode int = 500

/*ClientDeleteInternalServerError Internal sersver error

swagger:response clientDeleteInternalServerError
*/
type ClientDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *ClientDeleteInternalServerErrorBody `json:"body,omitempty"`
}

// NewClientDeleteInternalServerError creates ClientDeleteInternalServerError with default headers values
func NewClientDeleteInternalServerError() *ClientDeleteInternalServerError {

	return &ClientDeleteInternalServerError{}
}

// WithPayload adds the payload to the client delete internal server error response
func (o *ClientDeleteInternalServerError) WithPayload(payload *ClientDeleteInternalServerErrorBody) *ClientDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the client delete internal server error response
func (o *ClientDeleteInternalServerError) SetPayload(payload *ClientDeleteInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClientDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}