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

// PutMetricsReader is a Reader for the PutMetrics structure.
type PutMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutMetricsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMetricsNoContent creates a PutMetricsNoContent with default headers values
func NewPutMetricsNoContent() *PutMetricsNoContent {
	return &PutMetricsNoContent{}
}

/*
PutMetricsNoContent describes a response with status code 204, with default header values.

Metrics system created.
*/
type PutMetricsNoContent struct {
}

// IsSuccess returns true when this put metrics no content response has a 2xx status code
func (o *PutMetricsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put metrics no content response has a 3xx status code
func (o *PutMetricsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put metrics no content response has a 4xx status code
func (o *PutMetricsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put metrics no content response has a 5xx status code
func (o *PutMetricsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put metrics no content response a status code equal to that given
func (o *PutMetricsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put metrics no content response
func (o *PutMetricsNoContent) Code() int {
	return 204
}

func (o *PutMetricsNoContent) Error() string {
	return fmt.Sprintf("[PUT /metrics][%d] putMetricsNoContent ", 204)
}

func (o *PutMetricsNoContent) String() string {
	return fmt.Sprintf("[PUT /metrics][%d] putMetricsNoContent ", 204)
}

func (o *PutMetricsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMetricsBadRequest creates a PutMetricsBadRequest with default headers values
func NewPutMetricsBadRequest() *PutMetricsBadRequest {
	return &PutMetricsBadRequest{}
}

/*
PutMetricsBadRequest describes a response with status code 400, with default header values.

Metrics system cannot be initialized due to bad input request or metrics system already initialized.
*/
type PutMetricsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this put metrics bad request response has a 2xx status code
func (o *PutMetricsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put metrics bad request response has a 3xx status code
func (o *PutMetricsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put metrics bad request response has a 4xx status code
func (o *PutMetricsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put metrics bad request response has a 5xx status code
func (o *PutMetricsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put metrics bad request response a status code equal to that given
func (o *PutMetricsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put metrics bad request response
func (o *PutMetricsBadRequest) Code() int {
	return 400
}

func (o *PutMetricsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /metrics][%d] putMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *PutMetricsBadRequest) String() string {
	return fmt.Sprintf("[PUT /metrics][%d] putMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *PutMetricsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutMetricsDefault creates a PutMetricsDefault with default headers values
func NewPutMetricsDefault(code int) *PutMetricsDefault {
	return &PutMetricsDefault{
		_statusCode: code,
	}
}

/*
PutMetricsDefault describes a response with status code -1, with default header values.

Internal server error.
*/
type PutMetricsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this put metrics default response has a 2xx status code
func (o *PutMetricsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put metrics default response has a 3xx status code
func (o *PutMetricsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put metrics default response has a 4xx status code
func (o *PutMetricsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put metrics default response has a 5xx status code
func (o *PutMetricsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put metrics default response a status code equal to that given
func (o *PutMetricsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put metrics default response
func (o *PutMetricsDefault) Code() int {
	return o._statusCode
}

func (o *PutMetricsDefault) Error() string {
	return fmt.Sprintf("[PUT /metrics][%d] putMetrics default  %+v", o._statusCode, o.Payload)
}

func (o *PutMetricsDefault) String() string {
	return fmt.Sprintf("[PUT /metrics][%d] putMetrics default  %+v", o._statusCode, o.Payload)
}

func (o *PutMetricsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
