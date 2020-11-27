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

// ColumnFamilyMetricsTombstoneScannedHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsTombstoneScannedHistogramByNameGet structure.
type ColumnFamilyMetricsTombstoneScannedHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsTombstoneScannedHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK creates a ColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK() *ColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK {
	return &ColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK column family metrics tombstone scanned histogram by name get o k
*/
type ColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsTombstoneScannedHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault creates a ColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault with default headers values
func NewColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault(code int) *ColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault {
	return &ColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics tombstone scanned histogram by name get default response
func (o *ColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsTombstoneScannedHistogramByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
