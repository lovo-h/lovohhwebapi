package interfaces

import (
	"github.com/lovohh/lovohhwebapi/usecases"
	"net/http"
)

type WebResponder interface {
	BadRequest(rw http.ResponseWriter)
	Created(rw http.ResponseWriter)
	Forbidden(rw http.ResponseWriter)
	InternalServerError(rw http.ResponseWriter)
	NoContent(rw http.ResponseWriter)
	NotFound(rw http.ResponseWriter)
	ServiceUnavailable(rw http.ResponseWriter)
	Success(rw http.ResponseWriter)
	Unauthorized(rw http.ResponseWriter)
}

type JSONResponder interface {
	BadRequest(rw http.ResponseWriter, err string)
	Created(rw http.ResponseWriter, respMap usecases.M)
	Forbidden(rw http.ResponseWriter)
	InternalServerError(rw http.ResponseWriter)
	NoContent(rw http.ResponseWriter)
	NotFound(rw http.ResponseWriter)
	ServiceUnavailable(rw http.ResponseWriter)
	Success(rw http.ResponseWriter, respMap usecases.M)
	Unauthorized(rw http.ResponseWriter)
}

type Gmailer interface {
	Send(from, to, subject, msg string) error
}
