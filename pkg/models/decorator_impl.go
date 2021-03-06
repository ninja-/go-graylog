// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DecoratorImpl decorator impl
// swagger:model DecoratorImpl
type DecoratorImpl struct {

	// config
	Config interface{} `json:"config,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// order
	Order int64 `json:"order,omitempty"`

	// stream
	Stream string `json:"stream,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this decorator impl
func (m *DecoratorImpl) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DecoratorImpl) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DecoratorImpl) UnmarshalBinary(b []byte) error {
	var res DecoratorImpl
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
