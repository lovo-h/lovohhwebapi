package usecases

import "bytes"

type M map[string]interface{}

type EmailForm struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Subject         string `json:"subject"`
	Message         string `json:"message"`
	CaptchaID       string `json:"captchaid"`
	CaptchaSolution string `json:"captchasolution"`
}

type Logger interface {
	Log(message string) error
}

type EmailInteractor interface {
	SendEmail(EmailForm) error
}

type CaptchaInteractor interface {
	GetImageBuffer(id string, width, height int) (*bytes.Buffer, error)
	VerifyString(id, givenDigits string) bool
	Reload(id string) bool
	New() string
}
