// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "repo.nefrosovet.ru/maximus-platform/auth/index/models"
)

// PutClinicsClinicGUIDServicesServiceGUIDReader is a Reader for the PutClinicsClinicGUIDServicesServiceGUID structure.
type PutClinicsClinicGUIDServicesServiceGUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClinicsClinicGUIDServicesServiceGUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutClinicsClinicGUIDServicesServiceGUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutClinicsClinicGUIDServicesServiceGUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutClinicsClinicGUIDServicesServiceGUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutClinicsClinicGUIDServicesServiceGUIDOK creates a PutClinicsClinicGUIDServicesServiceGUIDOK with default headers values
func NewPutClinicsClinicGUIDServicesServiceGUIDOK() *PutClinicsClinicGUIDServicesServiceGUIDOK {
	return &PutClinicsClinicGUIDServicesServiceGUIDOK{}
}

/*PutClinicsClinicGUIDServicesServiceGUIDOK handles this case with default header values.

Коллекция сервисов
*/
type PutClinicsClinicGUIDServicesServiceGUIDOK struct {
	Payload *PutClinicsClinicGUIDServicesServiceGUIDOKBody
}

func (o *PutClinicsClinicGUIDServicesServiceGUIDOK) Error() string {
	return fmt.Sprintf("[PUT /clinics/{clinicGUID}/services/{serviceGUID}][%d] putClinicsClinicGuidServicesServiceGuidOK  %+v", 200, o.Payload)
}

func (o *PutClinicsClinicGUIDServicesServiceGUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutClinicsClinicGUIDServicesServiceGUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClinicsClinicGUIDServicesServiceGUIDBadRequest creates a PutClinicsClinicGUIDServicesServiceGUIDBadRequest with default headers values
func NewPutClinicsClinicGUIDServicesServiceGUIDBadRequest() *PutClinicsClinicGUIDServicesServiceGUIDBadRequest {
	return &PutClinicsClinicGUIDServicesServiceGUIDBadRequest{}
}

/*PutClinicsClinicGUIDServicesServiceGUIDBadRequest handles this case with default header values.

Validation error
*/
type PutClinicsClinicGUIDServicesServiceGUIDBadRequest struct {
	Payload *PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody
}

func (o *PutClinicsClinicGUIDServicesServiceGUIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /clinics/{clinicGUID}/services/{serviceGUID}][%d] putClinicsClinicGuidServicesServiceGuidBadRequest  %+v", 400, o.Payload)
}

func (o *PutClinicsClinicGUIDServicesServiceGUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClinicsClinicGUIDServicesServiceGUIDNotFound creates a PutClinicsClinicGUIDServicesServiceGUIDNotFound with default headers values
func NewPutClinicsClinicGUIDServicesServiceGUIDNotFound() *PutClinicsClinicGUIDServicesServiceGUIDNotFound {
	return &PutClinicsClinicGUIDServicesServiceGUIDNotFound{}
}

/*PutClinicsClinicGUIDServicesServiceGUIDNotFound handles this case with default header values.

Not found
*/
type PutClinicsClinicGUIDServicesServiceGUIDNotFound struct {
	Payload *PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody
}

func (o *PutClinicsClinicGUIDServicesServiceGUIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /clinics/{clinicGUID}/services/{serviceGUID}][%d] putClinicsClinicGuidServicesServiceGuidNotFound  %+v", 404, o.Payload)
}

