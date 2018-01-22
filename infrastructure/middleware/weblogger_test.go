package middleware

import (
	"github.com/lovohh/lovohhwebapi/shared"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

// =-=-=-=-=-=-=-=-=-=-=-=-=-= MOCKS

type MockHttpHandler struct {
	invokedServeHttp bool
}

func (handler *MockHttpHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	handler.invokedServeHttp = true
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= SETUP

func logger_weblogger() (*shared.MockLogger, *WebLoggerMiddleware) {
	logger := new(shared.MockLogger)
	weblogger := new(WebLoggerMiddleware)
	weblogger.Logger = logger

	return logger, weblogger
}

func TestMain(m *testing.M) {
	// setup
	retCode := m.Run()
	// teardown
	os.Exit(retCode)
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= TESTS

func TestWebLoggerMiddleware_ServeHTTP(t *testing.T) {
	logger, weblogger := logger_weblogger()
	nextHttpHandler := new(MockHttpHandler)
	weblogger.Handle(nextHttpHandler)

	rw, req := shared.GetDefaultRecorderAndRequest()
	weblogger.ServeHTTP(rw, req) // objective

	assert.True(t, logger.IsCalled)
	assert.Len(t, logger.StoredMessage, 1)
	assert.True(t, nextHttpHandler.invokedServeHttp)
}

func TestGetAndInitWebLoggerMiddleware(t *testing.T) {
  logger, _ := logger_weblogger()
  weblogger := GetAndInitWebLoggerMiddleware(logger)
  weblogger.Logger.Log("")

  assert.NotNil(t, weblogger)
  assert.True(t, logger.IsCalled)
}
