// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StreamRule stream rule
// swagger:model StreamRule
type StreamRule struct {

	// content pack
	ContentPack string `json:"content_pack,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// field
	Field string `json:"field,omitempty"`

	// fields
	Fields interface{} `json:"fields,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// inverted
	Inverted bool `json:"inverted,omitempty"`

	// stream id
	StreamID string `json:"stream_id,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// validations
	Validations interface{} `json:"validations,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this stream rule
func (m *StreamRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var streamRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXACT","REGEX","GREATER","SMALLER","PRESENCE","CONTAINS","ALWAYS_MATCH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		streamRuleTypeTypePropEnum = append(streamRuleTypeTypePropEnum, v)
	}
}

const (
	// StreamRuleTypeEXACT captures enum value "EXACT"
	StreamRuleTypeEXACT string = "EXACT"
	// StreamRuleTypeREGEX captures enum value "REGEX"
	StreamRuleTypeREGEX string = "REGEX"
	// StreamRuleTypeGREATER captures enum value "GREATER"
	StreamRuleTypeGREATER string = "GREATER"
	// StreamRuleTypeSMALLER captures enum value "SMALLER"
	StreamRuleTypeSMALLER string = "SMALLER"
	// StreamRuleTypePRESENCE captures enum value "PRESENCE"
	StreamRuleTypePRESENCE string = "PRESENCE"
	// StreamRuleTypeCONTAINS captures enum value "CONTAINS"
	StreamRuleTypeCONTAINS string = "CONTAINS"
	// StreamRuleTypeALWAYSMATCH captures enum value "ALWAYS_MATCH"
	StreamRuleTypeALWAYSMATCH string = "ALWAYS_MATCH"
)

// prop value enum
func (m *StreamRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, streamRuleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StreamRule) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StreamRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StreamRule) UnmarshalBinary(b []byte) error {
	var res StreamRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
