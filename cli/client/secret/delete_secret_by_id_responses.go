// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wuhoops/silenda/cli/models"
)

// DeleteSecretByIDReader is a Reader for the DeleteSecretByID structure.
type DeleteSecretByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecretByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSecretByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSecretByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSecretByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /secret.deleteSecretByID] deleteSecretById", response, response.Code())
	}
}

// NewDeleteSecretByIDOK creates a DeleteSecretByIDOK with default headers values
func NewDeleteSecretByIDOK() *DeleteSecretByIDOK {
	return &DeleteSecretByIDOK{}
}

/*
DeleteSecretByIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteSecretByIDOK struct {
	Payload *models.HandlersResponseString
}

// IsSuccess returns true when this delete secret by Id o k response has a 2xx status code
func (o *DeleteSecretByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete secret by Id o k response has a 3xx status code
func (o *DeleteSecretByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete secret by Id o k response has a 4xx status code
func (o *DeleteSecretByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete secret by Id o k response has a 5xx status code
func (o *DeleteSecretByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete secret by Id o k response a status code equal to that given
func (o *DeleteSecretByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete secret by Id o k response
func (o *DeleteSecretByIDOK) Code() int {
	return 200
}

func (o *DeleteSecretByIDOK) Error() string {
	return fmt.Sprintf("[POST /secret.deleteSecretByID][%d] deleteSecretByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteSecretByIDOK) String() string {
	return fmt.Sprintf("[POST /secret.deleteSecretByID][%d] deleteSecretByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteSecretByIDOK) GetPayload() *models.HandlersResponseString {
	return o.Payload
}

func (o *DeleteSecretByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersResponseString)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSecretByIDBadRequest creates a DeleteSecretByIDBadRequest with default headers values
func NewDeleteSecretByIDBadRequest() *DeleteSecretByIDBadRequest {
	return &DeleteSecretByIDBadRequest{}
}

/*
DeleteSecretByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteSecretByIDBadRequest struct {
	Payload *models.HandlersErrResponse
}

// IsSuccess returns true when this delete secret by Id bad request response has a 2xx status code
func (o *DeleteSecretByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete secret by Id bad request response has a 3xx status code
func (o *DeleteSecretByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete secret by Id bad request response has a 4xx status code
func (o *DeleteSecretByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete secret by Id bad request response has a 5xx status code
func (o *DeleteSecretByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete secret by Id bad request response a status code equal to that given
func (o *DeleteSecretByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete secret by Id bad request response
func (o *DeleteSecretByIDBadRequest) Code() int {
	return 400
}

func (o *DeleteSecretByIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /secret.deleteSecretByID][%d] deleteSecretByIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSecretByIDBadRequest) String() string {
	return fmt.Sprintf("[POST /secret.deleteSecretByID][%d] deleteSecretByIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSecretByIDBadRequest) GetPayload() *models.HandlersErrResponse {
	return o.Payload
}

func (o *DeleteSecretByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSecretByIDInternalServerError creates a DeleteSecretByIDInternalServerError with default headers values
func NewDeleteSecretByIDInternalServerError() *DeleteSecretByIDInternalServerError {
	return &DeleteSecretByIDInternalServerError{}
}

/*
DeleteSecretByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteSecretByIDInternalServerError struct {
	Payload *models.HandlersErrResponse
}

// IsSuccess returns true when this delete secret by Id internal server error response has a 2xx status code
func (o *DeleteSecretByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete secret by Id internal server error response has a 3xx status code
func (o *DeleteSecretByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete secret by Id internal server error response has a 4xx status code
func (o *DeleteSecretByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete secret by Id internal server error response has a 5xx status code
func (o *DeleteSecretByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete secret by Id internal server error response a status code equal to that given
func (o *DeleteSecretByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete secret by Id internal server error response
func (o *DeleteSecretByIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteSecretByIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /secret.deleteSecretByID][%d] deleteSecretByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSecretByIDInternalServerError) String() string {
	return fmt.Sprintf("[POST /secret.deleteSecretByID][%d] deleteSecretByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSecretByIDInternalServerError) GetPayload() *models.HandlersErrResponse {
	return o.Payload
}

func (o *DeleteSecretByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}