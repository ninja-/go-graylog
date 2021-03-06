// Code generated by go-swagger; DO NOT EDIT.

package search_saved

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// FieldHistogramRelativeReader is a Reader for the FieldHistogramRelative structure.
type FieldHistogramRelativeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FieldHistogramRelativeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFieldHistogramRelativeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewFieldHistogramRelativeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFieldHistogramRelativeOK creates a FieldHistogramRelativeOK with default headers values
func NewFieldHistogramRelativeOK() *FieldHistogramRelativeOK {
	return &FieldHistogramRelativeOK{}
}

/*FieldHistogramRelativeOK handles this case with default header values.

No response was specified
*/
type FieldHistogramRelativeOK struct {
	Payload *models.HistogramResult
}

func (o *FieldHistogramRelativeOK) Error() string {
	return fmt.Sprintf("[GET /search/universal/relative/fieldhistogram][%d] fieldHistogramRelativeOK  %+v", 200, o.Payload)
}

func (o *FieldHistogramRelativeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HistogramResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFieldHistogramRelativeBadRequest creates a FieldHistogramRelativeBadRequest with default headers values
func NewFieldHistogramRelativeBadRequest() *FieldHistogramRelativeBadRequest {
	return &FieldHistogramRelativeBadRequest{}
}

/*FieldHistogramRelativeBadRequest handles this case with default header values.

Field is not of numeric type.
*/
type FieldHistogramRelativeBadRequest struct {
}

func (o *FieldHistogramRelativeBadRequest) Error() string {
	return fmt.Sprintf("[GET /search/universal/relative/fieldhistogram][%d] fieldHistogramRelativeBadRequest ", 400)
}

func (o *FieldHistogramRelativeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
