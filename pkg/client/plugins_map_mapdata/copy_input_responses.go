// Code generated by go-swagger; DO NOT EDIT.

package plugins_map_mapdata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CopyInputReader is a Reader for the CopyInput structure.
type CopyInputReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopyInputReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCopyInputOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCopyInputBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCopyInputNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCopyInputOK creates a CopyInputOK with default headers values
func NewCopyInputOK() *CopyInputOK {
	return &CopyInputOK{}
}

/*CopyInputOK handles this case with default header values.

No response was specified
*/
type CopyInputOK struct {
}

func (o *CopyInputOK) Error() string {
	return fmt.Sprintf("[POST /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{inputId}/{name}][%d] copyInputOK ", 200)
}

func (o *CopyInputOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyInputBadRequest creates a CopyInputBadRequest with default headers values
func NewCopyInputBadRequest() *CopyInputBadRequest {
	return &CopyInputBadRequest{}
}

/*CopyInputBadRequest handles this case with default header values.

Invalid ObjectId.
*/
type CopyInputBadRequest struct {
}

func (o *CopyInputBadRequest) Error() string {
	return fmt.Sprintf("[POST /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{inputId}/{name}][%d] copyInputBadRequest ", 400)
}

func (o *CopyInputBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyInputNotFound creates a CopyInputNotFound with default headers values
func NewCopyInputNotFound() *CopyInputNotFound {
	return &CopyInputNotFound{}
}

/*CopyInputNotFound handles this case with default header values.

Configuration or Input not found.
*/
type CopyInputNotFound struct {
}

func (o *CopyInputNotFound) Error() string {
	return fmt.Sprintf("[POST /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{inputId}/{name}][%d] copyInputNotFound ", 404)
}

func (o *CopyInputNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
