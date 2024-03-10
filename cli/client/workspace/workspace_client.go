// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workspace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workspace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddUserToWorkspace(params *AddUserToWorkspaceParams, opts ...ClientOption) (*AddUserToWorkspaceOK, error)

	CreateWorkspace(params *CreateWorkspaceParams, opts ...ClientOption) (*CreateWorkspaceOK, error)

	DeleteWorkspace(params *DeleteWorkspaceParams, opts ...ClientOption) (*DeleteWorkspaceOK, error)

	FindAllWorkspaces(params *FindAllWorkspacesParams, opts ...ClientOption) (*FindAllWorkspacesOK, error)

	RemoveUserFromWorkspace(params *RemoveUserFromWorkspaceParams, opts ...ClientOption) (*RemoveUserFromWorkspaceOK, error)

	UpdateWorkspace(params *UpdateWorkspaceParams, opts ...ClientOption) (*UpdateWorkspaceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddUserToWorkspace adds user to workspace
*/
func (a *Client) AddUserToWorkspace(params *AddUserToWorkspaceParams, opts ...ClientOption) (*AddUserToWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUserToWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addUserToWorkspace",
		Method:             "POST",
		PathPattern:        "/workspace.addUserToWorkspace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddUserToWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddUserToWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addUserToWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateWorkspace creates a new workspace
*/
func (a *Client) CreateWorkspace(params *CreateWorkspaceParams, opts ...ClientOption) (*CreateWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createWorkspace",
		Method:             "POST",
		PathPattern:        "/workspace.createWorkspace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteWorkspace deletes workspace
*/
func (a *Client) DeleteWorkspace(params *DeleteWorkspaceParams, opts ...ClientOption) (*DeleteWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteWorkspace",
		Method:             "DELETE",
		PathPattern:        "/workspace.deleteWorkspace/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindAllWorkspaces finds all workspaces
*/
func (a *Client) FindAllWorkspaces(params *FindAllWorkspacesParams, opts ...ClientOption) (*FindAllWorkspacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindAllWorkspacesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findAllWorkspaces",
		Method:             "POST",
		PathPattern:        "/workspace.findAllWorkspaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindAllWorkspacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindAllWorkspacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findAllWorkspaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveUserFromWorkspace removes user from workspace
*/
func (a *Client) RemoveUserFromWorkspace(params *RemoveUserFromWorkspaceParams, opts ...ClientOption) (*RemoveUserFromWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveUserFromWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeUserFromWorkspace",
		Method:             "DELETE",
		PathPattern:        "/workspace.removeUserFromWorkspace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveUserFromWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveUserFromWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeUserFromWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateWorkspace updates workspace
*/
func (a *Client) UpdateWorkspace(params *UpdateWorkspaceParams, opts ...ClientOption) (*UpdateWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateWorkspace",
		Method:             "POST",
		PathPattern:        "/workspace.updateWorkspace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
