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

// StorageProxyHintsInProgressGetReader is a Reader for the StorageProxyHintsInProgressGet structure.
type StorageProxyHintsInProgressGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyHintsInProgressGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyHintsInProgressGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyHintsInProgressGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyHintsInProgressGetOK creates a StorageProxyHintsInProgressGetOK with default headers values
func NewStorageProxyHintsInProgressGetOK() *StorageProxyHintsInProgressGetOK {
	return &StorageProxyHintsInProgressGetOK{}
}

/*StorageProxyHintsInProgressGetOK handles this case with default header values.

StorageProxyHintsInProgressGetOK storage proxy hints in progress get o k
*/
type StorageProxyHintsInProgressGetOK struct {
	Payload int32
}

func (o *StorageProxyHintsInProgressGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageProxyHintsInProgressGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyHintsInProgressGetDefault creates a StorageProxyHintsInProgressGetDefault with default headers values
func NewStorageProxyHintsInProgressGetDefault(code int) *StorageProxyHintsInProgressGetDefault {
	return &StorageProxyHintsInProgressGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyHintsInProgressGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyHintsInProgressGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy hints in progress get default response
func (o *StorageProxyHintsInProgressGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyHintsInProgressGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyHintsInProgressGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyHintsInProgressGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
