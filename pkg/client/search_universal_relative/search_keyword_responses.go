// Code generated by go-swagger; DO NOT EDIT.

package search_universal_relative

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// SearchKeywordReader is a Reader for the SearchKeyword structure.
type SearchKeywordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchKeywordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchKeywordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSearchKeywordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchKeywordOK creates a SearchKeywordOK with default headers values
func NewSearchKeywordOK() *SearchKeywordOK {
	return &SearchKeywordOK{}
}

/*SearchKeywordOK handles this case with default header values.

No response was specified
*/
type SearchKeywordOK struct {
	Payload *models.SearchResponse
}

func (o *SearchKeywordOK) Error() string {
	return fmt.Sprintf("[GET /search/universal/keyword][%d] searchKeywordOK  %+v", 200, o.Payload)
}

func (o *SearchKeywordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchKeywordBadRequest creates a SearchKeywordBadRequest with default headers values
func NewSearchKeywordBadRequest() *SearchKeywordBadRequest {
	return &SearchKeywordBadRequest{}
}

/*SearchKeywordBadRequest handles this case with default header values.

Invalid keyword provided.
*/
type SearchKeywordBadRequest struct {
}

func (o *SearchKeywordBadRequest) Error() string {
	return fmt.Sprintf("[GET /search/universal/keyword][%d] searchKeywordBadRequest ", 400)
}

func (o *SearchKeywordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
