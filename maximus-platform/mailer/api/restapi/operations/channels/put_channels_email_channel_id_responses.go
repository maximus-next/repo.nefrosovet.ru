// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutChannelsEmailChannelIDOKCode is the HTTP code returned for type PutChannelsEmailChannelIDOK
const PutChannelsEmailChannelIDOKCode int = 200

/*PutChannelsEmailChannelIDOK Коллекция каналов

swagger:response putChannelsEmailChannelIdOK
*/
type PutChannelsEmailChannelIDOK struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsEmailChannelIDOKBody `json:"body,omitempty"`
}

// NewPutChannelsEmailChannelIDOK creates PutChannelsEmailChannelIDOK with default headers values
func NewPutChannelsEmailChannelIDOK() *PutChannelsEmailChannelIDOK {

	return &PutChannelsEmailChannelIDOK{}
}

// WithPayload adds the payload to the put channels email channel Id o k response
func (o *PutChannelsEmailChannelIDOK) WithPayload(payload *PutChannelsEmailChannelIDOKBody) *PutChannelsEmailChannelIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels email channel Id o k response
func (o *PutChannelsEmailChannelIDOK) SetPayload(payload *PutChannelsEmailChannelIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsEmailChannelIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsEmailChannelIDBadRequestCode is the HTTP code returned for type PutChannelsEmailChannelIDBadRequest
const PutChannelsEmailChannelIDBadRequestCode int = 400

/*PutChannelsEmailChannelIDBadRequest Validation error

swagger:response putChannelsEmailChannelIdBadRequest
*/
type PutChannelsEmailChannelIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsEmailChannelIDBadRequestBody `json:"body,omitempty"`
}

// NewPutChannelsEmailChannelIDBadRequest creates PutChannelsEmailChannelIDBadRequest with default headers values
func NewPutChannelsEmailChannelIDBadRequest() *PutChannelsEmailChannelIDBadRequest {

	return &PutChannelsEmailChannelIDBadRequest{}
}

// WithPayload adds the payload to the put channels email channel Id bad request response
func (o *PutChannelsEmailChannelIDBadRequest) WithPayload(payload *PutChannelsEmailChannelIDBadRequestBody) *PutChannelsEmailChannelIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels email channel Id bad request response
func (o *PutChannelsEmailChannelIDBadRequest) SetPayload(payload *PutChannelsEmailChannelIDBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsEmailChannelIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsEmailChannelIDForbiddenCode is the HTTP code returned for type PutChannelsEmailChannelIDForbidden
const PutChannelsEmailChannelIDForbiddenCode int = 403

/*PutChannelsEmailChannelIDForbidden Forbidden

swagger:response putChannelsEmailChannelIdForbidden
*/
type PutChannelsEmailChannelIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsEmailChannelIDForbiddenBody `json:"body,omitempty"`
}

// NewPutChannelsEmailChannelIDForbidden creates PutChannelsEmailChannelIDForbidden with default headers values
func NewPutChannelsEmailChannelIDForbidden() *PutChannelsEmailChannelIDForbidden {

	return &PutChannelsEmailChannelIDForbidden{}
}

// WithPayload adds the payload to the put channels email channel Id forbidden response
func (o *PutChannelsEmailChannelIDForbidden) WithPayload(payload *PutChannelsEmailChannelIDForbiddenBody) *PutChannelsEmailChannelIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels email channel Id forbidden response
func (o *PutChannelsEmailChannelIDForbidden) SetPayload(payload *PutChannelsEmailChannelIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsEmailChannelIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsEmailChannelIDNotFoundCode is the HTTP code returned for type PutChannelsEmailChannelIDNotFound
const PutChannelsEmailChannelIDNotFoundCode int = 404

/*PutChannelsEmailChannelIDNotFound Not found

swagger:response putChannelsEmailChannelIdNotFound
*/
type PutChannelsEmailChannelIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsEmailChannelIDNotFoundBody `json:"body,omitempty"`
}

// NewPutChannelsEmailChannelIDNotFound creates PutChannelsEmailChannelIDNotFound with default headers values
func NewPutChannelsEmailChannelIDNotFound() *PutChannelsEmailChannelIDNotFound {

	return &PutChannelsEmailChannelIDNotFound{}
}

// WithPayload adds the payload to the put channels email channel Id not found response
func (o *PutChannelsEmailChannelIDNotFound) WithPayload(payload *PutChannelsEmailChannelIDNotFoundBody) *PutChannelsEmailChannelIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels email channel Id not found response
func (o *PutChannelsEmailChannelIDNotFound) SetPayload(payload *PutChannelsEmailChannelIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsEmailChannelIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsEmailChannelIDMethodNotAllowedCode is the HTTP code returned for type PutChannelsEmailChannelIDMethodNotAllowed
const PutChannelsEmailChannelIDMethodNotAllowedCode int = 405

/*PutChannelsEmailChannelIDMethodNotAllowed Invalid Method

swagger:response putChannelsEmailChannelIdMethodNotAllowed
*/
type PutChannelsEmailChannelIDMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsEmailChannelIDMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPutChannelsEmailChannelIDMethodNotAllowed creates PutChannelsEmailChannelIDMethodNotAllowed with default headers values
func NewPutChannelsEmailChannelIDMethodNotAllowed() *PutChannelsEmailChannelIDMethodNotAllowed {

	return &PutChannelsEmailChannelIDMethodNotAllowed{}
}

// WithPayload adds the payload to the put channels email channel Id method not allowed response
func (o *PutChannelsEmailChannelIDMethodNotAllowed) WithPayload(payload *PutChannelsEmailChannelIDMethodNotAllowedBody) *PutChannelsEmailChannelIDMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels email channel Id method not allowed response
func (o *PutChannelsEmailChannelIDMethodNotAllowed) SetPayload(payload *PutChannelsEmailChannelIDMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsEmailChannelIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}