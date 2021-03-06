// Code generated by go-swagger; DO NOT EDIT.

package treatment_episodes

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

// PostTreatmentEpisodesReader is a Reader for the PostTreatmentEpisodes structure.
type PostTreatmentEpisodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTreatmentEpisodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostTreatmentEpisodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostTreatmentEpisodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostTreatmentEpisodesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTreatmentEpisodesOK creates a PostTreatmentEpisodesOK with default headers values
func NewPostTreatmentEpisodesOK() *PostTreatmentEpisodesOK {
	return &PostTreatmentEpisodesOK{}
}

/*PostTreatmentEpisodesOK handles this case with default header values.

Коллекция эпизодов лечения
*/
type PostTreatmentEpisodesOK struct {
	Payload *PostTreatmentEpisodesOKBody
}

func (o *PostTreatmentEpisodesOK) Error() string {
	return fmt.Sprintf("[POST /treatmentEpisodes][%d] postTreatmentEpisodesOK  %+v", 200, o.Payload)
}

func (o *PostTreatmentEpisodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostTreatmentEpisodesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTreatmentEpisodesBadRequest creates a PostTreatmentEpisodesBadRequest with default headers values
func NewPostTreatmentEpisodesBadRequest() *PostTreatmentEpisodesBadRequest {
	return &PostTreatmentEpisodesBadRequest{}
}

/*PostTreatmentEpisodesBadRequest handles this case with default header values.

Validation error
*/
type PostTreatmentEpisodesBadRequest struct {
	Payload *PostTreatmentEpisodesBadRequestBody
}

func (o *PostTreatmentEpisodesBadRequest) Error() string {
	return fmt.Sprintf("[POST /treatmentEpisodes][%d] postTreatmentEpisodesBadRequest  %+v", 400, o.Payload)
}

func (o *PostTreatmentEpisodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostTreatmentEpisodesBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTreatmentEpisodesMethodNotAllowed creates a PostTreatmentEpisodesMethodNotAllowed with default headers values
func NewPostTreatmentEpisodesMethodNotAllowed() *PostTreatmentEpisodesMethodNotAllowed {
	return &PostTreatmentEpisodesMethodNotAllowed{}
}

/*PostTreatmentEpisodesMethodNotAllowed handles this case with default header values.

Invalid Method
*/
type PostTreatmentEpisodesMethodNotAllowed struct {
	Payload *PostTreatmentEpisodesMethodNotAllowedBody
}

func (o *PostTreatmentEpisodesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /treatmentEpisodes][%d] postTreatmentEpisodesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostTreatmentEpisodesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostTreatmentEpisodesMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostTreatmentEpisodesBadRequestBody post treatment episodes bad request body
swagger:model PostTreatmentEpisodesBadRequestBody
*/
type PostTreatmentEpisodesBadRequestBody struct {
	models.Error400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostTreatmentEpisodesBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostTreatmentEpisodesBadRequestBodyAO0
	var postTreatmentEpisodesBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &postTreatmentEpisodesBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = postTreatmentEpisodesBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostTreatmentEpisodesBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postTreatmentEpisodesBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postTreatmentEpisodesBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post treatment episodes bad request body
func (o *PostTreatmentEpisodesBadRequestBody) Validate(formats strfmt.Registry) error {
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
func (o *PostTreatmentEpisodesBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTreatmentEpisodesBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostTreatmentEpisodesBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTreatmentEpisodesBody post treatment episodes body
swagger:model PostTreatmentEpisodesBody
*/
type PostTreatmentEpisodesBody struct {
	models.MainData

	models.TreatmentEpisodeObject

	PostTreatmentEpisodesParamsBodyAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostTreatmentEpisodesBody) UnmarshalJSON(raw []byte) error {
	// PostTreatmentEpisodesParamsBodyAO0
	var postTreatmentEpisodesParamsBodyAO0 models.MainData
	if err := swag.ReadJSON(raw, &postTreatmentEpisodesParamsBodyAO0); err != nil {
		return err
	}
	o.MainData = postTreatmentEpisodesParamsBodyAO0

	// PostTreatmentEpisodesParamsBodyAO1
	var postTreatmentEpisodesParamsBodyAO1 models.TreatmentEpisodeObject
	if err := swag.ReadJSON(raw, &postTreatmentEpisodesParamsBodyAO1); err != nil {
		return err
	}
	o.TreatmentEpisodeObject = postTreatmentEpisodesParamsBodyAO1

	// PostTreatmentEpisodesParamsBodyAO2
	var postTreatmentEpisodesParamsBodyAO2 PostTreatmentEpisodesParamsBodyAllOf2
	if err := swag.ReadJSON(raw, &postTreatmentEpisodesParamsBodyAO2); err != nil {
		return err
	}
	o.PostTreatmentEpisodesParamsBodyAllOf2 = postTreatmentEpisodesParamsBodyAO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostTreatmentEpisodesBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	postTreatmentEpisodesParamsBodyAO0, err := swag.WriteJSON(o.MainData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postTreatmentEpisodesParamsBodyAO0)

	postTreatmentEpisodesParamsBodyAO1, err := swag.WriteJSON(o.TreatmentEpisodeObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postTreatmentEpisodesParamsBodyAO1)

	postTreatmentEpisodesParamsBodyAO2, err := swag.WriteJSON(o.PostTreatmentEpisodesParamsBodyAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postTreatmentEpisodesParamsBodyAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post treatment episodes body
func (o *PostTreatmentEpisodesBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.MainData
	if err := o.MainData.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.TreatmentEpisodeObject
	if err := o.TreatmentEpisodeObject.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PostTreatmentEpisodesParamsBodyAllOf2

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostTreatmentEpisodesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTreatmentEpisodesBody) UnmarshalBinary(b []byte) error {
	var res PostTreatmentEpisodesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTreatmentEpisodesMethodNotAllowedBody post treatment episodes method not allowed body
swagger:model PostTreatmentEpisodesMethodNotAllowedBody
*/
type PostTreatmentEpisodesMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostTreatmentEpisodesMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostTreatmentEpisodesMethodNotAllowedBodyAO0
	var postTreatmentEpisodesMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postTreatmentEpisodesMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postTreatmentEpisodesMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostTreatmentEpisodesMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postTreatmentEpisodesMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postTreatmentEpisodesMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post treatment episodes method not allowed body
func (o *PostTreatmentEpisodesMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostTreatmentEpisodesMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTreatmentEpisodesMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostTreatmentEpisodesMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTreatmentEpisodesOKBody post treatment episodes o k body
swagger:model PostTreatmentEpisodesOKBody
*/
type PostTreatmentEpisodesOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostTreatmentEpisodesOKBody) UnmarshalJSON(raw []byte) error {
	// PostTreatmentEpisodesOKBodyAO0
	var postTreatmentEpisodesOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postTreatmentEpisodesOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postTreatmentEpisodesOKBodyAO0

	// PostTreatmentEpisodesOKBodyAO1
	var dataPostTreatmentEpisodesOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostTreatmentEpisodesOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostTreatmentEpisodesOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostTreatmentEpisodesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postTreatmentEpisodesOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postTreatmentEpisodesOKBodyAO0)

	var dataPostTreatmentEpisodesOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataPostTreatmentEpisodesOKBodyAO1.Data = o.Data

	jsonDataPostTreatmentEpisodesOKBodyAO1, errPostTreatmentEpisodesOKBodyAO1 := swag.WriteJSON(dataPostTreatmentEpisodesOKBodyAO1)
	if errPostTreatmentEpisodesOKBodyAO1 != nil {
		return nil, errPostTreatmentEpisodesOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostTreatmentEpisodesOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post treatment episodes o k body
func (o *PostTreatmentEpisodesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostTreatmentEpisodesOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("postTreatmentEpisodesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTreatmentEpisodesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTreatmentEpisodesOKBody) UnmarshalBinary(b []byte) error {
	var res PostTreatmentEpisodesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTreatmentEpisodesParamsBodyAllOf2 post treatment episodes params body all of2
swagger:model PostTreatmentEpisodesParamsBodyAllOf2
*/
type PostTreatmentEpisodesParamsBodyAllOf2 interface{}
