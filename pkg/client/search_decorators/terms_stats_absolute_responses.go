// Code generated by go-swagger; DO NOT EDIT.

package search_decorators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// TermsStatsAbsoluteReader is a Reader for the TermsStatsAbsolute structure.
type TermsStatsAbsoluteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TermsStatsAbsoluteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTermsStatsAbsoluteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTermsStatsAbsoluteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTermsStatsAbsoluteOK creates a TermsStatsAbsoluteOK with default headers values
func NewTermsStatsAbsoluteOK() *TermsStatsAbsoluteOK {
	return &TermsStatsAbsoluteOK{}
}

/*TermsStatsAbsoluteOK handles this case with default header values.

No response was specified
*/
type TermsStatsAbsoluteOK struct {
	Payload *models.TermsStatsResult
}

func (o *TermsStatsAbsoluteOK) Error() string {
	return fmt.Sprintf("[GET /search/universal/absolute/termsstats][%d] termsStatsAbsoluteOK  %+v", 200, o.Payload)
}

func (o *TermsStatsAbsoluteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TermsStatsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTermsStatsAbsoluteBadRequest creates a TermsStatsAbsoluteBadRequest with default headers values
func NewTermsStatsAbsoluteBadRequest() *TermsStatsAbsoluteBadRequest {
	return &TermsStatsAbsoluteBadRequest{}
}

/*TermsStatsAbsoluteBadRequest handles this case with default header values.

Invalid timerange parameters provided.
*/
type TermsStatsAbsoluteBadRequest struct {
}

func (o *TermsStatsAbsoluteBadRequest) Error() string {
	return fmt.Sprintf("[GET /search/universal/absolute/termsstats][%d] termsStatsAbsoluteBadRequest ", 400)
}

func (o *TermsStatsAbsoluteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
