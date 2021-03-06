// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChannelParamsViber Channel_params_viber
// swagger:model Channel_params_viber
type ChannelParamsViber struct {

	// Альтернативный текст приветствия
	// Required: true
	AlternateText *string `json:"alternateText"`

	// Текст ответа
	// Required: true
	AnswerText *string `json:"answerText"`

	// URL аватара бота
	// Required: true
	BotAvatar *string `json:"botAvatar"`

	// Название бота в Viber
	// Required: true
	BotName *string `json:"botName"`

	// Текст на кнопке
	// Required: true
	ButtonText *string `json:"buttonText"`

	// Текст приветствия
	// Required: true
	GreetingText *string `json:"greetingText"`

	// Токен доступа
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this channel params viber
func (m *ChannelParamsViber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlternateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnswerText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBotAvatar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBotName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateButtonText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGreetingText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelParamsViber) validateAlternateText(formats strfmt.Registry) error {

	if err := validate.Required("alternateText", "body", m.AlternateText); err != nil {
		return err
	}

	return nil
}

func (m *ChannelParamsViber) validateAnswerText(formats strfmt.Registry) error {

	if err := validate.Required("answerText", "body", m.AnswerText); err != nil {
		return err
	}

	return nil
}

func (m *ChannelParamsViber) validateBotAvatar(formats strfmt.Registry) error {

	if err := validate.Required("botAvatar", "body", m.BotAvatar); err != nil {
		return err
	}

	return nil
}

func (m *ChannelParamsViber) validateBotName(formats strfmt.Registry) error {

	if err := validate.Required("botName", "body", m.BotName); err != nil {
		return err
	}

	return nil
}

func (m *ChannelParamsViber) validateButtonText(formats strfmt.Registry) error {

	if err := validate.Required("buttonText", "body", m.ButtonText); err != nil {
		return err
	}

	return nil
}

func (m *ChannelParamsViber) validateGreetingText(formats strfmt.Registry) error {

	if err := validate.Required("greetingText", "body", m.GreetingText); err != nil {
		return err
	}

	return nil
}

func (m *ChannelParamsViber) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelParamsViber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelParamsViber) UnmarshalBinary(b []byte) error {
	var res ChannelParamsViber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
