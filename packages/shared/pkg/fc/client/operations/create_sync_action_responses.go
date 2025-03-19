// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/e2b-dev/infra/packages/shared/pkg/fc/models"
)

// CreateSyncActionReader is a Reader for the CreateSyncAction structure.
type CreateSyncActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSyncActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCreateSyncActionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSyncActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateSyncActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSyncActionNoContent creates a CreateSyncActionNoContent with default headers values
func NewCreateSyncActionNoContent() *CreateSyncActionNoContent {
	return &CreateSyncActionNoContent{}
}

/*
CreateSyncActionNoContent describes a response with status code 204, with default header values.

The update was successful
*/
type CreateSyncActionNoContent struct {
}

// IsSuccess returns true when this create sync action no content response has a 2xx status code
func (o *CreateSyncActionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create sync action no content response has a 3xx status code
func (o *CreateSyncActionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sync action no content response has a 4xx status code
func (o *CreateSyncActionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create sync action no content response has a 5xx status code
func (o *CreateSyncActionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create sync action no content response a status code equal to that given
func (o *CreateSyncActionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create sync action no content response
func (o *CreateSyncActionNoContent) Code() int {
	return 204
}

func (o *CreateSyncActionNoContent) Error() string {
	return fmt.Sprintf("[PUT /actions][%d] createSyncActionNoContent ", 204)
}

func (o *CreateSyncActionNoContent) String() string {
	return fmt.Sprintf("[PUT /actions][%d] createSyncActionNoContent ", 204)
}

func (o *CreateSyncActionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSyncActionBadRequest creates a CreateSyncActionBadRequest with default headers values
func NewCreateSyncActionBadRequest() *CreateSyncActionBadRequest {
	return &CreateSyncActionBadRequest{}
}

/*
CreateSyncActionBadRequest describes a response with status code 400, with default header values.

The action cannot be executed due to bad input
*/
type CreateSyncActionBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create sync action bad request response has a 2xx status code
func (o *CreateSyncActionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sync action bad request response has a 3xx status code
func (o *CreateSyncActionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sync action bad request response has a 4xx status code
func (o *CreateSyncActionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create sync action bad request response has a 5xx status code
func (o *CreateSyncActionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create sync action bad request response a status code equal to that given
func (o *CreateSyncActionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create sync action bad request response
func (o *CreateSyncActionBadRequest) Code() int {
	return 400
}

func (o *CreateSyncActionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /actions][%d] createSyncActionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSyncActionBadRequest) String() string {
	return fmt.Sprintf("[PUT /actions][%d] createSyncActionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSyncActionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSyncActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSyncActionDefault creates a CreateSyncActionDefault with default headers values
func NewCreateSyncActionDefault(code int) *CreateSyncActionDefault {
	return &CreateSyncActionDefault{
		_statusCode: code,
	}
}

/*
CreateSyncActionDefault describes a response with status code -1, with default header values.

Internal Server Error
*/
type CreateSyncActionDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create sync action default response has a 2xx status code
func (o *CreateSyncActionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create sync action default response has a 3xx status code
func (o *CreateSyncActionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create sync action default response has a 4xx status code
func (o *CreateSyncActionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create sync action default response has a 5xx status code
func (o *CreateSyncActionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create sync action default response a status code equal to that given
func (o *CreateSyncActionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create sync action default response
func (o *CreateSyncActionDefault) Code() int {
	return o._statusCode
}

func (o *CreateSyncActionDefault) Error() string {
	return fmt.Sprintf("[PUT /actions][%d] createSyncAction default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSyncActionDefault) String() string {
	return fmt.Sprintf("[PUT /actions][%d] createSyncAction default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSyncActionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSyncActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
