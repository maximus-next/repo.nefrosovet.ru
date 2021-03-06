// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostSendOKCode is the HTTP code returned for type PostSendOK
const PostSendOKCode int = 200

/*PostSendOK Коллекция сообщений

swagger:response postSendOK
*/
type PostSendOK struct {

	/*
	  In: Body
	*/
	Payload *PostSendOKBody `json:"body,omitempty"`
}

// NewPostSendOK creates PostSendOK with default headers values
func NewPostSendOK() *PostSendOK {

	return &PostSendOK{}
}

// WithPayload adds the payload to the post send o k response
func (o *PostSendOK) WithPayload(payload *PostSendOKBody) *PostSendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post send o k response
func (o *PostSendOK) SetPayload(payload *PostSendOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostSendBadRequestCode is the HTTP code returned for type PostSendBadRequest
const PostSendBadRequestCode int = 400

/*PostSendBadRequest Validation error

swagger:response postSendBadRequest
*/
type PostSendBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostSendBadRequestBody `json:"body,omitempty"`
}

// NewPostSendBadRequest creates PostSendBadRequest with default headers values
func NewPostSendBadRequest() *PostSendBadRequest {

	return &PostSendBadRequest{}
}

// WithPayload adds the payload to the post send bad request response
func (o *PostSendBadRequest) WithPayload(payload *PostSendBadRequestBody) *PostSendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post send bad request response
func (o *PostSendBadRequest) SetPayload(payload *PostSendBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostSendForbiddenCode is the HTTP code returned for type PostSendForbidden
const PostSendForbiddenCode int = 403

/*PostSendForbidden Forbidden

swagger:response postSendForbidden
*/
type PostSendForbidden struct {

	/*
	  In: Body
	*/
	Payload *PostSendForbiddenBody `json:"body,omitempty"`
}

// NewPostSendForbidden creates PostSendForbidden with default headers values
func NewPostSendForbidden() *PostSendForbidden {

	return &PostSendForbidden{}
}

// WithPayload adds the payload to the post send forbidden response
func (o *PostSendForbidden) WithPayload(payload *PostSendForbiddenBody) *PostSendForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post send forbidden response
func (o *PostSendForbidden) SetPayload(payload *PostSendForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSendForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostSendMethodNotAllowedCode is the HTTP code returned for type PostSendMethodNotAllowed
const PostSendMethodNotAllowedCode int = 405

/*PostSendMethodNotAllowed Invalid Method

swagger:response postSendMethodNotAllowed
*/
type PostSendMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PostSendMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPostSendMethodNotAllowed creates PostSendMethodNotAllowed with default headers values
func NewPostSendMethodNotAllowed() *PostSendMethodNotAllowed {

	return &PostSendMethodNotAllowed{}
}

// WithPayload adds the payload to the post send method not allowed response
func (o *PostSendMethodNotAllowed) WithPayload(payload *PostSendMethodNotAllowedBody) *PostSendMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post send method not allowed response
func (o *PostSendMethodNotAllowed) SetPayload(payload *PostSendMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSendMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostSendInternalServerErrorCode is the HTTP code returned for type PostSendInternalServerError
const PostSendInternalServerErrorCode int = 500

/*PostSendInternalServerError Internal server error

swagger:response postSendInternalServerError
*/
type PostSendInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostSendInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostSendInternalServerError creates PostSendInternalServerError with default headers values
func NewPostSendInternalServerError() *PostSendInternalServerError {

	return &PostSendInternalServerError{}
}

// WithPayload adds the payload to the post send internal server error response
func (o *PostSendInternalServerError) WithPayload(payload *PostSendInternalServerErrorBody) *PostSendInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post send internal server error response
func (o *PostSendInternalServerError) SetPayload(payload *PostSendInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSendInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
