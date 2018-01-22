package middleware

import (
	"fmt"
	"github.com/lovohh/lovohhwebapi/usecases"
	"net/http"
	"time"
)

type WebLoggerMiddleware struct {
	MiddlewareHandler
}

func (middleware *WebLoggerMiddleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	t1 := time.Now()
	middleware.next.ServeHTTP(rw, req)
	msg := fmt.Sprintf("\t%s | [%s] %q | %v\n", req.Host, req.Method, req.URL.String(), time.Now().Sub(t1))
	middleware.Logger.Log(msg)
}

func (middleware *WebLoggerMiddleware) Handle(next http.Handler) http.Handler {
	middleware.next = next
	return middleware
}

func GetAndInitWebLoggerMiddleware(logger usecases.Logger) *WebLoggerMiddleware {
	weblogger := new(WebLoggerMiddleware)
	weblogger.Logger = logger

	return weblogger
}
