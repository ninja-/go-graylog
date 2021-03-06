// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterStatsMongoDatabaseStatsExtentFreeList cluster stats mongo database stats extent free list
// swagger:model clusterStatsMongoDatabaseStatsExtentFreeList
type ClusterStatsMongoDatabaseStatsExtentFreeList struct {

	// num
	Num int64 `json:"num,omitempty"`

	// total size
	TotalSize int64 `json:"total_size,omitempty"`
}

// Validate validates this cluster stats mongo database stats extent free list
func (m *ClusterStatsMongoDatabaseStatsExtentFreeList) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatsMongoDatabaseStatsExtentFreeList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatsMongoDatabaseStatsExtentFreeList) UnmarshalBinary(b []byte) error {
	var res ClusterStatsMongoDatabaseStatsExtentFreeList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
