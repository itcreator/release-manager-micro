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

// SemverTagSet Semver set of tags
// swagger:model SemverTagSet
type SemverTagSet struct {

	// Array of all tags (e.g: [latest, v1.2.1, v1.2, v1])
	All []string `json:"all"`

	// The version tag which is generated for custom branch  (e.g: v1.2.0-rc for release branch)
	// Max Length: 150
	// Min Length: 3
	Branch string `json:"branch,omitempty"`

	// The full version tag (e.g: v1.2.1 or v1.2.0-rc.1 or v1.2.0-feature-22.1)
	// Required: true
	// Max Length: 150
	// Min Length: 3
	Full string `json:"full"`

	// True if this version is latest. False - if not.
	// Required: true
	IsLatest bool `json:"isLatest"`

	// The minor version tag (e.g: v1)
	// Max Length: 150
	// Min Length: 3
	Major string `json:"major,omitempty"`

	// The minor version tag (e.g: v1.2)
	// Max Length: 150
	// Min Length: 3
	Minor string `json:"minor,omitempty"`
}

// Validate validates this semver tag set
func (m *SemverTagSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAll(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBranch(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFull(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIsLatest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMajor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SemverTagSet) validateAll(formats strfmt.Registry) error {

	if swag.IsZero(m.All) { // not required
		return nil
	}

	return nil
}

func (m *SemverTagSet) validateBranch(formats strfmt.Registry) error {

	if swag.IsZero(m.Branch) { // not required
		return nil
	}

	if err := validate.MinLength("branch", "body", string(m.Branch), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("branch", "body", string(m.Branch), 150); err != nil {
		return err
	}

	return nil
}

func (m *SemverTagSet) validateFull(formats strfmt.Registry) error {

	if err := validate.RequiredString("full", "body", string(m.Full)); err != nil {
		return err
	}

	if err := validate.MinLength("full", "body", string(m.Full), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("full", "body", string(m.Full), 150); err != nil {
		return err
	}

	return nil
}

func (m *SemverTagSet) validateIsLatest(formats strfmt.Registry) error {

	if err := validate.Required("isLatest", "body", bool(m.IsLatest)); err != nil {
		return err
	}

	return nil
}

func (m *SemverTagSet) validateMajor(formats strfmt.Registry) error {

	if swag.IsZero(m.Major) { // not required
		return nil
	}

	if err := validate.MinLength("major", "body", string(m.Major), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("major", "body", string(m.Major), 150); err != nil {
		return err
	}

	return nil
}

func (m *SemverTagSet) validateMinor(formats strfmt.Registry) error {

	if swag.IsZero(m.Minor) { // not required
		return nil
	}

	if err := validate.MinLength("minor", "body", string(m.Minor), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("minor", "body", string(m.Minor), 150); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SemverTagSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SemverTagSet) UnmarshalBinary(b []byte) error {
	var res SemverTagSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
