// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla/models"
)

// CommitlogMetricsPendingTasksGetReader is a Reader for the CommitlogMetricsPendingTasksGet structure.
type CommitlogMetricsPendingTasksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitlogMetricsPendingTasksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitlogMetricsPendingTasksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCommitlogMetricsPendingTasksGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCommitlogMetricsPendingTasksGetOK creates a CommitlogMetricsPendingTasksGetOK with default headers values
func NewCommitlogMetricsPendingTasksGetOK() *CommitlogMetricsPendingTasksGetOK {
	return &CommitlogMetricsPendingTasksGetOK{}
}

/*CommitlogMetricsPendingTasksGetOK handles this case with default header values.

CommitlogMetricsPendingTasksGetOK commitlog metrics pending tasks get o k
*/
type CommitlogMetricsPendingTasksGetOK struct {
	Payload interface{}
}

func (o *CommitlogMetricsPendingTasksGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CommitlogMetricsPendingTasksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitlogMetricsPendingTasksGetDefault creates a CommitlogMetricsPendingTasksGetDefault with default headers values
func NewCommitlogMetricsPendingTasksGetDefault(code int) *CommitlogMetricsPendingTasksGetDefault {
	return &CommitlogMetricsPendingTasksGetDefault{
		_statusCode: code,
	}
}

/*CommitlogMetricsPendingTasksGetDefault handles this case with default header values.

internal server error
*/
type CommitlogMetricsPendingTasksGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the commitlog metrics pending tasks get default response
func (o *CommitlogMetricsPendingTasksGetDefault) Code() int {
	return o._statusCode
}

func (o *CommitlogMetricsPendingTasksGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CommitlogMetricsPendingTasksGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CommitlogMetricsPendingTasksGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
