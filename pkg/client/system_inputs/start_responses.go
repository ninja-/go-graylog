// Code generated by go-swagger; DO NOT EDIT.

package system_inputs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// StartReader is a Reader for the Start structure.
type StartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewStartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStartOK creates a StartOK with default headers values
func NewStartOK() *StartOK {
	return &StartOK{}
}

/*StartOK handles this case with default header values.

No response was specified
*/
type StartOK struct {
	Payload *models.InputCreated
}

func (o *StartOK) Error() string {
	return fmt.Sprintf("[PUT /system/inputstates/{inputId}][%d] startOK  %+v", 200, o.Payload)
}

func (o *StartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InputCreated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartNotFound creates a StartNotFound with default headers values
func NewStartNotFound() *StartNotFound {
	return &StartNotFound{}
}

/*StartNotFound handles this case with default header values.

No such input on this node.
*/
type StartNotFound struct {
}

func (o *StartNotFound) Error() string {
	return fmt.Sprintf("[PUT /system/inputstates/{inputId}][%d] startNotFound ", 404)
}

func (o *StartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
