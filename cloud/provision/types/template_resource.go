// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TemplateResource template resource
// swagger:model TemplateResource
type TemplateResource struct {

	// aws instance
	AwsInstance AWSInstance `json:"aws_instance,omitempty"`

	// aws internet gateway
	AwsInternetGateway AWSInternetGateway `json:"aws_internet_gateway,omitempty"`

	// aws main route table association
	AwsMainRouteTableAssociation AWSMainRouteTableAssociation `json:"aws_main_route_table_association,omitempty"`

	// aws route table
	AwsRouteTable AWSRouteTable `json:"aws_route_table,omitempty"`

	// aws security group
	AwsSecurityGroup AWSSecurityGroup `json:"aws_security_group,omitempty"`

	// aws subnet
	AwsSubnet AWSSubnet `json:"aws_subnet,omitempty"`

	// aws vpc
	AwsVpc AWSVPC `json:"aws_vpc,omitempty"`

	// digitalocean droplet
	DigitaloceanDroplet DigitalOceanDropletMap `json:"digitalocean_droplet,omitempty"`

	// google compute firewall
	GoogleComputeFirewall GoogleComputeFirewall `json:"google_compute_firewall,omitempty"`

	// google compute instance
	GoogleComputeInstance GoogleComputeInstance `json:"google_compute_instance,omitempty"`

	// google compute network
	GoogleComputeNetwork GoogleComputeNetwork `json:"google_compute_network,omitempty"`

	// google compute subnetwork
	GoogleComputeSubnetwork GoogleComputeSubnetwork `json:"google_compute_subnetwork,omitempty"`

	// packet device
	PacketDevice PacketDeviceMap `json:"packet_device,omitempty"`
}

// Validate validates this template resource
func (m *TemplateResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsInternetGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsRouteTable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsSecurityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsSubnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsVpc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigitaloceanDroplet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePacketDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateResource) validateAwsInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsInstance) { // not required
		return nil
	}

	if err := m.AwsInstance.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aws_instance")
		}
		return err
	}

	return nil
}

func (m *TemplateResource) validateAwsInternetGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsInternetGateway) { // not required
		return nil
	}

	if err := m.AwsInternetGateway.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aws_internet_gateway")
		}
		return err
	}

	return nil
}

func (m *TemplateResource) validateAwsRouteTable(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsRouteTable) { // not required
		return nil
	}

	if err := m.AwsRouteTable.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aws_route_table")
		}
		return err
	}

	return nil
}

func (m *TemplateResource) validateAwsSecurityGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsSecurityGroup) { // not required
		return nil
	}

	if err := m.AwsSecurityGroup.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aws_security_group")
		}
		return err
	}

	return nil
}

func (m *TemplateResource) validateAwsSubnet(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsSubnet) { // not required
		return nil
	}

	if err := m.AwsSubnet.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aws_subnet")
		}
		return err
	}

	return nil
}

func (m *TemplateResource) validateAwsVpc(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsVpc) { // not required
		return nil
	}

	if err := m.AwsVpc.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aws_vpc")
		}
		return err
	}

	return nil
}

func (m *TemplateResource) validateDigitaloceanDroplet(formats strfmt.Registry) error {

	if swag.IsZero(m.DigitaloceanDroplet) { // not required
		return nil
	}

	if err := m.DigitaloceanDroplet.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("digitalocean_droplet")
		}
		return err
	}

	return nil
}

func (m *TemplateResource) validatePacketDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.PacketDevice) { // not required
		return nil
	}

	if err := m.PacketDevice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("packet_device")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateResource) UnmarshalBinary(b []byte) error {
	var res TemplateResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
