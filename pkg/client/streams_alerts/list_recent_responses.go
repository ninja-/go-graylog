// Code generated by go-swagger; DO NOT EDIT.

package streams_alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// ListRecentReader is a Reader for the ListRecent structure.
type ListRecentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRecentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRecentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListRecentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRecentOK creates a ListRecentOK with default headers values
func NewListRecentOK() *ListRecentOK {
	return &ListRecentOK{}
}

/*ListRecentOK handles this case with default header values.

No response was specified
*/
type ListRecentOK struct {
	Payload *models.AlertListSummary
}

func (o *ListRecentOK) Error() string {
	return fmt.Sprintf("[GET /streams/alerts][%d] listRecentOK  %+v", 200, o.Payload)
}

func (o *ListRecentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertListSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRecentBadRequest creates a ListRecentBadRequest with default headers values
func NewListRecentBadRequest() *ListRecentBadRequest {
	return &ListRecentBadRequest{}
}

/*ListRecentBadRequest handles this case with default header values.

Invalid ObjectId.
*/
type ListRecentBadRequest struct {
}

func (o *ListRecentBadRequest) Error() string {
	return fmt.Sprintf("[GET /streams/alerts][%d] listRecentBadRequest ", 400)
}

func (o *ListRecentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
