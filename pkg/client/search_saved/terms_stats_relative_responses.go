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

// TermsStatsRelativeReader is a Reader for the TermsStatsRelative structure.
type TermsStatsRelativeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TermsStatsRelativeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTermsStatsRelativeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTermsStatsRelativeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTermsStatsRelativeOK creates a TermsStatsRelativeOK with default headers values
func NewTermsStatsRelativeOK() *TermsStatsRelativeOK {
	return &TermsStatsRelativeOK{}
}

/*TermsStatsRelativeOK handles this case with default header values.

No response was specified
*/
type TermsStatsRelativeOK struct {
	Payload *models.TermsStatsResult
}

func (o *TermsStatsRelativeOK) Error() string {
	return fmt.Sprintf("[GET /search/universal/relative/termsstats][%d] termsStatsRelativeOK  %+v", 200, o.Payload)
}

func (o *TermsStatsRelativeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TermsStatsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTermsStatsRelativeBadRequest creates a TermsStatsRelativeBadRequest with default headers values
func NewTermsStatsRelativeBadRequest() *TermsStatsRelativeBadRequest {
	return &TermsStatsRelativeBadRequest{}
}

/*TermsStatsRelativeBadRequest handles this case with default header values.

Invalid timerange parameters provided.
*/
type TermsStatsRelativeBadRequest struct {
}

func (o *TermsStatsRelativeBadRequest) Error() string {
	return fmt.Sprintf("[GET /search/universal/relative/termsstats][%d] termsStatsRelativeBadRequest ", 400)
}

func (o *TermsStatsRelativeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
