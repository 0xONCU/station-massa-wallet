// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra-plugin-wallet/api/server/models"
)

// RestAccountGetOKCode is the HTTP code returned for type RestAccountGetOK
const RestAccountGetOKCode int = 200

/*
RestAccountGetOK Account retrieved.

swagger:response restAccountGetOK
*/
type RestAccountGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Account `json:"body,omitempty"`
}

// NewRestAccountGetOK creates RestAccountGetOK with default headers values
func NewRestAccountGetOK() *RestAccountGetOK {

	return &RestAccountGetOK{}
}

// WithPayload adds the payload to the rest account get o k response
func (o *RestAccountGetOK) WithPayload(payload *models.Account) *RestAccountGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest account get o k response
func (o *RestAccountGetOK) SetPayload(payload *models.Account) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestAccountGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestAccountGetBadRequestCode is the HTTP code returned for type RestAccountGetBadRequest
const RestAccountGetBadRequestCode int = 400

/*
RestAccountGetBadRequest Bad request.

swagger:response restAccountGetBadRequest
*/
type RestAccountGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestAccountGetBadRequest creates RestAccountGetBadRequest with default headers values
func NewRestAccountGetBadRequest() *RestAccountGetBadRequest {

	return &RestAccountGetBadRequest{}
}

// WithPayload adds the payload to the rest account get bad request response
func (o *RestAccountGetBadRequest) WithPayload(payload *models.Error) *RestAccountGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest account get bad request response
func (o *RestAccountGetBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestAccountGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestAccountGetUnauthorizedCode is the HTTP code returned for type RestAccountGetUnauthorized
const RestAccountGetUnauthorizedCode int = 401

/*
RestAccountGetUnauthorized Unauthorized - The request requires user authentication. Only possible if ciphered is false.

swagger:response restAccountGetUnauthorized
*/
type RestAccountGetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestAccountGetUnauthorized creates RestAccountGetUnauthorized with default headers values
func NewRestAccountGetUnauthorized() *RestAccountGetUnauthorized {

	return &RestAccountGetUnauthorized{}
}

// WithPayload adds the payload to the rest account get unauthorized response
func (o *RestAccountGetUnauthorized) WithPayload(payload *models.Error) *RestAccountGetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest account get unauthorized response
func (o *RestAccountGetUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestAccountGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestAccountGetNotFoundCode is the HTTP code returned for type RestAccountGetNotFound
const RestAccountGetNotFoundCode int = 404

/*
RestAccountGetNotFound Not found.

swagger:response restAccountGetNotFound
*/
type RestAccountGetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestAccountGetNotFound creates RestAccountGetNotFound with default headers values
func NewRestAccountGetNotFound() *RestAccountGetNotFound {

	return &RestAccountGetNotFound{}
}

// WithPayload adds the payload to the rest account get not found response
func (o *RestAccountGetNotFound) WithPayload(payload *models.Error) *RestAccountGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest account get not found response
func (o *RestAccountGetNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestAccountGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestAccountGetInternalServerErrorCode is the HTTP code returned for type RestAccountGetInternalServerError
const RestAccountGetInternalServerErrorCode int = 500

/*
RestAccountGetInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response restAccountGetInternalServerError
*/
type RestAccountGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestAccountGetInternalServerError creates RestAccountGetInternalServerError with default headers values
func NewRestAccountGetInternalServerError() *RestAccountGetInternalServerError {

	return &RestAccountGetInternalServerError{}
}

// WithPayload adds the payload to the rest account get internal server error response
func (o *RestAccountGetInternalServerError) WithPayload(payload *models.Error) *RestAccountGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest account get internal server error response
func (o *RestAccountGetInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestAccountGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}