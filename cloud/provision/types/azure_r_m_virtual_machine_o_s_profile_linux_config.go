// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AzureRMVirtualMachineOSProfileLinuxConfig azure r m virtual machine o s profile linux config
// swagger:model AzureRMVirtualMachineOSProfileLinuxConfig
type AzureRMVirtualMachineOSProfileLinuxConfig struct {

	// disable password authentication
	DisablePasswordAuthentication bool `json:"disable_password_authentication,omitempty"`

	// ssh keys
	SSHKeys []*AzureRMSSHKey `json:"ssh_keys"`
}

// Validate validates this azure r m virtual machine o s profile linux config
func (m *AzureRMVirtualMachineOSProfileLinuxConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSSHKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureRMVirtualMachineOSProfileLinuxConfig) validateSSHKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.SSHKeys) { // not required
		return nil
	}

	for i := 0; i < len(m.SSHKeys); i++ {
		if swag.IsZero(m.SSHKeys[i]) { // not required
			continue
		}

		if m.SSHKeys[i] != nil {
			if err := m.SSHKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ssh_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureRMVirtualMachineOSProfileLinuxConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureRMVirtualMachineOSProfileLinuxConfig) UnmarshalBinary(b []byte) error {
	var res AzureRMVirtualMachineOSProfileLinuxConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
