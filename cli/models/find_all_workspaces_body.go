// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FindAllWorkspacesBody find all workspaces body
//
// swagger:model FindAllWorkspacesBody
type FindAllWorkspacesBody struct {

	// owner
	Owner string `json:"owner,omitempty"`
}

// Validate validates this find all workspaces body
func (m *FindAllWorkspacesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this find all workspaces body based on context it is used
func (m *FindAllWorkspacesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FindAllWorkspacesBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FindAllWorkspacesBody) UnmarshalBinary(b []byte) error {
	var res FindAllWorkspacesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
