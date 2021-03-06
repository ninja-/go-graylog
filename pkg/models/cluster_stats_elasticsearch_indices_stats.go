// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterStatsElasticsearchIndicesStats cluster stats elasticsearch indices stats
// swagger:model clusterStatsElasticsearchIndicesStats
type ClusterStatsElasticsearchIndicesStats struct {

	// field data size
	FieldDataSize int64 `json:"field_data_size,omitempty"`

	// id cache size
	IDCacheSize int64 `json:"id_cache_size,omitempty"`

	// index count
	IndexCount int64 `json:"index_count,omitempty"`

	// store size
	StoreSize int64 `json:"store_size,omitempty"`
}

// Validate validates this cluster stats elasticsearch indices stats
func (m *ClusterStatsElasticsearchIndicesStats) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatsElasticsearchIndicesStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatsElasticsearchIndicesStats) UnmarshalBinary(b []byte) error {
	var res ClusterStatsElasticsearchIndicesStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
