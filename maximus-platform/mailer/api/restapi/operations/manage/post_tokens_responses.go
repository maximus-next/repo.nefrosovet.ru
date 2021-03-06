// Code generated by go-swagger; DO NOT EDIT.

package manage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostTokensOKCode is the HTTP code returned for type PostTokensOK
const PostTokensOKCode int = 200

/*PostTokensOK Коллекция токенов

swagger:response postTokensOK
*/
type PostTokensOK struct {

	/*
	  In: Body
	*/
	Payload *PostTokensOKBody `json:"body,omitempty"`
}

// NewPostTokensOK creates PostTokensOK with default headers values
func NewPostTokensOK() *PostTokensOK {

	return &PostTokensOK{}
}

// WithPayload adds the payload to the post tokens o k response
func (o *PostTokensOK) WithPayload(payload *PostTokensOKBody) *PostTokensOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tokens o k response
func (o *PostTokensOK) SetPayload(payload *PostTokensOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTokensOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTokensBadRequestCode is the HTTP code returned for type PostTokensBadRequest
const PostTokensBadRequestCode int = 400

/*PostTokensBadRequest Validation error

swagger:response postTokensBadRequest
*/
type PostTokensBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostTokensBadRequestBody `json:"body,omitempty"`
}

// NewPostTokensBadRequest creates PostTokensBadRequest with default headers values
func NewPostTokensBadRequest() *PostTokensBadRequest {

	return &PostTokensBadRequest{}
}

// WithPayload adds the payload to the post tokens bad request response
func (o *PostTokensBadRequest) WithPayload(payload *PostTokensBadRequestBody) *PostTokensBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tokens bad request response
func (o *PostTokensBadRequest) SetPayload(payload *PostTokensBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTokensBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTokensForbiddenCode is the HTTP code returned for type PostTokensForbidden
const PostTokensForbiddenCode int = 403

/*PostTokensForbidden Forbidden

swagger:response postTokensForbidden
*/
type PostTokensForbidden struct {

	/*
	  In: Body
	*/
	Payload *PostTokensForbiddenBody `json:"body,omitempty"`
}

// NewPostTokensForbidden creates PostTokensForbidden with default headers values
func NewPostTokensForbidden() *PostTokensForbidden {

	return &PostTokensForbidden{}
}

// WithPayload adds the payload to the post tokens forbidden response
func (o *PostTokensForbidden) WithPayload(payload *PostTokensForbiddenBody) *PostTokensForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tokens forbidden response
func (o *PostTokensForbidden) SetPayload(payload *PostTokensForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTokensForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTokensMethodNotAllowedCode is the HTTP code returned for type PostTokensMethodNotAllowed
const PostTokensMethodNotAllowedCode int = 405

/*PostTokensMethodNotAllowed Invalid Method

swagger:response postTokensMethodNotAllowed
*/
type PostTokensMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PostTokensMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPostTokensMethodNotAllowed creates PostTokensMethodNotAllowed with default headers values
func NewPostTokensMethodNotAllowed() *PostTokensMethodNotAllowed {

	return &PostTokensMethodNotAllowed{}
}

// WithPayload adds the payload to the post tokens method not allowed response
func (o *PostTokensMethodNotAllowed) WithPayload(payload *PostTokensMethodNotAllowedBody) *PostTokensMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tokens method not allowed response
func (o *PostTokensMethodNotAllowed) SetPayload(payload *PostTokensMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTokensMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
