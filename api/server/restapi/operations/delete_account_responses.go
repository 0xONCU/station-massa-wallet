// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/massalabs/station-massa-wallet/api/server/models"
)

// DeleteAccountNoContentCode is the HTTP code returned for type DeleteAccountNoContent
const DeleteAccountNoContentCode int = 204

/*
DeleteAccountNoContent Account deleted successfully.

swagger:response deleteAccountNoContent
*/
type DeleteAccountNoContent struct {

	/*
	  In: Body
	*/
	Payload *models.Account `json:"body,omitempty"`
}

// NewDeleteAccountNoContent creates DeleteAccountNoContent with default headers values
func NewDeleteAccountNoContent() *DeleteAccountNoContent {

	return &DeleteAccountNoContent{}
}

// WithPayload adds the payload to the delete account no content response
func (o *DeleteAccountNoContent) WithPayload(payload *models.Account) *DeleteAccountNoContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete account no content response
func (o *DeleteAccountNoContent) SetPayload(payload *models.Account) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAccountNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAccountBadRequestCode is the HTTP code returned for type DeleteAccountBadRequest
const DeleteAccountBadRequestCode int = 400

/*
DeleteAccountBadRequest Bad request.

swagger:response deleteAccountBadRequest
*/
type DeleteAccountBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAccountBadRequest creates DeleteAccountBadRequest with default headers values
func NewDeleteAccountBadRequest() *DeleteAccountBadRequest {

	return &DeleteAccountBadRequest{}
}

// WithPayload adds the payload to the delete account bad request response
func (o *DeleteAccountBadRequest) WithPayload(payload *models.Error) *DeleteAccountBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete account bad request response
func (o *DeleteAccountBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAccountBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAccountUnauthorizedCode is the HTTP code returned for type DeleteAccountUnauthorized
const DeleteAccountUnauthorizedCode int = 401

/*
DeleteAccountUnauthorized Unauthorized - The request requires user authentication.

swagger:response deleteAccountUnauthorized
*/
type DeleteAccountUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAccountUnauthorized creates DeleteAccountUnauthorized with default headers values
func NewDeleteAccountUnauthorized() *DeleteAccountUnauthorized {

	return &DeleteAccountUnauthorized{}
}

// WithPayload adds the payload to the delete account unauthorized response
func (o *DeleteAccountUnauthorized) WithPayload(payload *models.Error) *DeleteAccountUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete account unauthorized response
func (o *DeleteAccountUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAccountUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAccountNotFoundCode is the HTTP code returned for type DeleteAccountNotFound
const DeleteAccountNotFoundCode int = 404

/*
DeleteAccountNotFound Not found.

swagger:response deleteAccountNotFound
*/
type DeleteAccountNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAccountNotFound creates DeleteAccountNotFound with default headers values
func NewDeleteAccountNotFound() *DeleteAccountNotFound {

	return &DeleteAccountNotFound{}
}

// WithPayload adds the payload to the delete account not found response
func (o *DeleteAccountNotFound) WithPayload(payload *models.Error) *DeleteAccountNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete account not found response
func (o *DeleteAccountNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAccountNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAccountInternalServerErrorCode is the HTTP code returned for type DeleteAccountInternalServerError
const DeleteAccountInternalServerErrorCode int = 500

/*
DeleteAccountInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response deleteAccountInternalServerError
*/
type DeleteAccountInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAccountInternalServerError creates DeleteAccountInternalServerError with default headers values
func NewDeleteAccountInternalServerError() *DeleteAccountInternalServerError {

	return &DeleteAccountInternalServerError{}
}

// WithPayload adds the payload to the delete account internal server error response
func (o *DeleteAccountInternalServerError) WithPayload(payload *models.Error) *DeleteAccountInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete account internal server error response
func (o *DeleteAccountInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAccountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
