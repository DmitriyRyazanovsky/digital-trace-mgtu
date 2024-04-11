// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mgtu/digital-trace/main-backend-service/internal/gen/models"
)

// UserProfileOKCode is the HTTP code returned for type UserProfileOK
const UserProfileOKCode int = 200

/*
UserProfileOK информация о достижениях успешно получена

swagger:response userProfileOK
*/
type UserProfileOK struct {

	/*
	  In: Body
	*/
	Payload *UserProfileOKBody `json:"body,omitempty"`
}

// NewUserProfileOK creates UserProfileOK with default headers values
func NewUserProfileOK() *UserProfileOK {

	return &UserProfileOK{}
}

// WithPayload adds the payload to the user profile o k response
func (o *UserProfileOK) WithPayload(payload *UserProfileOKBody) *UserProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user profile o k response
func (o *UserProfileOK) SetPayload(payload *UserProfileOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserProfileInternalServerErrorCode is the HTTP code returned for type UserProfileInternalServerError
const UserProfileInternalServerErrorCode int = 500

/*
UserProfileInternalServerError Ошибка сервера либо запроса

swagger:response userProfileInternalServerError
*/
type UserProfileInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error500 `json:"body,omitempty"`
}

// NewUserProfileInternalServerError creates UserProfileInternalServerError with default headers values
func NewUserProfileInternalServerError() *UserProfileInternalServerError {

	return &UserProfileInternalServerError{}
}

// WithPayload adds the payload to the user profile internal server error response
func (o *UserProfileInternalServerError) WithPayload(payload *models.Error500) *UserProfileInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user profile internal server error response
func (o *UserProfileInternalServerError) SetPayload(payload *models.Error500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserProfileInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}