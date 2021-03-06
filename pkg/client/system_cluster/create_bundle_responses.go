// Code generated by go-swagger; DO NOT EDIT.

package system_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateBundleReader is a Reader for the CreateBundle structure.
type CreateBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateBundleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateBundleOK creates a CreateBundleOK with default headers values
func NewCreateBundleOK() *CreateBundleOK {
	return &CreateBundleOK{}
}

/*CreateBundleOK handles this case with default header values.

No response was specified
*/
type CreateBundleOK struct {
}

func (o *CreateBundleOK) Error() string {
	return fmt.Sprintf("[POST /system/bundles][%d] createBundleOK ", 200)
}

func (o *CreateBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBundleBadRequest creates a CreateBundleBadRequest with default headers values
func NewCreateBundleBadRequest() *CreateBundleBadRequest {
	return &CreateBundleBadRequest{}
}

/*CreateBundleBadRequest handles this case with default header values.

Missing or invalid content pack
*/
type CreateBundleBadRequest struct {
}

func (o *CreateBundleBadRequest) Error() string {
	return fmt.Sprintf("[POST /system/bundles][%d] createBundleBadRequest ", 400)
}

func (o *CreateBundleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBundleInternalServerError creates a CreateBundleInternalServerError with default headers values
func NewCreateBundleInternalServerError() *CreateBundleInternalServerError {
	return &CreateBundleInternalServerError{}
}

/*CreateBundleInternalServerError handles this case with default header values.

Error while saving content pack
*/
type CreateBundleInternalServerError struct {
}

func (o *CreateBundleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /system/bundles][%d] createBundleInternalServerError ", 500)
}

func (o *CreateBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
