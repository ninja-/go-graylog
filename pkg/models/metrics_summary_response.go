// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MetricsSummaryResponse metrics summary response
// swagger:model MetricsSummaryResponse
type MetricsSummaryResponse struct {

	// metrics
	Metrics []interface{} `json:"metrics"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this metrics summary response
func (m *MetricsSummaryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetrics(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsSummaryResponse) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	for i := 0; i < len(m.Metrics); i++ {

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsSummaryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsSummaryResponse) UnmarshalBinary(b []byte) error {
	var res MetricsSummaryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
