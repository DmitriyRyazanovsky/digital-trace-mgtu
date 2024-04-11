// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestKlimov test klimov
//
// swagger:model testKlimov
type TestKlimov struct {

	// human human
	// Required: true
	HumanHuman *uint64 `json:"human_human"`

	// human nature
	// Required: true
	HumanNature *uint64 `json:"human_nature"`

	// human sign
	// Required: true
	HumanSign *uint64 `json:"human_sign"`

	// human sign system
	// Required: true
	HumanSignSystem *uint64 `json:"human_sign_system"`

	// human technic
	// Required: true
	HumanTechnic *uint64 `json:"human_technic"`
}

// Validate validates this test klimov
func (m *TestKlimov) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHumanHuman(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHumanNature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHumanSign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHumanSignSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHumanTechnic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestKlimov) validateHumanHuman(formats strfmt.Registry) error {

	if err := validate.Required("human_human", "body", m.HumanHuman); err != nil {
		return err
	}

	return nil
}

func (m *TestKlimov) validateHumanNature(formats strfmt.Registry) error {

	if err := validate.Required("human_nature", "body", m.HumanNature); err != nil {
		return err
	}

	return nil
}

func (m *TestKlimov) validateHumanSign(formats strfmt.Registry) error {

	if err := validate.Required("human_sign", "body", m.HumanSign); err != nil {
		return err
	}

	return nil
}

func (m *TestKlimov) validateHumanSignSystem(formats strfmt.Registry) error {

	if err := validate.Required("human_sign_system", "body", m.HumanSignSystem); err != nil {
		return err
	}

	return nil
}

func (m *TestKlimov) validateHumanTechnic(formats strfmt.Registry) error {

	if err := validate.Required("human_technic", "body", m.HumanTechnic); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestKlimov) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestKlimov) UnmarshalBinary(b []byte) error {
	var res TestKlimov
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}