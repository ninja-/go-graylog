// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateUserRequest create user request
// swagger:model CreateUserRequest
type CreateUserRequest struct {

	// email
	Email string `json:"email,omitempty"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// permissions
	Permissions []string `json:"permissions"`

	// roles
	Roles []string `json:"roles"`

	// session timeout ms
	SessionTimeoutMs int64 `json:"session_timeout_ms,omitempty"`

	// startpage
	Startpage *CreateUserRequestStartpage `json:"startpage,omitempty"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this create user request
func (m *CreateUserRequest) Validate(formats strfmt.Registry) error {
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

func (m *CreateUserRequest) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	return nil
}

func (m *CreateUserRequest) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	return nil
}

func (m *CreateUserRequest) validateStartpage(formats strfmt.Registry) error {

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
func (m *CreateUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUserRequest) UnmarshalBinary(b []byte) error {
	var res CreateUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
