// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterStatsMongoHostInfoExtra cluster stats mongo host info extra
// swagger:model clusterStatsMongoHostInfoExtra
type ClusterStatsMongoHostInfoExtra struct {

	// cpu features
	CPUFeatures string `json:"cpu_features,omitempty"`

	// cpu frequency mhz
	CPUFrequencyMhz string `json:"cpu_frequency_mhz,omitempty"`

	// kernel version
	KernelVersion string `json:"kernel_version,omitempty"`

	// libc version
	LibcVersion string `json:"libc_version,omitempty"`

	// max open files
	MaxOpenFiles int64 `json:"max_open_files,omitempty"`

	// num pages
	NumPages int64 `json:"num_pages,omitempty"`

	// page size
	PageSize int64 `json:"page_size,omitempty"`

	// scheduler
	Scheduler string `json:"scheduler,omitempty"`

	// version string
	VersionString string `json:"version_string,omitempty"`
}

// Validate validates this cluster stats mongo host info extra
func (m *ClusterStatsMongoHostInfoExtra) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatsMongoHostInfoExtra) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatsMongoHostInfoExtra) UnmarshalBinary(b []byte) error {
	var res ClusterStatsMongoHostInfoExtra
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
