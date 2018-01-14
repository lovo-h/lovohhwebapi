package infrastructure

import (
	"net/http"
)

type WebResponder struct{}

func (responder *WebResponder) BadRequest(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)
}

func (responder *WebResponder) Created(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusCreated)
}

func (responder *WebResponder) Forbidden(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusForbidden)
}

func (responder *WebResponder) InternalServerError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
}

func (responder *WebResponder) NoContent(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNoContent)
}

func (responder *WebResponder) NotFound(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNotFound)
}

func (responder *WebResponder) ServiceUnavailable(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusServiceUnavailable)
}

func (responder *WebResponder) Success(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusOK)
}

func (responder *WebResponder) Unauthorized(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusUnauthorized)
}
