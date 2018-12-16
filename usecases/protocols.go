package usecases

import "bytes"

type M map[string]interface{}

type EmailForm struct {
	Email           string `json:"email"`
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
