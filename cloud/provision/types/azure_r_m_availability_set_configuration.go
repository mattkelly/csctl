// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AzureRMAvailabilitySetConfiguration azure r m availability set configuration
// swagger:model AzureRMAvailabilitySetConfiguration
type AzureRMAvailabilitySetConfiguration struct {

	// location
	Location string `json:"location,omitempty"`

	// managed
	Managed bool `json:"managed,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// platform fault domain count
	PlatformFaultDomainCount int64 `json:"platform_fault_domain_count,omitempty"`

	// resource group name
	ResourceGroupName string `json:"resource_group_name,omitempty"`
}

// Validate validates this azure r m availability set configuration
func (m *AzureRMAvailabilitySetConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureRMAvailabilitySetConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureRMAvailabilitySetConfiguration) UnmarshalBinary(b []byte) error {
	var res AzureRMAvailabilitySetConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}