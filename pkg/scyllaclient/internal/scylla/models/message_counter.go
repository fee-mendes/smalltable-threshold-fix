// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MessageCounter message_counter
//
// Holds command counters
//
// swagger:model message_counter
type MessageCounter struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this message counter
func (m *MessageCounter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MessageCounter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageCounter) UnmarshalBinary(b []byte) error {
	var res MessageCounter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
