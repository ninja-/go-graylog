// Code generated by go-swagger; DO NOT EDIT.

package cluster_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// CancelJobReader is a Reader for the CancelJob structure.
type CancelJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCancelJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelJobOK creates a CancelJobOK with default headers values
func NewCancelJobOK() *CancelJobOK {
	return &CancelJobOK{}
}

/*CancelJobOK handles this case with default header values.

No response was specified
*/
type CancelJobOK struct {
	Payload *models.SystemJobSummary
}

func (o *CancelJobOK) Error() string {
	return fmt.Sprintf("[DELETE /cluster/jobs/{jobId}][%d] cancelJobOK  %+v", 200, o.Payload)
}

func (o *CancelJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemJobSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
