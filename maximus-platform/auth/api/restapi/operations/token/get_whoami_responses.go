// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetWhoamiOKCode is the HTTP code returned for type GetWhoamiOK
const GetWhoamiOKCode int = 200

/*GetWhoamiOK Инфо о токене

swagger:response getWhoamiOK
*/
type GetWhoamiOK struct {

	/*
	  In: Body
	*/
	Payload *GetWhoamiOKBody `json:"body,omitempty"`
}

// NewGetWhoamiOK creates GetWhoamiOK with default headers values
func NewGetWhoamiOK() *GetWhoamiOK {

	return &GetWhoamiOK{}
}

// WithPayload adds the payload to the get whoami o k response
func (o *GetWhoamiOK) WithPayload(payload *GetWhoamiOKBody) *GetWhoamiOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get whoami o k response
func (o *GetWhoamiOK) SetPayload(payload *GetWhoamiOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWhoamiOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetWhoamiUnauthorizedCode is the HTTP code returned for type GetWhoamiUnauthorized
const GetWhoamiUnauthorizedCode int = 401

/*GetWhoamiUnauthorized Access denied

swagger:response getWhoamiUnauthorized
*/
type GetWhoamiUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *GetWhoamiUnauthorizedBody `json:"body,omitempty"`
}

// NewGetWhoamiUnauthorized creates GetWhoamiUnauthorized with default headers values
func NewGetWhoamiUnauthorized() *GetWhoamiUnauthorized {

	return &GetWhoamiUnauthorized{}
}

// WithPayload adds the payload to the get whoami unauthorized response
func (o *GetWhoamiUnauthorized) WithPayload(payload *GetWhoamiUnauthorizedBody) *GetWhoamiUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get whoami unauthorized response
func (o *GetWhoamiUnauthorized) SetPayload(payload *GetWhoamiUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWhoamiUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetWhoamiInternalServerErrorCode is the HTTP code returned for type GetWhoamiInternalServerError
const GetWhoamiInternalServerErrorCode int = 500

/*GetWhoamiInternalServerError Internal server error

swagger:response getWhoamiInternalServerError
*/
type GetWhoamiInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetWhoamiInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetWhoamiInternalServerError creates GetWhoamiInternalServerError with default headers values
func NewGetWhoamiInternalServerError() *GetWhoamiInternalServerError {

	return &GetWhoamiInternalServerError{}
}

// WithPayload adds the payload to the get whoami internal server error response
func (o *GetWhoamiInternalServerError) WithPayload(payload *GetWhoamiInternalServerErrorBody) *GetWhoamiInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get whoami internal server error response
func (o *GetWhoamiInternalServerError) SetPayload(payload *GetWhoamiInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWhoamiInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}