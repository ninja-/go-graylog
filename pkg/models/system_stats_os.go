// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SystemStatsOs system stats os
// swagger:model systemStatsOs
type SystemStatsOs struct {

	// load average
	LoadAverage []float64 `json:"load_average"`

	// memory
	Memory *SystemStatsOsMemory `json:"memory,omitempty"`

	// processor
	Processor *SystemStatsOsProcessor `json:"processor,omitempty"`

	// swap
	Swap *SystemStatsOsSwap `json:"swap,omitempty"`

	// uptime
	Uptime int64 `json:"uptime,omitempty"`
}

// Validate validates this system stats os
func (m *SystemStatsOs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadAverage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProcessor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSwap(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemStatsOs) validateLoadAverage(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadAverage) { // not required
		return nil
	}

	return nil
}

func (m *SystemStatsOs) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {

		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

func (m *SystemStatsOs) validateProcessor(formats strfmt.Registry) error {

	if swag.IsZero(m.Processor) { // not required
		return nil
	}

	if m.Processor != nil {

		if err := m.Processor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processor")
			}
			return err
		}
	}

	return nil
}

func (m *SystemStatsOs) validateSwap(formats strfmt.Registry) error {

	if swag.IsZero(m.Swap) { // not required
		return nil
	}

	if m.Swap != nil {

		if err := m.Swap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemStatsOs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemStatsOs) UnmarshalBinary(b []byte) error {
	var res SystemStatsOs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
