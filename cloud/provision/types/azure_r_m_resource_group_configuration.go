// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AzureRMResourceGroupConfiguration azure r m resource group configuration
// swagger:model AzureRMResourceGroupConfiguration
type AzureRMResourceGroupConfiguration struct {

	// location
	Location string `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this azure r m resource group configuration
func (m *AzureRMResourceGroupConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureRMResourceGroupConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureRMResourceGroupConfiguration) UnmarshalBinary(b []byte) error {
	var res AzureRMResourceGroupConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
