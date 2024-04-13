// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mgtu/digital-trace/main-backend-service/internal/gen/models"
)

// UserProfileGetOKCode is the HTTP code returned for type UserProfileGetOK
const UserProfileGetOKCode int = 200

/*
UserProfileGetOK OK

swagger:response userProfileGetOK
*/
type UserProfileGetOK struct {

	/*
	  In: Body
	*/
	Payload *UserProfileGetOKBody `json:"body,omitempty"`
}

// NewUserProfileGetOK creates UserProfileGetOK with default headers values
func NewUserProfileGetOK() *UserProfileGetOK {

	return &UserProfileGetOK{}
}

// WithPayload adds the payload to the user profile get o k response
func (o *UserProfileGetOK) WithPayload(payload *UserProfileGetOKBody) *UserProfileGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user profile get o k response
func (o *UserProfileGetOK) SetPayload(payload *UserProfileGetOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserProfileGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserProfileGetInternalServerErrorCode is the HTTP code returned for type UserProfileGetInternalServerError
const UserProfileGetInternalServerErrorCode int = 500

/*
UserProfileGetInternalServerError Ошибка сервера

swagger:response userProfileGetInternalServerError
*/
type UserProfileGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error500 `json:"body,omitempty"`
}

// NewUserProfileGetInternalServerError creates UserProfileGetInternalServerError with default headers values
func NewUserProfileGetInternalServerError() *UserProfileGetInternalServerError {

	return &UserProfileGetInternalServerError{}
}

// WithPayload adds the payload to the user profile get internal server error response
func (o *UserProfileGetInternalServerError) WithPayload(payload *models.Error500) *UserProfileGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user profile get internal server error response
func (o *UserProfileGetInternalServerError) SetPayload(payload *models.Error500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserProfileGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}