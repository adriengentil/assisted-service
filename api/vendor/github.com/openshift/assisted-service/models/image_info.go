// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	timeext "time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageInfo image info
//
// swagger:model image_info
type ImageInfo struct {

	// created at
	// Format: date-time
	CreatedAt timeext.Time `json:"created_at,omitempty" gorm:"type:timestamp with time zone"`

	// download url
	DownloadURL string `json:"download_url,omitempty"`

	// expires at
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty" gorm:"type:timestamp with time zone"`

	// Image generator version.
	GeneratorVersion string `json:"generator_version,omitempty"`

	// size bytes
	// Minimum: 0
	SizeBytes *int64 `json:"size_bytes,omitempty"`

	// SSH public key for debugging the installation.
	SSHPublicKey string `json:"ssh_public_key,omitempty"`

	// static network configuration string in the format expected by discovery ignition generation
	StaticNetworkConfig string `json:"static_network_config,omitempty"`

	// type
	Type ImageType `json:"type,omitempty"`
}

// Validate validates this image info
func (m *ImageInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizeBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageInfo) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageInfo) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ImageInfo) validateSizeBytes(formats strfmt.Registry) error {
	if swag.IsZero(m.SizeBytes) { // not required
		return nil
	}

	if err := validate.MinimumInt("size_bytes", "body", *m.SizeBytes, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ImageInfo) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this image info based on the context it is used
func (m *ImageInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageInfo) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageInfo) UnmarshalBinary(b []byte) error {
	var res ImageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}