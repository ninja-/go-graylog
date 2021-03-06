// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InputTypeInfo input type info
// swagger:model InputTypeInfo
type InputTypeInfo struct {

	// is exclusive
	IsExclusive bool `json:"is_exclusive,omitempty"`

	// link to docs
	LinkToDocs string `json:"link_to_docs,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// requested configuration
	RequestedConfiguration interface{} `json:"requested_configuration,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this input type info
func (m *InputTypeInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InputTypeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InputTypeInfo) UnmarshalBinary(b []byte) error {
	var res InputTypeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
