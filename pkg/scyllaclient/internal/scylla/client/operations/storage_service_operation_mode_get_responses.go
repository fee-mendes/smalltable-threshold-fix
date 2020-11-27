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

// StorageServiceOperationModeGetReader is a Reader for the StorageServiceOperationModeGet structure.
type StorageServiceOperationModeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceOperationModeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceOperationModeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceOperationModeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceOperationModeGetOK creates a StorageServiceOperationModeGetOK with default headers values
func NewStorageServiceOperationModeGetOK() *StorageServiceOperationModeGetOK {
	return &StorageServiceOperationModeGetOK{}
}

/*StorageServiceOperationModeGetOK handles this case with default header values.

StorageServiceOperationModeGetOK storage service operation mode get o k
*/
type StorageServiceOperationModeGetOK struct {
	Payload string
}

func (o *StorageServiceOperationModeGetOK) GetPayload() string {
	return o.Payload
}

func (o *StorageServiceOperationModeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceOperationModeGetDefault creates a StorageServiceOperationModeGetDefault with default headers values
func NewStorageServiceOperationModeGetDefault(code int) *StorageServiceOperationModeGetDefault {
	return &StorageServiceOperationModeGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceOperationModeGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceOperationModeGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service operation mode get default response
func (o *StorageServiceOperationModeGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceOperationModeGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceOperationModeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceOperationModeGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
