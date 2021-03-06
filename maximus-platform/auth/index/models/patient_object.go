// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatientObject Patient_object
// swagger:model Patient_object
type PatientObject struct {

	// Дата рождения
	BirthDate string `json:"birthDate,omitempty"`

	// Номер документа
	DocumentNumber string `json:"documentNumber,omitempty"`

	// Серия документа
	DocumentSeries string `json:"documentSeries,omitempty"`

	// Тип документа
	DocumentType int64 `json:"documentType,omitempty"`

	// E-mail адрес
	Email string `json:"email,omitempty"`

	// Имя
	FirstName string `json:"firstName,omitempty"`

	// Пол
	// Enum: [MALE FEMALE UNCLASSIFIED]
	Gender string `json:"gender,omitempty"`

	// Номер страхового полиса
	InsurancePolicyNumber string `json:"insurancePolicyNumber,omitempty"`

	// Серия страхового полиса
	InsurancePolicySeries string `json:"insurancePolicySeries,omitempty"`

	// Фамилия
	LastName string `json:"lastName,omitempty"`

	// Номер мобильного телефона
	Mobile string `json:"mobile,omitempty"`

	// Отчество
	Patronymic string `json:"patronymic,omitempty"`

	// Номер смарт-карты
	SmartCardNumber string `json:"smartCardNumber,omitempty"`

	// СНИЛС
	Snils string `json:"snils,omitempty"`

	// Имя пользователя
	Username string `json:"username,omitempty"`
}

// Validate validates this patient object
func (m *PatientObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patientObjectTypeGenderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MALE","FEMALE","UNCLASSIFIED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patientObjectTypeGenderPropEnum = append(patientObjectTypeGenderPropEnum, v)
	}
}

const (

	// PatientObjectGenderMALE captures enum value "MALE"
	PatientObjectGenderMALE string = "MALE"

	// PatientObjectGenderFEMALE captures enum value "FEMALE"
	PatientObjectGenderFEMALE string = "FEMALE"

	// PatientObjectGenderUNCLASSIFIED captures enum value "UNCLASSIFIED"
	PatientObjectGenderUNCLASSIFIED string = "UNCLASSIFIED"
)

// prop value enum
func (m *PatientObject) validateGenderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patientObjectTypeGenderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PatientObject) validateGender(formats strfmt.Registry) error {

	if swag.IsZero(m.Gender) { // not required
		return nil
	}

	// value enum
	if err := m.validateGenderEnum("gender", "body", m.Gender); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatientObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatientObject) UnmarshalBinary(b []byte) error {
	var res PatientObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
