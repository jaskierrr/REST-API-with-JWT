// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"main/internal/models"
)

// PatchUsersIDBalanceCreatedCode is the HTTP code returned for type PatchUsersIDBalanceCreated
const PatchUsersIDBalanceCreatedCode int = 201

/*
PatchUsersIDBalanceCreated Update balance updated

swagger:response patchUsersIdBalanceCreated
*/
type PatchUsersIDBalanceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPatchUsersIDBalanceCreated creates PatchUsersIDBalanceCreated with default headers values
func NewPatchUsersIDBalanceCreated() *PatchUsersIDBalanceCreated {

	return &PatchUsersIDBalanceCreated{}
}

// WithPayload adds the payload to the patch users Id balance created response
func (o *PatchUsersIDBalanceCreated) WithPayload(payload *models.User) *PatchUsersIDBalanceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch users Id balance created response
func (o *PatchUsersIDBalanceCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUsersIDBalanceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PatchUsersIDBalanceDefault Общая ошибка

swagger:response patchUsersIdBalanceDefault
*/
type PatchUsersIDBalanceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPatchUsersIDBalanceDefault creates PatchUsersIDBalanceDefault with default headers values
func NewPatchUsersIDBalanceDefault(code int) *PatchUsersIDBalanceDefault {
	if code <= 0 {
		code = 500
	}

	return &PatchUsersIDBalanceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the patch users ID balance default response
func (o *PatchUsersIDBalanceDefault) WithStatusCode(code int) *PatchUsersIDBalanceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the patch users ID balance default response
func (o *PatchUsersIDBalanceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the patch users ID balance default response
func (o *PatchUsersIDBalanceDefault) WithPayload(payload *models.ErrorResponse) *PatchUsersIDBalanceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch users ID balance default response
func (o *PatchUsersIDBalanceDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUsersIDBalanceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
