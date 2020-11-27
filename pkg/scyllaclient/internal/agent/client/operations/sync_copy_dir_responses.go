// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/agent/models"
)

// SyncCopyDirReader is a Reader for the SyncCopyDir structure.
type SyncCopyDirReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncCopyDirReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncCopyDirOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSyncCopyDirDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSyncCopyDirOK creates a SyncCopyDirOK with default headers values
func NewSyncCopyDirOK() *SyncCopyDirOK {
	return &SyncCopyDirOK{}
}

/*SyncCopyDirOK handles this case with default header values.

Job ID
*/
type SyncCopyDirOK struct {
	Payload *models.Jobid
	JobID   int64
}

func (o *SyncCopyDirOK) GetPayload() *models.Jobid {
	return o.Payload
}

func (o *SyncCopyDirOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Jobid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

// NewSyncCopyDirDefault creates a SyncCopyDirDefault with default headers values
func NewSyncCopyDirDefault(code int) *SyncCopyDirDefault {
	return &SyncCopyDirDefault{
		_statusCode: code,
	}
}

/*SyncCopyDirDefault handles this case with default header values.

Server error
*/
type SyncCopyDirDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
	JobID   int64
}

// Code gets the status code for the sync copy dir default response
func (o *SyncCopyDirDefault) Code() int {
	return o._statusCode
}

func (o *SyncCopyDirDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SyncCopyDirDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

func (o *SyncCopyDirDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