func (o *PutClinicsClinicGUIDServicesServiceGUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowed creates a PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowed with default headers values
func NewPutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowed() *PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowed {
	return &PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowed{}
}

/*PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowed handles this case with default header values.

Invalid Method
*/
type PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowed struct {
	Payload *PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody
}

func (o *PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /clinics/{clinicGUID}/services/{serviceGUID}][%d] putClinicsClinicGuidServicesServiceGuidMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody put clinics clinic GUID services service GUID bad request body
swagger:model PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody
*/
type PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody struct {
	models.Error400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PutClinicsClinicGUIDServicesServiceGUIDBadRequestBodyAO0
	var putClinicsClinicGUIDServicesServiceGUIDBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &putClinicsClinicGUIDServicesServiceGUIDBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = putClinicsClinicGUIDServicesServiceGUIDBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	putClinicsClinicGUIDServicesServiceGUIDBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putClinicsClinicGUIDServicesServiceGUIDBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put clinics clinic GUID services service GUID bad request body
func (o *PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error400Data
	if err := o.Error400Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutClinicsClinicGUIDServicesServiceGUIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutClinicsClinicGUIDServicesServiceGUIDBody put clinics clinic GUID services service GUID body
swagger:model PutClinicsClinicGUIDServicesServiceGUIDBody
*/
type PutClinicsClinicGUIDServicesServiceGUIDBody struct {
	models.ServiceObject

	PutClinicsClinicGUIDServicesServiceGUIDParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutClinicsClinicGUIDServicesServiceGUIDBody) UnmarshalJSON(raw []byte) error {
	// PutClinicsClinicGUIDServicesServiceGUIDParamsBodyAO0
	var putClinicsClinicGUIDServicesServiceGUIDParamsBodyAO0 models.ServiceObject
	if err := swag.ReadJSON(raw, &putClinicsClinicGUIDServicesServiceGUIDParamsBodyAO0); err != nil {
		return err
	}
	o.ServiceObject = putClinicsClinicGUIDServicesServiceGUIDParamsBodyAO0

	// PutClinicsClinicGUIDServicesServiceGUIDParamsBodyAO1
	var putClinicsClinicGUIDServicesServiceGUIDParamsBodyAO1 PutClinicsClinicGUIDServicesServiceGUIDParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &putClinicsClinicGUIDServicesServiceGUIDParamsBodyAO1); err != nil {
		return err
	}
	o.PutClinicsClinicGUIDServicesServiceGUIDParamsBodyAllOf1 = putClinicsClinicGUIDServicesServiceGUIDParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutClinicsClinicGUIDServicesServiceGUIDBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putClinicsClinicGUIDServicesServiceGUIDParamsBodyAO0, err := swag.WriteJSON(o.ServiceObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putClinicsClinicGUIDServicesServiceGUIDParamsBodyAO0)

	putClinicsClinicGUIDServicesServiceGUIDParamsBodyAO1, err := swag.WriteJSON(o.PutClinicsClinicGUIDServicesServiceGUIDParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putClinicsClinicGUIDServicesServiceGUIDParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put clinics clinic GUID services service GUID body
func (o *PutClinicsClinicGUIDServicesServiceGUIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ServiceObject
	if err := o.ServiceObject.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PutClinicsClinicGUIDServicesServiceGUIDParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PutClinicsClinicGUIDServicesServiceGUIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutClinicsClinicGUIDServicesServiceGUIDBody) UnmarshalBinary(b []byte) error {
	var res PutClinicsClinicGUIDServicesServiceGUIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody put clinics clinic GUID services service GUID method not allowed body
swagger:model PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody
*/
type PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBodyAO0
	var putClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &putClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = putClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	putClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put clinics clinic GUID services service GUID method not allowed body
func (o *PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error405Data
	if err := o.Error405Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PutClinicsClinicGUIDServicesServiceGUIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody put clinics clinic GUID services service GUID not found body
swagger:model PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody
*/
type PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// PutClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0
	var putClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &putClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = putClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	putClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put clinics clinic GUID services service GUID not found body
func (o *PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error404Data
	if err := o.Error404Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutClinicsClinicGUIDServicesServiceGUIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutClinicsClinicGUIDServicesServiceGUIDOKBody put clinics clinic GUID services service GUID o k body
swagger:model PutClinicsClinicGUIDServicesServiceGUIDOKBody
*/
type PutClinicsClinicGUIDServicesServiceGUIDOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutClinicsClinicGUIDServicesServiceGUIDOKBody) UnmarshalJSON(raw []byte) error {
	// PutClinicsClinicGUIDServicesServiceGUIDOKBodyAO0
	var putClinicsClinicGUIDServicesServiceGUIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &putClinicsClinicGUIDServicesServiceGUIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = putClinicsClinicGUIDServicesServiceGUIDOKBodyAO0

	// PutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1
	var dataPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutClinicsClinicGUIDServicesServiceGUIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putClinicsClinicGUIDServicesServiceGUIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putClinicsClinicGUIDServicesServiceGUIDOKBodyAO0)

	var dataPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1.Data = o.Data

	jsonDataPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1, errPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1 := swag.WriteJSON(dataPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1)
	if errPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1 != nil {
		return nil, errPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutClinicsClinicGUIDServicesServiceGUIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put clinics clinic GUID services service GUID o k body
func (o *PutClinicsClinicGUIDServicesServiceGUIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SuccessData
	if err := o.SuccessData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutClinicsClinicGUIDServicesServiceGUIDOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("putClinicsClinicGuidServicesServiceGuidOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutClinicsClinicGUIDServicesServiceGUIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutClinicsClinicGUIDServicesServiceGUIDOKBody) UnmarshalBinary(b []byte) error {
	var res PutClinicsClinicGUIDServicesServiceGUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutClinicsClinicGUIDServicesServiceGUIDParamsBodyAllOf1 put clinics clinic GUID services service GUID params body all of1
swagger:model PutClinicsClinicGUIDServicesServiceGUIDParamsBodyAllOf1
*/
type PutClinicsClinicGUIDServicesServiceGUIDParamsBodyAllOf1 interface{}
