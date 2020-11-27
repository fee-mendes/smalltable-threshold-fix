// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/pkg/managerclient/internal/models"
)

// PostClustersReader is a Reader for the PostClusters structure.
type PostClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostClustersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostClustersCreated creates a PostClustersCreated with default headers values
func NewPostClustersCreated() *PostClustersCreated {
	return &PostClustersCreated{}
}

/*PostClustersCreated handles this case with default header values.

Cluster added
*/
type PostClustersCreated struct {
	Location string
}

func (o *PostClustersCreated) Error() string {
	return fmt.Sprintf("[POST /clusters][%d] postClustersCreated ", 201)
}

func (o *PostClustersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewPostClustersDefault creates a PostClustersDefault with default headers values
func NewPostClustersDefault(code int) *PostClustersDefault {
	return &PostClustersDefault{
		_statusCode: code,
	}
}

/*PostClustersDefault handles this case with default header values.

Error
*/
type PostClustersDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the post clusters default response
func (o *PostClustersDefault) Code() int {
	return o._statusCode
}

func (o *PostClustersDefault) Error() string {
	return fmt.Sprintf("[POST /clusters][%d] PostClusters default  %+v", o._statusCode, o.Payload)
}

func (o *PostClustersDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
