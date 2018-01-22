package interfaces

import (
	"testing"

	"bytes"
	"encoding/json"
	"errors"
	"github.com/lovohh/lovohhwebapi/shared"
	"github.com/lovohh/lovohhwebapi/usecases"
	"github.com/stretchr/testify/assert"
	"net/http"
)

// =-=-=-=-=-=-=-=-=-=-=-=-=-= MOCKS

type MockResponder struct {
	ErrString  string
	SuccessMap usecases.M
}

func (interactor *MockResponder) BadRequest(rw http.ResponseWriter, errStr string) {
	interactor.ErrString = errStr
	rw.WriteHeader(http.StatusBadRequest)
}
func (interactor *MockResponder) Created(rw http.ResponseWriter, respMap usecases.M) {
	rw.WriteHeader(http.StatusCreated)
}
func (interactor *MockResponder) Forbidden(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusForbidden)
}
func (interactor *MockResponder) InternalServerError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
}
func (interactor *MockResponder) NoContent(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNoContent)
}
func (interactor *MockResponder) NotFound(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNotFound)
}
func (interactor *MockResponder) ServiceUnavailable(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusServiceUnavailable)
}
func (interactor *MockResponder) Success(rw http.ResponseWriter, respMap usecases.M) {
	interactor.SuccessMap = respMap
	rw.WriteHeader(http.StatusOK)
}
func (interactor *MockResponder) Unauthorized(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusUnauthorized)
}

type MockEmailInteractor struct {
	CalledSendEmail     bool
	SendEmailErr2Return error
}

func (interactor *MockEmailInteractor) SendEmail(form usecases.EmailForm) error {
	interactor.CalledSendEmail = true
	return interactor.SendEmailErr2Return
}

type MockCaptchaInteractor struct {
	CalledGetImageBuffer          bool
	CalledReload                  bool
	CalledNew                     bool
	CalledVerifyString            bool
	ShouldReturnFalseVerifyString bool
	ShouldReturnFalseReload       bool
	GivenGetImageBufferID         string
	GivenCaptchaID                string
	GivenCaptchaSolution          string
	GivenReloadID                 string
	GetImageBufferErr2Return      error
}

