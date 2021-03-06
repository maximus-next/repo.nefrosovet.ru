// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PasswordRecoveryOKCode is the HTTP code returned for type PasswordRecoveryOK
const PasswordRecoveryOKCode int = 200

/*PasswordRecoveryOK SUCCESS

swagger:response passwordRecoveryOK
*/
type PasswordRecoveryOK struct {

	/*
	  In: Body
	*/
	Payload *PasswordRecoveryOKBody `json:"body,omitempty"`
}

// NewPasswordRecoveryOK creates PasswordRecoveryOK with default headers values
func NewPasswordRecoveryOK() *PasswordRecoveryOK {

	return &PasswordRecoveryOK{}
}

// WithPayload adds the payload to the password recovery o k response
func (o *PasswordRecoveryOK) WithPayload(payload *PasswordRecoveryOKBody) *PasswordRecoveryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the password recovery o k response
func (o *PasswordRecoveryOK) SetPayload(payload *PasswordRecoveryOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PasswordRecoveryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PasswordRecoveryBadRequestCode is the HTTP code returned for type PasswordRecoveryBadRequest
const PasswordRecoveryBadRequestCode int = 400

/*PasswordRecoveryBadRequest Validation error

swagger:response passwordRecoveryBadRequest
*/
type PasswordRecoveryBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PasswordRecoveryBadRequestBody `json:"body,omitempty"`
}

// NewPasswordRecoveryBadRequest creates PasswordRecoveryBadRequest with default headers values
func NewPasswordRecoveryBadRequest() *PasswordRecoveryBadRequest {

	return &PasswordRecoveryBadRequest{}
}

// WithPayload adds the payload to the password recovery bad request response
func (o *PasswordRecoveryBadRequest) WithPayload(payload *PasswordRecoveryBadRequestBody) *PasswordRecoveryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the password recovery bad request response
func (o *PasswordRecoveryBadRequest) SetPayload(payload *PasswordRecoveryBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PasswordRecoveryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PasswordRecoveryNotFoundCode is the HTTP code returned for type PasswordRecoveryNotFound
const PasswordRecoveryNotFoundCode int = 404

/*PasswordRecoveryNotFound Not found

swagger:response passwordRecoveryNotFound
*/
type PasswordRecoveryNotFound struct {

	/*
	  In: Body
	*/
	Payload *PasswordRecoveryNotFoundBody `json:"body,omitempty"`
}

// NewPasswordRecoveryNotFound creates PasswordRecoveryNotFound with default headers values
func NewPasswordRecoveryNotFound() *PasswordRecoveryNotFound {

	return &PasswordRecoveryNotFound{}
}

// WithPayload adds the payload to the password recovery not found response
func (o *PasswordRecoveryNotFound) WithPayload(payload *PasswordRecoveryNotFoundBody) *PasswordRecoveryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the password recovery not found response
func (o *PasswordRecoveryNotFound) SetPayload(payload *PasswordRecoveryNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PasswordRecoveryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PasswordRecoveryMethodNotAllowedCode is the HTTP code returned for type PasswordRecoveryMethodNotAllowed
const PasswordRecoveryMethodNotAllowedCode int = 405

/*PasswordRecoveryMethodNotAllowed Invalid Method

swagger:response passwordRecoveryMethodNotAllowed
*/
type PasswordRecoveryMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PasswordRecoveryMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPasswordRecoveryMethodNotAllowed creates PasswordRecoveryMethodNotAllowed with default headers values
func NewPasswordRecoveryMethodNotAllowed() *PasswordRecoveryMethodNotAllowed {

	return &PasswordRecoveryMethodNotAllowed{}
}

// WithPayload adds the payload to the password recovery method not allowed response
func (o *PasswordRecoveryMethodNotAllowed) WithPayload(payload *PasswordRecoveryMethodNotAllowedBody) *PasswordRecoveryMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the password recovery method not allowed response
func (o *PasswordRecoveryMethodNotAllowed) SetPayload(payload *PasswordRecoveryMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PasswordRecoveryMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PasswordRecoveryInternalServerErrorCode is the HTTP code returned for type PasswordRecoveryInternalServerError
const PasswordRecoveryInternalServerErrorCode int = 500

/*PasswordRecoveryInternalServerError Internal server error

swagger:response passwordRecoveryInternalServerError
*/
type PasswordRecoveryInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PasswordRecoveryInternalServerErrorBody `json:"body,omitempty"`
}

// NewPasswordRecoveryInternalServerError creates PasswordRecoveryInternalServerError with default headers values
func NewPasswordRecoveryInternalServerError() *PasswordRecoveryInternalServerError {

	return &PasswordRecoveryInternalServerError{}
}

// WithPayload adds the payload to the password recovery internal server error response
func (o *PasswordRecoveryInternalServerError) WithPayload(payload *PasswordRecoveryInternalServerErrorBody) *PasswordRecoveryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the password recovery internal server error response
func (o *PasswordRecoveryInternalServerError) SetPayload(payload *PasswordRecoveryInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PasswordRecoveryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
