// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ChangeUserRequest change user request
// swagger:model ChangeUserRequest
type ChangeUserRequest struct {

	// email
	Email string `json:"email,omitempty"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// permissions
	Permissions []string `json:"permissions"`

	// roles
	Roles []string `json:"roles"`

	// session timeout ms
	SessionTimeoutMs int64 `json:"session_timeout_ms,omitempty"`

	// startpage
	Startpage *ChangeUserRequestStartpage `json:"startpage,omitempty"`

	// timezone
	Timezone string `json:"timezone,omitempty"`
}

// Validate validates this change user request
func (m *ChangeUserRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartpage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangeUserRequest) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	return nil
}

func (m *ChangeUserRequest) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	return nil
}

func (m *ChangeUserRequest) validateStartpage(formats strfmt.Registry) error {

	if swag.IsZero(m.Startpage) { // not required
		return nil
	}

	if m.Startpage != nil {

		if err := m.Startpage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startpage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangeUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeUserRequest) UnmarshalBinary(b []byte) error {
	var res ChangeUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
