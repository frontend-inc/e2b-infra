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

// DescribeBalloonStatsReader is a Reader for the DescribeBalloonStats structure.
type DescribeBalloonStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeBalloonStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeBalloonStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDescribeBalloonStatsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDescribeBalloonStatsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeBalloonStatsOK creates a DescribeBalloonStatsOK with default headers values
func NewDescribeBalloonStatsOK() *DescribeBalloonStatsOK {
	return &DescribeBalloonStatsOK{}
}

/*
DescribeBalloonStatsOK describes a response with status code 200, with default header values.

The balloon device statistics
*/
type DescribeBalloonStatsOK struct {
	Payload *models.BalloonStats
}

// IsSuccess returns true when this describe balloon stats o k response has a 2xx status code
func (o *DescribeBalloonStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe balloon stats o k response has a 3xx status code
func (o *DescribeBalloonStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe balloon stats o k response has a 4xx status code
func (o *DescribeBalloonStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe balloon stats o k response has a 5xx status code
func (o *DescribeBalloonStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe balloon stats o k response a status code equal to that given
func (o *DescribeBalloonStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe balloon stats o k response
func (o *DescribeBalloonStatsOK) Code() int {
	return 200
}

func (o *DescribeBalloonStatsOK) Error() string {
	return fmt.Sprintf("[GET /balloon/statistics][%d] describeBalloonStatsOK  %+v", 200, o.Payload)
}

func (o *DescribeBalloonStatsOK) String() string {
	return fmt.Sprintf("[GET /balloon/statistics][%d] describeBalloonStatsOK  %+v", 200, o.Payload)
}

func (o *DescribeBalloonStatsOK) GetPayload() *models.BalloonStats {
	return o.Payload
}

func (o *DescribeBalloonStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BalloonStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeBalloonStatsBadRequest creates a DescribeBalloonStatsBadRequest with default headers values
func NewDescribeBalloonStatsBadRequest() *DescribeBalloonStatsBadRequest {
	return &DescribeBalloonStatsBadRequest{}
}

/*
DescribeBalloonStatsBadRequest describes a response with status code 400, with default header values.

The balloon device statistics were not enabled when the device was configured.
*/
type DescribeBalloonStatsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this describe balloon stats bad request response has a 2xx status code
func (o *DescribeBalloonStatsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe balloon stats bad request response has a 3xx status code
func (o *DescribeBalloonStatsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe balloon stats bad request response has a 4xx status code
func (o *DescribeBalloonStatsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe balloon stats bad request response has a 5xx status code
func (o *DescribeBalloonStatsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this describe balloon stats bad request response a status code equal to that given
func (o *DescribeBalloonStatsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the describe balloon stats bad request response
func (o *DescribeBalloonStatsBadRequest) Code() int {
	return 400
}

func (o *DescribeBalloonStatsBadRequest) Error() string {
	return fmt.Sprintf("[GET /balloon/statistics][%d] describeBalloonStatsBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeBalloonStatsBadRequest) String() string {
	return fmt.Sprintf("[GET /balloon/statistics][%d] describeBalloonStatsBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeBalloonStatsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeBalloonStatsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeBalloonStatsDefault creates a DescribeBalloonStatsDefault with default headers values
func NewDescribeBalloonStatsDefault(code int) *DescribeBalloonStatsDefault {
	return &DescribeBalloonStatsDefault{
		_statusCode: code,
	}
}

/*
DescribeBalloonStatsDefault describes a response with status code -1, with default header values.

Internal Server Error
*/
type DescribeBalloonStatsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe balloon stats default response has a 2xx status code
func (o *DescribeBalloonStatsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe balloon stats default response has a 3xx status code
func (o *DescribeBalloonStatsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe balloon stats default response has a 4xx status code
func (o *DescribeBalloonStatsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe balloon stats default response has a 5xx status code
func (o *DescribeBalloonStatsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe balloon stats default response a status code equal to that given
func (o *DescribeBalloonStatsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe balloon stats default response
func (o *DescribeBalloonStatsDefault) Code() int {
	return o._statusCode
}

func (o *DescribeBalloonStatsDefault) Error() string {
	return fmt.Sprintf("[GET /balloon/statistics][%d] describeBalloonStats default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeBalloonStatsDefault) String() string {
	return fmt.Sprintf("[GET /balloon/statistics][%d] describeBalloonStats default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeBalloonStatsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeBalloonStatsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
