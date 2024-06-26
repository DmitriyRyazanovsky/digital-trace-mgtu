// Code generated by go-swagger; DO NOT EDIT.

package attempt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mgtu/digital-trace/main-backend-service/internal/gen/models"
)

// UserAttemptAttemptIDAnswerPutOKCode is the HTTP code returned for type UserAttemptAttemptIDAnswerPutOK
const UserAttemptAttemptIDAnswerPutOKCode int = 200

/*
UserAttemptAttemptIDAnswerPutOK OK

swagger:response userAttemptAttemptIdAnswerPutOK
*/
type UserAttemptAttemptIDAnswerPutOK struct {
}

// NewUserAttemptAttemptIDAnswerPutOK creates UserAttemptAttemptIDAnswerPutOK with default headers values
func NewUserAttemptAttemptIDAnswerPutOK() *UserAttemptAttemptIDAnswerPutOK {

	return &UserAttemptAttemptIDAnswerPutOK{}
}

// WriteResponse to the client
func (o *UserAttemptAttemptIDAnswerPutOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UserAttemptAttemptIDAnswerPutInternalServerErrorCode is the HTTP code returned for type UserAttemptAttemptIDAnswerPutInternalServerError
const UserAttemptAttemptIDAnswerPutInternalServerErrorCode int = 500

/*
UserAttemptAttemptIDAnswerPutInternalServerError Ошибка сервера

swagger:response userAttemptAttemptIdAnswerPutInternalServerError
*/
type UserAttemptAttemptIDAnswerPutInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error500 `json:"body,omitempty"`
}

// NewUserAttemptAttemptIDAnswerPutInternalServerError creates UserAttemptAttemptIDAnswerPutInternalServerError with default headers values
func NewUserAttemptAttemptIDAnswerPutInternalServerError() *UserAttemptAttemptIDAnswerPutInternalServerError {

	return &UserAttemptAttemptIDAnswerPutInternalServerError{}
}

// WithPayload adds the payload to the user attempt attempt Id answer put internal server error response
func (o *UserAttemptAttemptIDAnswerPutInternalServerError) WithPayload(payload *models.Error500) *UserAttemptAttemptIDAnswerPutInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user attempt attempt Id answer put internal server error response
func (o *UserAttemptAttemptIDAnswerPutInternalServerError) SetPayload(payload *models.Error500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserAttemptAttemptIDAnswerPutInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
