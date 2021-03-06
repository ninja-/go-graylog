// Code generated by go-swagger; DO NOT EDIT.

package system_messageprocessors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// AdaptersReader is a Reader for the Adapters structure.
type AdaptersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdaptersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAdaptersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAdaptersOK creates a AdaptersOK with default headers values
func NewAdaptersOK() *AdaptersOK {
	return &AdaptersOK{}
}

/*AdaptersOK handles this case with default header values.

No response was specified
*/
type AdaptersOK struct {
	Payload *models.DataAdapterPage
}

func (o *AdaptersOK) Error() string {
	return fmt.Sprintf("[GET /system/lookup/adapters][%d] adaptersOK  %+v", 200, o.Payload)
}

func (o *AdaptersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataAdapterPage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
