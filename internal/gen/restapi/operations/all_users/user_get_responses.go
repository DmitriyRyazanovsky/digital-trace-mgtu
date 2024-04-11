// Code generated by go-swagger; DO NOT EDIT.

package all_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mgtu/digital-trace/main-backend-service/internal/gen/models"
)

// UserGetOKCode is the HTTP code returned for type UserGetOK
const UserGetOKCode int = 200

/*
UserGetOK Информация о пользователях успешно получена

swagger:response userGetOK
*/
type UserGetOK struct {

	/*
	  In: Body
	*/
	Payload []*UserGetOKBodyItems0 `json:"body,omitempty"`
}

// NewUserGetOK creates UserGetOK with default headers values
func NewUserGetOK() *UserGetOK {

	return &UserGetOK{}
}

// WithPayload adds the payload to the user get o k response
func (o *UserGetOK) WithPayload(payload []*UserGetOKBodyItems0) *UserGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user get o k response
func (o *UserGetOK) SetPayload(payload []*UserGetOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*UserGetOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UserGetInternalServerErrorCode is the HTTP code returned for type UserGetInternalServerError
const UserGetInternalServerErrorCode int = 500

/*
UserGetInternalServerError Ошибка сервера

swagger:response userGetInternalServerError
*/
type UserGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error500 `json:"body,omitempty"`
}

// NewUserGetInternalServerError creates UserGetInternalServerError with default headers values
func NewUserGetInternalServerError() *UserGetInternalServerError {

	return &UserGetInternalServerError{}
}

// WithPayload adds the payload to the user get internal server error response
func (o *UserGetInternalServerError) WithPayload(payload *models.Error500) *UserGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user get internal server error response
func (o *UserGetInternalServerError) SetPayload(payload *models.Error500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}