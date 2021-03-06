// Code generated by go-swagger; DO NOT EDIT.

package search_saved

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SearchRelativeChunkedReader is a Reader for the SearchRelativeChunked structure.
type SearchRelativeChunkedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchRelativeChunkedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchRelativeChunkedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSearchRelativeChunkedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchRelativeChunkedOK creates a SearchRelativeChunkedOK with default headers values
func NewSearchRelativeChunkedOK() *SearchRelativeChunkedOK {
	return &SearchRelativeChunkedOK{}
}

/*SearchRelativeChunkedOK handles this case with default header values.

No response was specified
*/
type SearchRelativeChunkedOK struct {
	Payload SearchRelativeChunkedOKBody
}

func (o *SearchRelativeChunkedOK) Error() string {
	return fmt.Sprintf("[GET /search/universal/relative][%d] searchRelativeChunkedOK  %+v", 200, o.Payload)
}

func (o *SearchRelativeChunkedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRelativeChunkedBadRequest creates a SearchRelativeChunkedBadRequest with default headers values
func NewSearchRelativeChunkedBadRequest() *SearchRelativeChunkedBadRequest {
	return &SearchRelativeChunkedBadRequest{}
}

/*SearchRelativeChunkedBadRequest handles this case with default header values.

Invalid timerange parameters provided.
*/
type SearchRelativeChunkedBadRequest struct {
}

func (o *SearchRelativeChunkedBadRequest) Error() string {
	return fmt.Sprintf("[GET /search/universal/relative][%d] searchRelativeChunkedBadRequest ", 400)
}

func (o *SearchRelativeChunkedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*SearchRelativeChunkedOKBody search relative chunked o k body
swagger:model SearchRelativeChunkedOKBody
*/
type SearchRelativeChunkedOKBody interface{}
