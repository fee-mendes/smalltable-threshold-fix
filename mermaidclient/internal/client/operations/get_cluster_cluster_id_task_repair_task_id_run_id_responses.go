// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/mermaidclient/internal/models"
)

// GetClusterClusterIDTaskRepairTaskIDRunIDReader is a Reader for the GetClusterClusterIDTaskRepairTaskIDRunID structure.
type GetClusterClusterIDTaskRepairTaskIDRunIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDTaskRepairTaskIDRunIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterClusterIDTaskRepairTaskIDRunIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		body := response.Body()
		defer body.Close()

		var m json.RawMessage
		if err := json.NewDecoder(body).Decode(&m); err != nil {
			return nil, err
		}

		return nil, runtime.NewAPIError("API error", m, response.Code())
	}
}

// NewGetClusterClusterIDTaskRepairTaskIDRunIDOK creates a GetClusterClusterIDTaskRepairTaskIDRunIDOK with default headers values
func NewGetClusterClusterIDTaskRepairTaskIDRunIDOK() *GetClusterClusterIDTaskRepairTaskIDRunIDOK {
	return &GetClusterClusterIDTaskRepairTaskIDRunIDOK{}
}

/*GetClusterClusterIDTaskRepairTaskIDRunIDOK handles this case with default header values.

Repair progress
*/
type GetClusterClusterIDTaskRepairTaskIDRunIDOK struct {
	Payload *models.TaskRunRepairProgress
}

func (o *GetClusterClusterIDTaskRepairTaskIDRunIDOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/repair/{task_id}/{run_id}][%d] getClusterClusterIdTaskRepairTaskIdRunIdOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDTaskRepairTaskIDRunIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskRunRepairProgress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
