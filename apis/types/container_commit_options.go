// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerCommitOptions options of committing a container into an image
// swagger:model ContainerCommitOptions
type ContainerCommitOptions struct {

	// author is the one build the image
	Author string `json:"Author,omitempty"`

	// comment is external information add for the image
	Comment string `json:"Comment,omitempty"`

	// repository is the image name
	Repository string `json:"Repository,omitempty"`

	// tag is the image tag
	Tag string `json:"Tag,omitempty"`
}

// Validate validates this container commit options
func (m *ContainerCommitOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerCommitOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerCommitOptions) UnmarshalBinary(b []byte) error {
	var res ContainerCommitOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
