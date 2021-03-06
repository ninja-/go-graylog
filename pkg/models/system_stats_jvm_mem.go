// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SystemStatsJvmMem system stats jvm mem
// swagger:model systemStatsJvmMem
type SystemStatsJvmMem struct {

	// direct memory max
	DirectMemoryMax int64 `json:"direct_memory_max,omitempty"`

	// heap init
	HeapInit int64 `json:"heap_init,omitempty"`

	// heap max
	HeapMax int64 `json:"heap_max,omitempty"`

	// non heap init
	NonHeapInit int64 `json:"non_heap_init,omitempty"`

	// non heap max
	NonHeapMax int64 `json:"non_heap_max,omitempty"`
}

// Validate validates this system stats jvm mem
func (m *SystemStatsJvmMem) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SystemStatsJvmMem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemStatsJvmMem) UnmarshalBinary(b []byte) error {
	var res SystemStatsJvmMem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
