// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SingleStreamRuleSummaryResponse single stream rule summary response
// swagger:model SingleStreamRuleSummaryResponse
type SingleStreamRuleSummaryResponse struct {

	// streamrule id
	StreamruleID string `json:"streamrule_id,omitempty"`
}

// Validate validates this single stream rule summary response
func (m *SingleStreamRuleSummaryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SingleStreamRuleSummaryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SingleStreamRuleSummaryResponse) UnmarshalBinary(b []byte) error {
	var res SingleStreamRuleSummaryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
