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

// GetClusterClusterIDTaskTaskTypeTaskIDHistoryReader is a Reader for the GetClusterClusterIDTaskTaskTypeTaskIDHistory structure.
type GetClusterClusterIDTaskTaskTypeTaskIDHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryOK()
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

// NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryOK creates a GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK with default headers values
func NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryOK() *GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK {
	return &GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK{}
}

/*GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK handles this case with default header values.

List of task runs
*/
type GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK struct {
	Payload []*models.TaskRun
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/{task_type}/{task_id}/history][%d] getClusterClusterIdTaskTaskTypeTaskIdHistoryOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
