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

// CompactionManagerMetricsCompletedTasksGetReader is a Reader for the CompactionManagerMetricsCompletedTasksGet structure.
type CompactionManagerMetricsCompletedTasksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompactionManagerMetricsCompletedTasksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompactionManagerMetricsCompletedTasksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCompactionManagerMetricsCompletedTasksGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompactionManagerMetricsCompletedTasksGetOK creates a CompactionManagerMetricsCompletedTasksGetOK with default headers values
func NewCompactionManagerMetricsCompletedTasksGetOK() *CompactionManagerMetricsCompletedTasksGetOK {
	return &CompactionManagerMetricsCompletedTasksGetOK{}
}

/*CompactionManagerMetricsCompletedTasksGetOK handles this case with default header values.

CompactionManagerMetricsCompletedTasksGetOK compaction manager metrics completed tasks get o k
*/
type CompactionManagerMetricsCompletedTasksGetOK struct {
	Payload interface{}
}

func (o *CompactionManagerMetricsCompletedTasksGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CompactionManagerMetricsCompletedTasksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompactionManagerMetricsCompletedTasksGetDefault creates a CompactionManagerMetricsCompletedTasksGetDefault with default headers values
func NewCompactionManagerMetricsCompletedTasksGetDefault(code int) *CompactionManagerMetricsCompletedTasksGetDefault {
	return &CompactionManagerMetricsCompletedTasksGetDefault{
		_statusCode: code,
	}
}

/*CompactionManagerMetricsCompletedTasksGetDefault handles this case with default header values.

internal server error
*/
type CompactionManagerMetricsCompletedTasksGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the compaction manager metrics completed tasks get default response
func (o *CompactionManagerMetricsCompletedTasksGetDefault) Code() int {
	return o._statusCode
}

func (o *CompactionManagerMetricsCompletedTasksGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CompactionManagerMetricsCompletedTasksGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CompactionManagerMetricsCompletedTasksGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
