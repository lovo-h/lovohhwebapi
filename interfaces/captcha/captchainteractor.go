package captcha

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/lovohh/lovohhwebapi/usecases"
	"time"
)

const (
	DefaultLen        = 6
	DefaultCollectNum = 100
	DefaultExpiration = 20
)

type CaptchaInteractor struct {
	Store  captcha.Store
	Logger usecases.Logger
}

func (interactor *CaptchaInteractor) New() string {
	interactor.Logger.Log("Generating new CAPTCHA")
	id := randomId()
	interactor.Store.Set(id, RandomDigits(DefaultLen))
	return id
}

func (interactor *CaptchaInteractor) Reload(id string) bool {
	interactor.Logger.Log("Reloading some predefined CAPTCHA")
	old := interactor.Store.Get(id, false)
	if old == nil {
		interactor.Logger.Log("\t - Invalid CAPTCHA id")
		return false
	}

	interactor.Store.Set(id, RandomDigits(len(old)))

	return true
}

func (interactor *CaptchaInteractor) GetImageBuffer(id string, width, height int) (*bytes.Buffer, error) {
	interactor.Logger.Log("Retrieving CAPTCHA image")
	content := new(bytes.Buffer)

	d := interactor.Store.Get(id, false)
	if d == nil {
		interactor.Logger.Log("\t - Invalid CAPTCHA id")
		return nil, captcha.ErrNotFound
	}

	_, err := captcha.NewImage(id, d, width, height).WriteTo(content)

	return content, err
}

func (interactor *CaptchaInteractor) VerifyString(id, givenDigits string) bool {
	interactor.Logger.Log("Attempting to verify CAPTCHA")
	if len(givenDigits) == 0 {
		interactor.Logger.Log("\t- Empty given-digits")
		return false
	}

	digitsByte := make([]byte, len(givenDigits))

	for idx := range digitsByte {
		d := givenDigits[idx]
		switch {
		case '0' <= d && d <= '9':
			digitsByte[idx] = d - '0'
		default:
			interactor.Logger.Log("\t- Invalid given-digits")
			return false
		}
	}
	return interactor.verify(id, digitsByte)
}

// Accessory function to VerifyString. Should only be called from VerifyString.
func (interactor *CaptchaInteractor) verify(id string, givenDigits []byte) bool {
	expectedDigits := interactor.Store.Get(id, true)

	if expectedDigits == nil {
		interactor.Logger.Log("\t- Invalid CAPTCHA id")
		return false
	}

	isVerified := bytes.Equal(givenDigits, expectedDigits)

	if !isVerified {
		interactor.Logger.Log("\t- Failed verify CAPTCHA")
		return false
	}

	interactor.Logger.Log("\t- Successfully verified CAPTCHA")
	return isVerified
}

func InitCaptchaHandler(logger usecases.Logger) *CaptchaInteractor {
	captchaHandler := new(CaptchaInteractor)
	captchaHandler.Store = captcha.NewMemoryStore(DefaultCollectNum, DefaultExpiration*time.Minute)
	captchaHandler.Logger = logger

	return captchaHandler
}
