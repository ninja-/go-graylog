// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AddWidgetRequest add widget request
// swagger:model AddWidgetRequest
type AddWidgetRequest struct {

	// cache time
	CacheTime int64 `json:"cache_time,omitempty"`

	// config
	Config interface{} `json:"config,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this add widget request
func (m *AddWidgetRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AddWidgetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddWidgetRequest) UnmarshalBinary(b []byte) error {
	var res AddWidgetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
