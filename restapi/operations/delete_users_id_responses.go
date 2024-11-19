// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"main/models"
)

// DeleteUsersIDNoContentCode is the HTTP code returned for type DeleteUsersIDNoContent
const DeleteUsersIDNoContentCode int = 204

/*
DeleteUsersIDNoContent User deleted

swagger:response deleteUsersIdNoContent
*/
type DeleteUsersIDNoContent struct {
}

// NewDeleteUsersIDNoContent creates DeleteUsersIDNoContent with default headers values
func NewDeleteUsersIDNoContent() *DeleteUsersIDNoContent {

	return &DeleteUsersIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteUsersIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*
DeleteUsersIDDefault Общая ошибка

swagger:response deleteUsersIdDefault
*/
type DeleteUsersIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteUsersIDDefault creates DeleteUsersIDDefault with default headers values
func NewDeleteUsersIDDefault(code int) *DeleteUsersIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteUsersIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete users ID default response
func (o *DeleteUsersIDDefault) WithStatusCode(code int) *DeleteUsersIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete users ID default response
func (o *DeleteUsersIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete users ID default response
func (o *DeleteUsersIDDefault) WithPayload(payload *models.ErrorResponse) *DeleteUsersIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete users ID default response
func (o *DeleteUsersIDDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUsersIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
