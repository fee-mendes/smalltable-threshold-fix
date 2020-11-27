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

// ColumnFamilyMetricsMemtableColumnsCountGetReader is a Reader for the ColumnFamilyMetricsMemtableColumnsCountGet structure.
type ColumnFamilyMetricsMemtableColumnsCountGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableColumnsCountGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMemtableColumnsCountGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMemtableColumnsCountGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMemtableColumnsCountGetOK creates a ColumnFamilyMetricsMemtableColumnsCountGetOK with default headers values
func NewColumnFamilyMetricsMemtableColumnsCountGetOK() *ColumnFamilyMetricsMemtableColumnsCountGetOK {
	return &ColumnFamilyMetricsMemtableColumnsCountGetOK{}
}

/*ColumnFamilyMetricsMemtableColumnsCountGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableColumnsCountGetOK column family metrics memtable columns count get o k
*/
type ColumnFamilyMetricsMemtableColumnsCountGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMemtableColumnsCountGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableColumnsCountGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMemtableColumnsCountGetDefault creates a ColumnFamilyMetricsMemtableColumnsCountGetDefault with default headers values
func NewColumnFamilyMetricsMemtableColumnsCountGetDefault(code int) *ColumnFamilyMetricsMemtableColumnsCountGetDefault {
	return &ColumnFamilyMetricsMemtableColumnsCountGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsMemtableColumnsCountGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMemtableColumnsCountGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics memtable columns count get default response
func (o *ColumnFamilyMetricsMemtableColumnsCountGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMemtableColumnsCountGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableColumnsCountGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMemtableColumnsCountGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
