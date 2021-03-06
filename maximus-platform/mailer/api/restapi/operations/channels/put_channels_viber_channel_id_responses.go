// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutChannelsViberChannelIDOKCode is the HTTP code returned for type PutChannelsViberChannelIDOK
const PutChannelsViberChannelIDOKCode int = 200

/*PutChannelsViberChannelIDOK Коллекция каналов

swagger:response putChannelsViberChannelIdOK
*/
type PutChannelsViberChannelIDOK struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsViberChannelIDOKBody `json:"body,omitempty"`
}

// NewPutChannelsViberChannelIDOK creates PutChannelsViberChannelIDOK with default headers values
func NewPutChannelsViberChannelIDOK() *PutChannelsViberChannelIDOK {

	return &PutChannelsViberChannelIDOK{}
}

// WithPayload adds the payload to the put channels viber channel Id o k response
func (o *PutChannelsViberChannelIDOK) WithPayload(payload *PutChannelsViberChannelIDOKBody) *PutChannelsViberChannelIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels viber channel Id o k response
func (o *PutChannelsViberChannelIDOK) SetPayload(payload *PutChannelsViberChannelIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsViberChannelIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsViberChannelIDBadRequestCode is the HTTP code returned for type PutChannelsViberChannelIDBadRequest
const PutChannelsViberChannelIDBadRequestCode int = 400

/*PutChannelsViberChannelIDBadRequest Validation error

swagger:response putChannelsViberChannelIdBadRequest
*/
type PutChannelsViberChannelIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsViberChannelIDBadRequestBody `json:"body,omitempty"`
}

// NewPutChannelsViberChannelIDBadRequest creates PutChannelsViberChannelIDBadRequest with default headers values
func NewPutChannelsViberChannelIDBadRequest() *PutChannelsViberChannelIDBadRequest {

	return &PutChannelsViberChannelIDBadRequest{}
}

// WithPayload adds the payload to the put channels viber channel Id bad request response
func (o *PutChannelsViberChannelIDBadRequest) WithPayload(payload *PutChannelsViberChannelIDBadRequestBody) *PutChannelsViberChannelIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels viber channel Id bad request response
func (o *PutChannelsViberChannelIDBadRequest) SetPayload(payload *PutChannelsViberChannelIDBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsViberChannelIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsViberChannelIDForbiddenCode is the HTTP code returned for type PutChannelsViberChannelIDForbidden
const PutChannelsViberChannelIDForbiddenCode int = 403

/*PutChannelsViberChannelIDForbidden Forbidden

swagger:response putChannelsViberChannelIdForbidden
*/
type PutChannelsViberChannelIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsViberChannelIDForbiddenBody `json:"body,omitempty"`
}

// NewPutChannelsViberChannelIDForbidden creates PutChannelsViberChannelIDForbidden with default headers values
func NewPutChannelsViberChannelIDForbidden() *PutChannelsViberChannelIDForbidden {

	return &PutChannelsViberChannelIDForbidden{}
}

// WithPayload adds the payload to the put channels viber channel Id forbidden response
func (o *PutChannelsViberChannelIDForbidden) WithPayload(payload *PutChannelsViberChannelIDForbiddenBody) *PutChannelsViberChannelIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels viber channel Id forbidden response
func (o *PutChannelsViberChannelIDForbidden) SetPayload(payload *PutChannelsViberChannelIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsViberChannelIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsViberChannelIDNotFoundCode is the HTTP code returned for type PutChannelsViberChannelIDNotFound
const PutChannelsViberChannelIDNotFoundCode int = 404

/*PutChannelsViberChannelIDNotFound Not found

swagger:response putChannelsViberChannelIdNotFound
*/
type PutChannelsViberChannelIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsViberChannelIDNotFoundBody `json:"body,omitempty"`
}

// NewPutChannelsViberChannelIDNotFound creates PutChannelsViberChannelIDNotFound with default headers values
func NewPutChannelsViberChannelIDNotFound() *PutChannelsViberChannelIDNotFound {

	return &PutChannelsViberChannelIDNotFound{}
}

// WithPayload adds the payload to the put channels viber channel Id not found response
func (o *PutChannelsViberChannelIDNotFound) WithPayload(payload *PutChannelsViberChannelIDNotFoundBody) *PutChannelsViberChannelIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels viber channel Id not found response
func (o *PutChannelsViberChannelIDNotFound) SetPayload(payload *PutChannelsViberChannelIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsViberChannelIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsViberChannelIDMethodNotAllowedCode is the HTTP code returned for type PutChannelsViberChannelIDMethodNotAllowed
const PutChannelsViberChannelIDMethodNotAllowedCode int = 405

/*PutChannelsViberChannelIDMethodNotAllowed Invalid Method

swagger:response putChannelsViberChannelIdMethodNotAllowed
*/
type PutChannelsViberChannelIDMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsViberChannelIDMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPutChannelsViberChannelIDMethodNotAllowed creates PutChannelsViberChannelIDMethodNotAllowed with default headers values
func NewPutChannelsViberChannelIDMethodNotAllowed() *PutChannelsViberChannelIDMethodNotAllowed {

	return &PutChannelsViberChannelIDMethodNotAllowed{}
}

// WithPayload adds the payload to the put channels viber channel Id method not allowed response
func (o *PutChannelsViberChannelIDMethodNotAllowed) WithPayload(payload *PutChannelsViberChannelIDMethodNotAllowedBody) *PutChannelsViberChannelIDMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels viber channel Id method not allowed response
func (o *PutChannelsViberChannelIDMethodNotAllowed) SetPayload(payload *PutChannelsViberChannelIDMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsViberChannelIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
