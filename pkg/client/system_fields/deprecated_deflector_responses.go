// Code generated by go-swagger; DO NOT EDIT.

package system_fields

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// DeprecatedDeflectorReader is a Reader for the DeprecatedDeflector structure.
type DeprecatedDeflectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeprecatedDeflectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeprecatedDeflectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeprecatedDeflectorOK creates a DeprecatedDeflectorOK with default headers values
func NewDeprecatedDeflectorOK() *DeprecatedDeflectorOK {
	return &DeprecatedDeflectorOK{}
}

/*DeprecatedDeflectorOK handles this case with default header values.

No response was specified
*/
type DeprecatedDeflectorOK struct {
	Payload *models.DeflectorSummary
}

func (o *DeprecatedDeflectorOK) Error() string {
	return fmt.Sprintf("[GET /system/deflector][%d] deprecatedDeflectorOK  %+v", 200, o.Payload)
}

func (o *DeprecatedDeflectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeflectorSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
