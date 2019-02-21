// Code generated by go-swagger; DO NOT EDIT.

package version_incremental

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "api/models"
)

// IncrementalUpdateReader is a Reader for the IncrementalUpdate structure.
type IncrementalUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IncrementalUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIncrementalUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewIncrementalUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIncrementalUpdateOK creates a IncrementalUpdateOK with default headers values
func NewIncrementalUpdateOK() *IncrementalUpdateOK {
	return &IncrementalUpdateOK{}
}

/*IncrementalUpdateOK handles this case with default header values.

Update incremental version response
*/
type IncrementalUpdateOK struct {
	Payload *models.IncrementalVersionNumber
}

func (o *IncrementalUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /increamental_version/{projectName}][%d] incrementalUpdateOK  %+v", 200, o.Payload)
}

func (o *IncrementalUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncrementalVersionNumber)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIncrementalUpdateInternalServerError creates a IncrementalUpdateInternalServerError with default headers values
func NewIncrementalUpdateInternalServerError() *IncrementalUpdateInternalServerError {
	return &IncrementalUpdateInternalServerError{}
}

/*IncrementalUpdateInternalServerError handles this case with default header values.

Error response
*/
type IncrementalUpdateInternalServerError struct {
	XErrorCode string

	Payload *models.Error
}

func (o *IncrementalUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /increamental_version/{projectName}][%d] incrementalUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *IncrementalUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Error-Code
	o.XErrorCode = response.GetHeader("X-Error-Code")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}