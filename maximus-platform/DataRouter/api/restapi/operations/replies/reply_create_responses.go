// Code generated by go-swagger; DO NOT EDIT.

package replies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ReplyCreateOKCode is the HTTP code returned for type ReplyCreateOK
const ReplyCreateOKCode int = 200

/*ReplyCreateOK Коллекция шаблонов ответов

swagger:response replyCreateOK
*/
type ReplyCreateOK struct {

	/*
	  In: Body
	*/
	Payload *ReplyCreateOKBody `json:"body,omitempty"`
}

// NewReplyCreateOK creates ReplyCreateOK with default headers values
func NewReplyCreateOK() *ReplyCreateOK {

	return &ReplyCreateOK{}
}

// WithPayload adds the payload to the reply create o k response
func (o *ReplyCreateOK) WithPayload(payload *ReplyCreateOKBody) *ReplyCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reply create o k response
func (o *ReplyCreateOK) SetPayload(payload *ReplyCreateOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplyCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplyCreateBadRequestCode is the HTTP code returned for type ReplyCreateBadRequest
const ReplyCreateBadRequestCode int = 400

/*ReplyCreateBadRequest Validation error

swagger:response replyCreateBadRequest
*/
type ReplyCreateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *ReplyCreateBadRequestBody `json:"body,omitempty"`
}

// NewReplyCreateBadRequest creates ReplyCreateBadRequest with default headers values
func NewReplyCreateBadRequest() *ReplyCreateBadRequest {

	return &ReplyCreateBadRequest{}
}

// WithPayload adds the payload to the reply create bad request response
func (o *ReplyCreateBadRequest) WithPayload(payload *ReplyCreateBadRequestBody) *ReplyCreateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reply create bad request response
func (o *ReplyCreateBadRequest) SetPayload(payload *ReplyCreateBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplyCreateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplyCreateMethodNotAllowedCode is the HTTP code returned for type ReplyCreateMethodNotAllowed
const ReplyCreateMethodNotAllowedCode int = 405

/*ReplyCreateMethodNotAllowed Invalid Method

swagger:response replyCreateMethodNotAllowed
*/
type ReplyCreateMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *ReplyCreateMethodNotAllowedBody `json:"body,omitempty"`
}

// NewReplyCreateMethodNotAllowed creates ReplyCreateMethodNotAllowed with default headers values
func NewReplyCreateMethodNotAllowed() *ReplyCreateMethodNotAllowed {

	return &ReplyCreateMethodNotAllowed{}
}

// WithPayload adds the payload to the reply create method not allowed response
func (o *ReplyCreateMethodNotAllowed) WithPayload(payload *ReplyCreateMethodNotAllowedBody) *ReplyCreateMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reply create method not allowed response
func (o *ReplyCreateMethodNotAllowed) SetPayload(payload *ReplyCreateMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplyCreateMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplyCreateInternalServerErrorCode is the HTTP code returned for type ReplyCreateInternalServerError
const ReplyCreateInternalServerErrorCode int = 500

/*ReplyCreateInternalServerError Internal sersver error

swagger:response replyCreateInternalServerError
*/
type ReplyCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *ReplyCreateInternalServerErrorBody `json:"body,omitempty"`
}

// NewReplyCreateInternalServerError creates ReplyCreateInternalServerError with default headers values
func NewReplyCreateInternalServerError() *ReplyCreateInternalServerError {

	return &ReplyCreateInternalServerError{}
}

// WithPayload adds the payload to the reply create internal server error response
func (o *ReplyCreateInternalServerError) WithPayload(payload *ReplyCreateInternalServerErrorBody) *ReplyCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reply create internal server error response
func (o *ReplyCreateInternalServerError) SetPayload(payload *ReplyCreateInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplyCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}