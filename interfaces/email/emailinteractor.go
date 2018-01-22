package email

import (
	"errors"
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

	sendErr := interactor.Gmailer.Send(form.Email, "lovohh@gmail.com", form.Subject, form.Message)

	if sendErr != nil {
		interactor.Logger.Log("\t- Failed to send email: " + sendErr.Error())
		return sendErr
	}

	interactor.Logger.Log("\t- Successfully sent email")
	return nil
}

func (interactor *EmailInteractor) verifyEmailForm(ef usecases.EmailForm) error {
	errExists := false
	for _, el := range []string{ef.Name, ef.Email, ef.Subject, ef.Message} {
		if len(el) == 0 {
			errExists = true
			break
		}
	}

	if errExists {
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
