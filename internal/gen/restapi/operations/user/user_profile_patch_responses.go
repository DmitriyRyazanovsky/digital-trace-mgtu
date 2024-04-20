// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mgtu/digital-trace/main-backend-service/internal/gen/models"
)

// UserProfilePatchOKCode is the HTTP code returned for type UserProfilePatchOK
const UserProfilePatchOKCode int = 200

/*
UserProfilePatchOK OK

swagger:response userProfilePatchOK
*/
type UserProfilePatchOK struct {
}

// NewUserProfilePatchOK creates UserProfilePatchOK with default headers values
func NewUserProfilePatchOK() *UserProfilePatchOK {

	return &UserProfilePatchOK{}
}

// WriteResponse to the client
func (o *UserProfilePatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UserProfilePatchInternalServerErrorCode is the HTTP code returned for type UserProfilePatchInternalServerError
const UserProfilePatchInternalServerErrorCode int = 500

/*
UserProfilePatchInternalServerError Ошибка сервера

swagger:response userProfilePatchInternalServerError
*/
type UserProfilePatchInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error500 `json:"body,omitempty"`
}

// NewUserProfilePatchInternalServerError creates UserProfilePatchInternalServerError with default headers values
func NewUserProfilePatchInternalServerError() *UserProfilePatchInternalServerError {

	return &UserProfilePatchInternalServerError{}
}

// WithPayload adds the payload to the user profile patch internal server error response
func (o *UserProfilePatchInternalServerError) WithPayload(payload *models.Error500) *UserProfilePatchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user profile patch internal server error response
func (o *UserProfilePatchInternalServerError) SetPayload(payload *models.Error500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserProfilePatchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
