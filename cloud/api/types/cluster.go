// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Cluster A cluster attached or provisioned through Containership
// swagger:model Cluster
type Cluster struct {

	// API server address
	// Required: true
	APIServerAddress *string `json:"api_server_address"`

	// Timestamp at which the cluster was created
	// Required: true
	CreatedAt *string `json:"created_at"`

	// Cluster environment
	// Required: true
	Environment *string `json:"environment"`

	// Cluster ID
	// Required: true
	ID UUID `json:"id"`

	// Cluster name
	// Required: true
	Name *string `json:"name"`

	// Organization ID of the organization the cluster belongs to
	// Required: true
	OrganizationID UUID `json:"organization_id"`

	// Account ID of the cluster owner
	// Required: true
	OwnerID UUID `json:"owner_id"`

	// Name of the provider through which the cluster is provisioned
	// Required: true
	ProviderName *string `json:"provider_name"`

	// Timestamp at which the cluster was updated
	// Required: true
	UpdatedAt *string `json:"updated_at"`

	// Kubernetes version
	// Required: true
	Version *string `json:"version"`

	// Worker nodes address
	// Required: true
	WorkerNodesAddress *string `json:"worker_nodes_address"`
}

// Validate validates this cluster
func (m *Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIServerAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkerNodesAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) validateAPIServerAddress(formats strfmt.Registry) error {

	if err := validate.Required("api_server_address", "body", m.APIServerAddress); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Cluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateOrganizationID(formats strfmt.Registry) error {

	if err := m.OrganizationID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("organization_id")
		}
		return err
	}

	return nil
}

func (m *Cluster) validateOwnerID(formats strfmt.Registry) error {

	if err := m.OwnerID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("owner_id")
		}
		return err
	}

	return nil
}

func (m *Cluster) validateProviderName(formats strfmt.Registry) error {

	if err := validate.Required("provider_name", "body", m.ProviderName); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateWorkerNodesAddress(formats strfmt.Registry) error {

	if err := validate.Required("worker_nodes_address", "body", m.WorkerNodesAddress); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cluster) UnmarshalBinary(b []byte) error {
	var res Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}