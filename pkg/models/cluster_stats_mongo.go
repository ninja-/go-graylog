// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterStatsMongo cluster stats mongo
// swagger:model clusterStatsMongo
type ClusterStatsMongo struct {

	// build info
	BuildInfo *ClusterStatsMongoBuildInfo `json:"build_info,omitempty"`

	// database stats
	DatabaseStats *ClusterStatsMongoDatabaseStats `json:"database_stats,omitempty"`

	// host info
	HostInfo *ClusterStatsMongoHostInfo `json:"host_info,omitempty"`

	// server status
	ServerStatus *ClusterStatsMongoServerStatus `json:"server_status,omitempty"`

	// servers
	Servers []string `json:"servers"`
}

// Validate validates this cluster stats mongo
func (m *ClusterStatsMongo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDatabaseStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHostInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServerStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStatsMongo) validateBuildInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildInfo) { // not required
		return nil
	}

	if m.BuildInfo != nil {

		if err := m.BuildInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build_info")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterStatsMongo) validateDatabaseStats(formats strfmt.Registry) error {

	if swag.IsZero(m.DatabaseStats) { // not required
		return nil
	}

	if m.DatabaseStats != nil {

		if err := m.DatabaseStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database_stats")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterStatsMongo) validateHostInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.HostInfo) { // not required
		return nil
	}

	if m.HostInfo != nil {

		if err := m.HostInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host_info")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterStatsMongo) validateServerStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerStatus) { // not required
		return nil
	}

	if m.ServerStatus != nil {

		if err := m.ServerStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server_status")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterStatsMongo) validateServers(formats strfmt.Registry) error {

	if swag.IsZero(m.Servers) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatsMongo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatsMongo) UnmarshalBinary(b []byte) error {
	var res ClusterStatsMongo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
