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

// RemoveUserFromWorkspaceReader is a Reader for the RemoveUserFromWorkspace structure.
type RemoveUserFromWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveUserFromWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveUserFromWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveUserFromWorkspaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveUserFromWorkspaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /workspace.removeUserFromWorkspace] removeUserFromWorkspace", response, response.Code())
	}
}

// NewRemoveUserFromWorkspaceOK creates a RemoveUserFromWorkspaceOK with default headers values
func NewRemoveUserFromWorkspaceOK() *RemoveUserFromWorkspaceOK {
	return &RemoveUserFromWorkspaceOK{}
}

/*
RemoveUserFromWorkspaceOK describes a response with status code 200, with default header values.

OK
*/
type RemoveUserFromWorkspaceOK struct {
	Payload *models.HandlersResponseString
}

// IsSuccess returns true when this remove user from workspace o k response has a 2xx status code
func (o *RemoveUserFromWorkspaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove user from workspace o k response has a 3xx status code
func (o *RemoveUserFromWorkspaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user from workspace o k response has a 4xx status code
func (o *RemoveUserFromWorkspaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove user from workspace o k response has a 5xx status code
func (o *RemoveUserFromWorkspaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user from workspace o k response a status code equal to that given
func (o *RemoveUserFromWorkspaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove user from workspace o k response
func (o *RemoveUserFromWorkspaceOK) Code() int {
	return 200
}

func (o *RemoveUserFromWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /workspace.removeUserFromWorkspace][%d] removeUserFromWorkspaceOK  %+v", 200, o.Payload)
}

func (o *RemoveUserFromWorkspaceOK) String() string {
	return fmt.Sprintf("[DELETE /workspace.removeUserFromWorkspace][%d] removeUserFromWorkspaceOK  %+v", 200, o.Payload)
}

func (o *RemoveUserFromWorkspaceOK) GetPayload() *models.HandlersResponseString {
	return o.Payload
}

func (o *RemoveUserFromWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersResponseString)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserFromWorkspaceBadRequest creates a RemoveUserFromWorkspaceBadRequest with default headers values
func NewRemoveUserFromWorkspaceBadRequest() *RemoveUserFromWorkspaceBadRequest {
	return &RemoveUserFromWorkspaceBadRequest{}
}

/*
RemoveUserFromWorkspaceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RemoveUserFromWorkspaceBadRequest struct {
	Payload *models.HandlersErrResponse
}

// IsSuccess returns true when this remove user from workspace bad request response has a 2xx status code
func (o *RemoveUserFromWorkspaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user from workspace bad request response has a 3xx status code
func (o *RemoveUserFromWorkspaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user from workspace bad request response has a 4xx status code
func (o *RemoveUserFromWorkspaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user from workspace bad request response has a 5xx status code
func (o *RemoveUserFromWorkspaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user from workspace bad request response a status code equal to that given
func (o *RemoveUserFromWorkspaceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove user from workspace bad request response
func (o *RemoveUserFromWorkspaceBadRequest) Code() int {
	return 400
}

func (o *RemoveUserFromWorkspaceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /workspace.removeUserFromWorkspace][%d] removeUserFromWorkspaceBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveUserFromWorkspaceBadRequest) String() string {
	return fmt.Sprintf("[DELETE /workspace.removeUserFromWorkspace][%d] removeUserFromWorkspaceBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveUserFromWorkspaceBadRequest) GetPayload() *models.HandlersErrResponse {
	return o.Payload
}

func (o *RemoveUserFromWorkspaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserFromWorkspaceInternalServerError creates a RemoveUserFromWorkspaceInternalServerError with default headers values
func NewRemoveUserFromWorkspaceInternalServerError() *RemoveUserFromWorkspaceInternalServerError {
	return &RemoveUserFromWorkspaceInternalServerError{}
}

/*
RemoveUserFromWorkspaceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RemoveUserFromWorkspaceInternalServerError struct {
	Payload *models.HandlersErrResponse
}

// IsSuccess returns true when this remove user from workspace internal server error response has a 2xx status code
func (o *RemoveUserFromWorkspaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user from workspace internal server error response has a 3xx status code
func (o *RemoveUserFromWorkspaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user from workspace internal server error response has a 4xx status code
func (o *RemoveUserFromWorkspaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove user from workspace internal server error response has a 5xx status code
func (o *RemoveUserFromWorkspaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove user from workspace internal server error response a status code equal to that given
func (o *RemoveUserFromWorkspaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove user from workspace internal server error response
func (o *RemoveUserFromWorkspaceInternalServerError) Code() int {
	return 500
}

func (o *RemoveUserFromWorkspaceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /workspace.removeUserFromWorkspace][%d] removeUserFromWorkspaceInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveUserFromWorkspaceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /workspace.removeUserFromWorkspace][%d] removeUserFromWorkspaceInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveUserFromWorkspaceInternalServerError) GetPayload() *models.HandlersErrResponse {
	return o.Payload
}

func (o *RemoveUserFromWorkspaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
