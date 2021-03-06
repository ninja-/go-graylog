// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OsStats os stats
// swagger:model OsStats
type OsStats struct {

	// load average
	LoadAverage []float64 `json:"load_average"`

	// memory
	Memory *OsStatsMemory `json:"memory,omitempty"`

	// processor
	Processor *OsStatsProcessor `json:"processor,omitempty"`

	// swap
	Swap *OsStatsSwap `json:"swap,omitempty"`

	// uptime
	Uptime int64 `json:"uptime,omitempty"`
}

// Validate validates this os stats
func (m *OsStats) Validate(formats strfmt.Registry) error {
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

func (m *OsStats) validateLoadAverage(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadAverage) { // not required
		return nil
	}

	return nil
}

func (m *OsStats) validateMemory(formats strfmt.Registry) error {

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

func (m *OsStats) validateProcessor(formats strfmt.Registry) error {

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

func (m *OsStats) validateSwap(formats strfmt.Registry) error {

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
func (m *OsStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsStats) UnmarshalBinary(b []byte) error {
	var res OsStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
