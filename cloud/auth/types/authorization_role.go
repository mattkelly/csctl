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

// AuthorizationRole authorization role
// swagger:model AuthorizationRole
type AuthorizationRole struct {

	// Timestamp at which the role was created
	// Required: true
	CreatedAt *string `json:"created_at"`

	// Description of the role
	Description string `json:"description,omitempty"`

	// Role ID
	// Required: true
	ID UUID `json:"id"`

	// Name of the role
	// Required: true
	Name *string `json:"name"`

	// Organization ID of the organization the role belongs to
	// Required: true
	OrganizationID UUID `json:"organization_id"`

	// Account ID of the role owner
	// Required: true
	OwnerID UUID `json:"owner_id"`

	// Timestamp at which the role was updated
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this authorization role
func (m *AuthorizationRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
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

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizationRole) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *AuthorizationRole) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *AuthorizationRole) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AuthorizationRole) validateOrganizationID(formats strfmt.Registry) error {

	if err := m.OrganizationID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("organization_id")
		}
		return err
	}

	return nil
}

func (m *AuthorizationRole) validateOwnerID(formats strfmt.Registry) error {

	if err := m.OwnerID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("owner_id")
		}
		return err
	}

	return nil
}

func (m *AuthorizationRole) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthorizationRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizationRole) UnmarshalBinary(b []byte) error {
	var res AuthorizationRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
