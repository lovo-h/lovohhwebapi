package middleware

import (
	"github.com/lovohh/lovohhwebapi/interfaces"
	"github.com/lovohh/lovohhwebapi/usecases"
	"net/http"
)

type MiddlewareHandler struct {
	Responder interfaces.WebResponder
	Logger    usecases.Logger
	next      http.Handler
}
