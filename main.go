package main

import (
	"github.com/justinas/alice"
	"github.com/lovohh/lovohhwebapi/infrastructure"
	_ "github.com/lovohh/lovohhwebapi/infrastructure/deps"
	"github.com/lovohh/lovohhwebapi/infrastructure/middleware"
	"github.com/lovohh/lovohhwebapi/interfaces"
	"github.com/lovohh/lovohhwebapi/interfaces/captcha"
	"github.com/lovohh/lovohhwebapi/interfaces/email"
	"net/http"
)

func main() {
	// Webservice
	logger := new(interfaces.Logger)
	gmailer := infrastructure.GetAndInitGMailer()
	jsonResponder := new(infrastructure.JSONWebResponder)
	emailInteractor := email.GetAndInitEmailInteractor(logger, gmailer)
	captchaInteractor := captcha.InitCaptchaHandler(logger)
	webserviceHandler := &interfaces.WebserviceHandler{
		jsonResponder,
		emailInteractor,
		captchaInteractor,
	}

	// Middleware

	weblogger := middleware.GetAndInitWebLoggerMiddleware(logger)

	globalMiddleware := []alice.Constructor{weblogger.Handle}

	router := infrastructure.GetRouterWithRoutes(webserviceHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: alice.New(globalMiddleware...).Then(router),
	}

	if serverErr := server.ListenAndServe(); serverErr != nil {
		logger.Log("Server failed to boot: " + serverErr.Error())
	}
}
