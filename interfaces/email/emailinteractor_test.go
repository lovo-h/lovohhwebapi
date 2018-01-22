package email

import (
	"errors"
	"github.com/lovohh/lovohhwebapi/shared"
	"github.com/lovohh/lovohhwebapi/usecases"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// =-=-=-=-=-=-=-=-=-=-=-=-=-= MOCKS

type MockGmailer struct {
	ShouldReturnSendErr bool
}

func (interactor *MockGmailer) Send(from, to, subject, msg string) error {
	if interactor.ShouldReturnSendErr {
		return errors.New("mocked error")
	}
	return nil
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= SETUP

func TestMain(m *testing.M) {
	// setup
	retCode := m.Run()
	// teardown
	os.Exit(retCode)
}

func logger_email() (*shared.MockLogger, *MockGmailer, *EmailInteractor) {
	logger := new(shared.MockLogger)
	gmailer := new(MockGmailer)
	email := new(EmailInteractor)
	email.Logger = logger
	email.Gmailer = gmailer

	return logger, gmailer, email
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= ACCESSORY

func getValidEmailForm() usecases.EmailForm {
	return usecases.EmailForm{
		Name:    "John Does",
		Email:   "john.doe@email.com",
		Subject: "Unit Test",
		Message: "Testing SendEmail()",
	}
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= TESTS

func TestGetAndInitEmailInteractor(t *testing.T) {
	logger, gmailer, _ := logger_email()
	email := GetAndInitEmailInteractor(logger, gmailer)
	email.Logger.Log("")

	assert.NotNil(t, email)
	assert.NotNil(t, email.Logger)
	assert.NotNil(t, email.Gmailer)
	assert.True(t, logger.IsCalled)
}

func TestEmailInteractor_SendEmail(t *testing.T) {
	logger, _, email := logger_email()
	emailForm := getValidEmailForm()
	sendErr := email.SendEmail(emailForm)

	assert.Nil(t, sendErr)
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "Attempting to send email", logger.StoredMessage[0])
	assert.Equal(t, "\t- Successfully sent email", logger.StoredMessage[1])
}

func TestEmailInteractor_SendEmail_EmptyFormName(t *testing.T) {
	logger, _, email := logger_email()
	emailForm := getValidEmailForm()
	emailForm.Name = ""
	sendErr := email.SendEmail(emailForm)

	assert.NotNil(t, sendErr)
	assert.Equal(t, "invalid form", sendErr.Error())
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "\t- Failed to verify email-form", logger.StoredMessage[1])
}

func TestEmailInteractor_SendEmail_EmptyFormEmail(t *testing.T) {
	logger, _, email := logger_email()
	emailForm := getValidEmailForm()
	emailForm.Email = ""
	sendErr := email.SendEmail(emailForm)

	assert.NotNil(t, sendErr)
	assert.Equal(t, "invalid form", sendErr.Error())
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "\t- Failed to verify email-form", logger.StoredMessage[1])
}

func TestEmailInteractor_SendEmail_EmptyFormSubject(t *testing.T) {
	logger, _, email := logger_email()
	emailForm := getValidEmailForm()
	emailForm.Subject = ""
	sendErr := email.SendEmail(emailForm)

	assert.NotNil(t, sendErr)
	assert.Equal(t, "invalid form", sendErr.Error())
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "\t- Failed to verify email-form", logger.StoredMessage[1])
}

func TestEmailInteractor_SendEmail_EmptyFormMessage(t *testing.T) {
	logger, _, email := logger_email()
	emailForm := getValidEmailForm()
	emailForm.Message = ""
	sendErr := email.SendEmail(emailForm)

	assert.NotNil(t, sendErr)
	assert.Equal(t, "invalid form", sendErr.Error())
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "\t- Failed to verify email-form", logger.StoredMessage[1])
}

func TestEmailInteractor_SendEmail_SendError(t *testing.T) {
	logger, gmailer, email := logger_email()
	gmailer.ShouldReturnSendErr = true
	emailForm := getValidEmailForm()
	sendErr := email.SendEmail(emailForm)

	assert.NotNil(t, sendErr)
	assert.Equal(t, "mocked error", sendErr.Error())
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "\t- Failed to send email: mocked error", logger.StoredMessage[1])
}
