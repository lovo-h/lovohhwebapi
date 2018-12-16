package email

import (
	"errors"
	"fmt"
	"github.com/lovohh/lovohhwebapi/interfaces"
	"github.com/lovohh/lovohhwebapi/usecases"
)

type EmailInteractor struct {
	Logger  usecases.Logger
	Gmailer interfaces.Gmailer
}

func (interactor *EmailInteractor) SendEmail(form usecases.EmailForm) error {
	interactor.Logger.Log("Attempting to send email")

	if verifyErr := interactor.verifyEmailForm(form); verifyErr != nil {
		interactor.Logger.Log("\t- Failed to verify email-form")
		return verifyErr
	}

	subject := fmt.Sprintf("LovoHH Website Notification: From %s", form.Email)
	sendErr := interactor.Gmailer.Send(form.Email, "lovohh@gmail.com", subject, form.Message)

	if sendErr != nil {
		interactor.Logger.Log("\t- Failed to send email: " + sendErr.Error())
		return sendErr
	}

	interactor.Logger.Log("\t- Successfully sent email")
	return nil
}

func (interactor *EmailInteractor) verifyEmailForm(ef usecases.EmailForm) error {
	if len(ef.Message) == 0 {
		return errors.New("invalid form")
	}

	return nil
}

func GetAndInitEmailInteractor(logger usecases.Logger, gmailer interfaces.Gmailer) *EmailInteractor {
	emailInteractor := new(EmailInteractor)
	emailInteractor.Logger = logger
	emailInteractor.Gmailer = gmailer

	return emailInteractor
}