func (interactor *MockCaptchaInteractor) GetImageBuffer(id string, width, height int) (*bytes.Buffer, error) {
	interactor.CalledGetImageBuffer = true
	interactor.GivenGetImageBufferID = id

	return new(bytes.Buffer), interactor.GetImageBufferErr2Return
}
func (interactor *MockCaptchaInteractor) VerifyString(id, givenDigits string) bool {
	interactor.CalledVerifyString = true
	interactor.GivenCaptchaID = id
	interactor.GivenCaptchaSolution = givenDigits
	return !interactor.ShouldReturnFalseVerifyString
}
func (interactor *MockCaptchaInteractor) Reload(id string) bool {
	interactor.CalledReload = true
	interactor.GivenReloadID = id
	return !interactor.ShouldReturnFalseReload
}
func (interactor *MockCaptchaInteractor) New() string {
	interactor.CalledNew = true
	return "123456"
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= ACCESSORY

func webservice_captcha_email_responder() (*WebserviceHandler, *MockCaptchaInteractor, *MockEmailInteractor, *MockResponder) {
	webservice := new(WebserviceHandler)

	captcha := new(MockCaptchaInteractor)
	email := new(MockEmailInteractor)
	responder := new(MockResponder)
	webservice.CaptchaInteractor = captcha
	webservice.EmailInteractor = email
	webservice.Responder = responder

	return webservice, captcha, email, responder
}

func getValidEmailForm() usecases.EmailForm {
	return usecases.EmailForm{
		Name:            "Guest User",
		Email:           "guest.user@email.com",
		Subject:         "Unit Test",
		Message:         "This is a valid form.",
		CaptchaID:       "BnZET3i2zecyA8qXhLIb",
		CaptchaSolution: "748548",
	}
}

func convertStruct2Str(givenStruct interface{}) string {
	givenBytes, _ := json.Marshal(givenStruct)
	return string(givenBytes[:])
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= TESTS

func TestWebserviceHandler_SendEmail(t *testing.T) {
	webservice, captcha, email, _ := webservice_captcha_email_responder()
	emailFormJSONStr := convertStruct2Str(getValidEmailForm())
	rr, req := shared.GetRecorderAndModdedRequest("GET", "/", emailFormJSONStr)

	webservice.SendEmail(rr, req) // objective

	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.True(t, email.CalledSendEmail)
	assert.True(t, captcha.CalledVerifyString)
	assert.Equal(t, "BnZET3i2zecyA8qXhLIb", captcha.GivenCaptchaID)
	assert.Equal(t, "748548", captcha.GivenCaptchaSolution)
}

func TestWebserviceHandler_SendEmail_InvalidCaptchaDigitsOrID(t *testing.T) {
	webservice, captcha, email, responder := webservice_captcha_email_responder()
	emailFormJSONStr := convertStruct2Str(getValidEmailForm())
	rr, req := shared.GetRecorderAndModdedRequest("GET", "/", emailFormJSONStr)
	captcha.ShouldReturnFalseVerifyString = true
	webservice.SendEmail(rr, req) // objective

	assert.True(t, captcha.CalledVerifyString)
	assert.False(t, email.CalledSendEmail)
	assert.Equal(t, "unverified captcha", responder.ErrString)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestWebserviceHandler_SendEmail_InvalidForm(t *testing.T) {
	webservice, captcha, email, responder := webservice_captcha_email_responder()
	emailFormJSONStr := convertStruct2Str(getValidEmailForm())
	rr, req := shared.GetRecorderAndModdedRequest("GET", "/", emailFormJSONStr)
	email.SendEmailErr2Return = errors.New("invalid form")
	webservice.SendEmail(rr, req) // objective

	assert.True(t, captcha.CalledVerifyString)
	assert.True(t, email.CalledSendEmail)
	assert.Equal(t, "invalid form", responder.ErrString)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestWebserviceHandler_SendEmail_FailedSendEmail(t *testing.T) {
	webservice, captcha, email, _ := webservice_captcha_email_responder()
	emailFormJSONStr := convertStruct2Str(getValidEmailForm())
	rr, req := shared.GetRecorderAndModdedRequest("GET", "/", emailFormJSONStr)
	email.SendEmailErr2Return = errors.New("irrelevant error")
	webservice.SendEmail(rr, req) // objective

	assert.True(t, captcha.CalledVerifyString)
	assert.True(t, email.CalledSendEmail)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestWebserviceHandler_NewCaptcha(t *testing.T) {
	webservice, captcha, _, responder := webservice_captcha_email_responder()
	rr, _ := shared.GetDefaultRecorderAndRequest()

	webservice.NewCaptcha(rr)

	assert.True(t, captcha.CalledNew)
	assert.Equal(t, responder.SuccessMap, usecases.M{"id": "123456"})
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestWebserviceHandler_ReloadCaptcha(t *testing.T) {
	webservice, captcha, _, _ := webservice_captcha_email_responder()
	rr, req := shared.GetRecorderAndModdedRequest("PUT", "/", `{"id":"987654"}`)
	webservice.ReloadCaptcha(rr, req)

	assert.True(t, captcha.CalledReload)
	assert.Equal(t, "987654", captcha.GivenReloadID)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestWebserviceHandler_ReloadCaptcha_FailedReload(t *testing.T) {
	webservice, captcha, _, responder := webservice_captcha_email_responder()
	rr, req := shared.GetRecorderAndModdedRequest("GET", "/", `{"id":"987654"}`)
	captcha.ShouldReturnFalseReload = true
	webservice.ReloadCaptcha(rr, req)

	assert.True(t, captcha.CalledReload)
	assert.Equal(t, "987654", captcha.GivenReloadID)
	assert.Equal(t, "invalid captcha id", responder.ErrString)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestWebserviceHandler_GetCaptchaImage(t *testing.T) {
	webservice, captcha, _, _ := webservice_captcha_email_responder()
	rr, req := shared.GetRecorderAndModdedRequest("GET", "/", "")
	webservice.GetCaptchaImage(rr, req, map[string]string{"id": "456789"}) // objective

	assert.True(t, captcha.CalledGetImageBuffer)
	assert.Equal(t, "456789", captcha.GivenGetImageBufferID)
	assert.Equal(t, "no-cache, no-store, must-revalidate", rr.Header().Get("Cache-Control"))
	assert.Equal(t, "no-cache", rr.Header().Get("Pragma"))
	assert.Equal(t, "0", rr.Header().Get("Expires"))
	assert.Equal(t, "image/png", rr.Header().Get("Content-Type"))
	assert.Equal(t, "bytes", rr.Header().Get("Accept-Ranges"))
	assert.Equal(t, "0", rr.Header().Get("Content-Length"))
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestWebserviceHandler_GetCaptchaImage_ImageErr(t *testing.T) {
	webservice, captcha, _, _ := webservice_captcha_email_responder()
	rr, req := shared.GetRecorderAndModdedRequest("GET", "/", "")
	captcha.GetImageBufferErr2Return = errors.New("irrelevant error")
	webservice.GetCaptchaImage(rr, req, map[string]string{"id": "456789"}) // objective

	assert.True(t, captcha.CalledGetImageBuffer)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
