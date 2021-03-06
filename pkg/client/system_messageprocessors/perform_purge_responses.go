// Code generated by go-swagger; DO NOT EDIT.

package system_messageprocessors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PerformPurgeReader is a Reader for the PerformPurge structure.
type PerformPurgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformPurgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPerformPurgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPerformPurgeOK creates a PerformPurgeOK with default headers values
func NewPerformPurgeOK() *PerformPurgeOK {
	return &PerformPurgeOK{}
}

/*PerformPurgeOK handles this case with default header values.

No response was specified
*/
type PerformPurgeOK struct {
}

func (o *PerformPurgeOK) Error() string {
	return fmt.Sprintf("[POST /system/lookup/tables/{idOrName}/purge][%d] performPurgeOK ", 200)
}

func (o *PerformPurgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
