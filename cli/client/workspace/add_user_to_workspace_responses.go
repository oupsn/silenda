// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wuhoops/silenda/cli/models"
)

// AddUserToWorkspaceReader is a Reader for the AddUserToWorkspace structure.
type AddUserToWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserToWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUserToWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddUserToWorkspaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddUserToWorkspaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /workspace.addUserToWorkspace] addUserToWorkspace", response, response.Code())
	}
}

// NewAddUserToWorkspaceOK creates a AddUserToWorkspaceOK with default headers values
func NewAddUserToWorkspaceOK() *AddUserToWorkspaceOK {
	return &AddUserToWorkspaceOK{}
}

/*
AddUserToWorkspaceOK describes a response with status code 200, with default header values.

OK
*/
type AddUserToWorkspaceOK struct {
	Payload *models.HandlersResponseString
}

// IsSuccess returns true when this add user to workspace o k response has a 2xx status code
func (o *AddUserToWorkspaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add user to workspace o k response has a 3xx status code
func (o *AddUserToWorkspaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to workspace o k response has a 4xx status code
func (o *AddUserToWorkspaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user to workspace o k response has a 5xx status code
func (o *AddUserToWorkspaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to workspace o k response a status code equal to that given
func (o *AddUserToWorkspaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add user to workspace o k response
func (o *AddUserToWorkspaceOK) Code() int {
	return 200
}

func (o *AddUserToWorkspaceOK) Error() string {
	return fmt.Sprintf("[POST /workspace.addUserToWorkspace][%d] addUserToWorkspaceOK  %+v", 200, o.Payload)
}

func (o *AddUserToWorkspaceOK) String() string {
	return fmt.Sprintf("[POST /workspace.addUserToWorkspace][%d] addUserToWorkspaceOK  %+v", 200, o.Payload)
}

func (o *AddUserToWorkspaceOK) GetPayload() *models.HandlersResponseString {
	return o.Payload
}

func (o *AddUserToWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersResponseString)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToWorkspaceBadRequest creates a AddUserToWorkspaceBadRequest with default headers values
func NewAddUserToWorkspaceBadRequest() *AddUserToWorkspaceBadRequest {
	return &AddUserToWorkspaceBadRequest{}
}

/*
AddUserToWorkspaceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddUserToWorkspaceBadRequest struct {
	Payload *models.HandlersErrResponse
}

// IsSuccess returns true when this add user to workspace bad request response has a 2xx status code
func (o *AddUserToWorkspaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user to workspace bad request response has a 3xx status code
func (o *AddUserToWorkspaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to workspace bad request response has a 4xx status code
func (o *AddUserToWorkspaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user to workspace bad request response has a 5xx status code
func (o *AddUserToWorkspaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to workspace bad request response a status code equal to that given
func (o *AddUserToWorkspaceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add user to workspace bad request response
func (o *AddUserToWorkspaceBadRequest) Code() int {
	return 400
}

func (o *AddUserToWorkspaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /workspace.addUserToWorkspace][%d] addUserToWorkspaceBadRequest  %+v", 400, o.Payload)
}

func (o *AddUserToWorkspaceBadRequest) String() string {
	return fmt.Sprintf("[POST /workspace.addUserToWorkspace][%d] addUserToWorkspaceBadRequest  %+v", 400, o.Payload)
}

func (o *AddUserToWorkspaceBadRequest) GetPayload() *models.HandlersErrResponse {
	return o.Payload
}

func (o *AddUserToWorkspaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToWorkspaceInternalServerError creates a AddUserToWorkspaceInternalServerError with default headers values
func NewAddUserToWorkspaceInternalServerError() *AddUserToWorkspaceInternalServerError {
	return &AddUserToWorkspaceInternalServerError{}
}

/*
AddUserToWorkspaceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AddUserToWorkspaceInternalServerError struct {
	Payload *models.HandlersErrResponse
}

// IsSuccess returns true when this add user to workspace internal server error response has a 2xx status code
func (o *AddUserToWorkspaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user to workspace internal server error response has a 3xx status code
func (o *AddUserToWorkspaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to workspace internal server error response has a 4xx status code
func (o *AddUserToWorkspaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user to workspace internal server error response has a 5xx status code
func (o *AddUserToWorkspaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add user to workspace internal server error response a status code equal to that given
func (o *AddUserToWorkspaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add user to workspace internal server error response
func (o *AddUserToWorkspaceInternalServerError) Code() int {
	return 500
}

func (o *AddUserToWorkspaceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /workspace.addUserToWorkspace][%d] addUserToWorkspaceInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUserToWorkspaceInternalServerError) String() string {
	return fmt.Sprintf("[POST /workspace.addUserToWorkspace][%d] addUserToWorkspaceInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUserToWorkspaceInternalServerError) GetPayload() *models.HandlersErrResponse {
	return o.Payload
}

func (o *AddUserToWorkspaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
