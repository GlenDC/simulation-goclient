package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// PutThermalSimulationReader is a Reader for the PutThermalSimulation structure.
type PutThermalSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutThermalSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutThermalSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPutThermalSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutThermalSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutThermalSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutThermalSimulationOK creates a PutThermalSimulationOK with default headers values
func NewPutThermalSimulationOK() *PutThermalSimulationOK {
	return &PutThermalSimulationOK{}
}

/*PutThermalSimulationOK handles this case with default header values.

Successfully updated a simulation
*/
type PutThermalSimulationOK struct {
	Payload *models.ThermalSimulation
}

func (o *PutThermalSimulationOK) Error() string {
	return fmt.Sprintf("[PUT /thermalsimulations/{id}][%d] putThermalSimulationOK  %+v", 200, o.Payload)
}

func (o *PutThermalSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThermalSimulation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutThermalSimulationUnauthorized creates a PutThermalSimulationUnauthorized with default headers values
func NewPutThermalSimulationUnauthorized() *PutThermalSimulationUnauthorized {
	return &PutThermalSimulationUnauthorized{}
}

/*PutThermalSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type PutThermalSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *PutThermalSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /thermalsimulations/{id}][%d] putThermalSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *PutThermalSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutThermalSimulationForbidden creates a PutThermalSimulationForbidden with default headers values
func NewPutThermalSimulationForbidden() *PutThermalSimulationForbidden {
	return &PutThermalSimulationForbidden{}
}

/*PutThermalSimulationForbidden handles this case with default header values.

Forbidden
*/
type PutThermalSimulationForbidden struct {
	Payload *models.Error
}

func (o *PutThermalSimulationForbidden) Error() string {
	return fmt.Sprintf("[PUT /thermalsimulations/{id}][%d] putThermalSimulationForbidden  %+v", 403, o.Payload)
}

func (o *PutThermalSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutThermalSimulationDefault creates a PutThermalSimulationDefault with default headers values
func NewPutThermalSimulationDefault(code int) *PutThermalSimulationDefault {
	return &PutThermalSimulationDefault{
		_statusCode: code,
	}
}

/*PutThermalSimulationDefault handles this case with default header values.

unexpected error
*/
type PutThermalSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put thermal simulation default response
func (o *PutThermalSimulationDefault) Code() int {
	return o._statusCode
}

func (o *PutThermalSimulationDefault) Error() string {
	return fmt.Sprintf("[PUT /thermalsimulations/{id}][%d] putThermalSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *PutThermalSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
