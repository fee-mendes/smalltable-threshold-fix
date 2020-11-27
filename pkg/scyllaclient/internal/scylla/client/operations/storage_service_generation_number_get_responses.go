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

// StorageServiceGenerationNumberGetReader is a Reader for the StorageServiceGenerationNumberGet structure.
type StorageServiceGenerationNumberGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceGenerationNumberGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceGenerationNumberGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceGenerationNumberGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceGenerationNumberGetOK creates a StorageServiceGenerationNumberGetOK with default headers values
func NewStorageServiceGenerationNumberGetOK() *StorageServiceGenerationNumberGetOK {
	return &StorageServiceGenerationNumberGetOK{}
}

/*StorageServiceGenerationNumberGetOK handles this case with default header values.

StorageServiceGenerationNumberGetOK storage service generation number get o k
*/
type StorageServiceGenerationNumberGetOK struct {
	Payload int32
}

func (o *StorageServiceGenerationNumberGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageServiceGenerationNumberGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceGenerationNumberGetDefault creates a StorageServiceGenerationNumberGetDefault with default headers values
func NewStorageServiceGenerationNumberGetDefault(code int) *StorageServiceGenerationNumberGetDefault {
	return &StorageServiceGenerationNumberGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceGenerationNumberGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceGenerationNumberGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service generation number get default response
func (o *StorageServiceGenerationNumberGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceGenerationNumberGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceGenerationNumberGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceGenerationNumberGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
