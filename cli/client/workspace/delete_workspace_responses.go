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

// DeleteWorkspaceReader is a Reader for the DeleteWorkspace structure.
type DeleteWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWorkspaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWorkspaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /workspace.deleteWorkspace/] deleteWorkspace", response, response.Code())
	}
}

// NewDeleteWorkspaceOK creates a DeleteWorkspaceOK with default headers values
func NewDeleteWorkspaceOK() *DeleteWorkspaceOK {
	return &DeleteWorkspaceOK{}
}

/*
DeleteWorkspaceOK describes a response with status code 200, with default header values.

OK
*/
type DeleteWorkspaceOK struct {
	Payload *models.HandlersResponseString
}

// IsSuccess returns true when this delete workspace o k response has a 2xx status code
func (o *DeleteWorkspaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete workspace o k response has a 3xx status code
func (o *DeleteWorkspaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workspace o k response has a 4xx status code
func (o *DeleteWorkspaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete workspace o k response has a 5xx status code
func (o *DeleteWorkspaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workspace o k response a status code equal to that given
func (o *DeleteWorkspaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete workspace o k response
func (o *DeleteWorkspaceOK) Code() int {
	return 200
}

func (o *DeleteWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /workspace.deleteWorkspace/][%d] deleteWorkspaceOK  %+v", 200, o.Payload)
}

func (o *DeleteWorkspaceOK) String() string {
	return fmt.Sprintf("[DELETE /workspace.deleteWorkspace/][%d] deleteWorkspaceOK  %+v", 200, o.Payload)
}

func (o *DeleteWorkspaceOK) GetPayload() *models.HandlersResponseString {
	return o.Payload
}

func (o *DeleteWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersResponseString)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkspaceBadRequest creates a DeleteWorkspaceBadRequest with default headers values
func NewDeleteWorkspaceBadRequest() *DeleteWorkspaceBadRequest {
	return &DeleteWorkspaceBadRequest{}
}

/*
DeleteWorkspaceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteWorkspaceBadRequest struct {
	Payload *models.HandlersErrResponse
}

// IsSuccess returns true when this delete workspace bad request response has a 2xx status code
func (o *DeleteWorkspaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workspace bad request response has a 3xx status code
func (o *DeleteWorkspaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workspace bad request response has a 4xx status code
func (o *DeleteWorkspaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workspace bad request response has a 5xx status code
func (o *DeleteWorkspaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workspace bad request response a status code equal to that given
func (o *DeleteWorkspaceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete workspace bad request response
func (o *DeleteWorkspaceBadRequest) Code() int {
	return 400
}

func (o *DeleteWorkspaceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /workspace.deleteWorkspace/][%d] deleteWorkspaceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkspaceBadRequest) String() string {
	return fmt.Sprintf("[DELETE /workspace.deleteWorkspace/][%d] deleteWorkspaceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkspaceBadRequest) GetPayload() *models.HandlersErrResponse {
	return o.Payload
}

func (o *DeleteWorkspaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkspaceInternalServerError creates a DeleteWorkspaceInternalServerError with default headers values
func NewDeleteWorkspaceInternalServerError() *DeleteWorkspaceInternalServerError {
	return &DeleteWorkspaceInternalServerError{}
}

/*
DeleteWorkspaceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteWorkspaceInternalServerError struct {
	Payload *models.HandlersErrResponse
}

// IsSuccess returns true when this delete workspace internal server error response has a 2xx status code
func (o *DeleteWorkspaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workspace internal server error response has a 3xx status code
func (o *DeleteWorkspaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workspace internal server error response has a 4xx status code
func (o *DeleteWorkspaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete workspace internal server error response has a 5xx status code
func (o *DeleteWorkspaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete workspace internal server error response a status code equal to that given
func (o *DeleteWorkspaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete workspace internal server error response
func (o *DeleteWorkspaceInternalServerError) Code() int {
	return 500
}

func (o *DeleteWorkspaceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /workspace.deleteWorkspace/][%d] deleteWorkspaceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkspaceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /workspace.deleteWorkspace/][%d] deleteWorkspaceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkspaceInternalServerError) GetPayload() *models.HandlersErrResponse {
	return o.Payload
}

func (o *DeleteWorkspaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
