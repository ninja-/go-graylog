// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateSavedSearchRequest create saved search request
// swagger:model CreateSavedSearchRequest
type CreateSavedSearchRequest struct {

	// query
	Query interface{} `json:"query,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this create saved search request
func (m *CreateSavedSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CreateSavedSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSavedSearchRequest) UnmarshalBinary(b []byte) error {
	var res CreateSavedSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
