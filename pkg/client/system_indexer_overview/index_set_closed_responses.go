// Code generated by go-swagger; DO NOT EDIT.

package system_indexer_overview

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// IndexSetClosedReader is a Reader for the IndexSetClosed structure.
type IndexSetClosedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexSetClosedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexSetClosedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexSetClosedOK creates a IndexSetClosedOK with default headers values
func NewIndexSetClosedOK() *IndexSetClosedOK {
	return &IndexSetClosedOK{}
}

/*IndexSetClosedOK handles this case with default header values.

No response was specified
*/
type IndexSetClosedOK struct {
	Payload *models.ClosedIndices
}

func (o *IndexSetClosedOK) Error() string {
	return fmt.Sprintf("[GET /system/indexer/indices/{indexSetId}/closed][%d] indexSetClosedOK  %+v", 200, o.Payload)
}

func (o *IndexSetClosedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClosedIndices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
