// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigCompactionLargeCellWarningThresholdMbReader is a Reader for the FindConfigCompactionLargeCellWarningThresholdMb structure.
type FindConfigCompactionLargeCellWarningThresholdMbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCompactionLargeCellWarningThresholdMbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCompactionLargeCellWarningThresholdMbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCompactionLargeCellWarningThresholdMbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCompactionLargeCellWarningThresholdMbOK creates a FindConfigCompactionLargeCellWarningThresholdMbOK with default headers values
func NewFindConfigCompactionLargeCellWarningThresholdMbOK() *FindConfigCompactionLargeCellWarningThresholdMbOK {
	return &FindConfigCompactionLargeCellWarningThresholdMbOK{}
}

/*FindConfigCompactionLargeCellWarningThresholdMbOK handles this case with default header values.

Config value
*/
type FindConfigCompactionLargeCellWarningThresholdMbOK struct {
	Payload int64
}

func (o *FindConfigCompactionLargeCellWarningThresholdMbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigCompactionLargeCellWarningThresholdMbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCompactionLargeCellWarningThresholdMbDefault creates a FindConfigCompactionLargeCellWarningThresholdMbDefault with default headers values
func NewFindConfigCompactionLargeCellWarningThresholdMbDefault(code int) *FindConfigCompactionLargeCellWarningThresholdMbDefault {
	return &FindConfigCompactionLargeCellWarningThresholdMbDefault{
		_statusCode: code,
	}
}

/*FindConfigCompactionLargeCellWarningThresholdMbDefault handles this case with default header values.

unexpected error
*/
type FindConfigCompactionLargeCellWarningThresholdMbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config compaction large cell warning threshold mb default response
func (o *FindConfigCompactionLargeCellWarningThresholdMbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCompactionLargeCellWarningThresholdMbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCompactionLargeCellWarningThresholdMbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCompactionLargeCellWarningThresholdMbDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
