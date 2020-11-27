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

// ColumnFamilyMetricsBloomFilterFalseRatioByNameGetReader is a Reader for the ColumnFamilyMetricsBloomFilterFalseRatioByNameGet structure.
type ColumnFamilyMetricsBloomFilterFalseRatioByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK creates a ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK with default headers values
func NewColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK() *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK {
	return &ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK{}
}

/*ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK handles this case with default header values.

ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK column family metrics bloom filter false ratio by name get o k
*/
type ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault creates a ColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault with default headers values
func NewColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault(code int) *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault {
	return &ColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics bloom filter false ratio by name get default response
func (o *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
