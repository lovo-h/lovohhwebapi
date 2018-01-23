package interfaces

import (
	"bytes"
	"encoding/json"
	"github.com/lovohh/lovohhwebapi/usecases"
	"io"
	"net/http"
	"time"
)

const (
	CaptchaWidth  = 310
	CaptchaHeight = 46
)

type WebserviceHandler struct {
	Responder         JSONResponder
	EmailInteractor   usecases.EmailInteractor
	CaptchaInteractor usecases.CaptchaInteractor
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-=- EMAIL

//  /api/email [POST]
func (webhandler *WebserviceHandler) SendEmail(rw http.ResponseWriter, req *http.Request) {
	emailForm := usecases.EmailForm{}

	if decErr := json.NewDecoder(req.Body).Decode(&emailForm); decErr != nil && decErr != io.EOF {
		webhandler.Responder.InternalServerError(rw)
		return
	}

	if !webhandler.CaptchaInteractor.VerifyString(emailForm.CaptchaID, emailForm.CaptchaSolution) {
		webhandler.Responder.BadRequest(rw, "unverified captcha")
		return
	}

	if sendErr := webhandler.EmailInteractor.SendEmail(emailForm); sendErr != nil {
		if sendErr.Error() == "invalid form" {
			webhandler.Responder.BadRequest(rw, sendErr.Error())
		} else {
			webhandler.Responder.InternalServerError(rw)
		}

		return
	}

	webhandler.Responder.NoContent(rw)
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-=- CAPTCHA

//  /api/captcha [GET]
func (webhandler *WebserviceHandler) NewCaptcha(rw http.ResponseWriter) {
	id := webhandler.CaptchaInteractor.New()

	webhandler.Responder.Success(rw, usecases.M{"id": id})
}

//  /api/captcha [PUT]
func (webhandler *WebserviceHandler) ReloadCaptcha(rw http.ResponseWriter, req *http.Request) {
	idMap := usecases.M{}
	if decErr := json.NewDecoder(req.Body).Decode(&idMap); decErr != nil && decErr != io.EOF {
		webhandler.Responder.InternalServerError(rw)
		return
	}

	id := idMap["id"].(string)
	if !webhandler.CaptchaInteractor.Reload(id) {
		webhandler.Responder.BadRequest(rw, "invalid captcha id")
		return
	}

	webhandler.Responder.NoContent(rw)
}

//  /api/captcha/img/{id} [GET]
func (webhandler *WebserviceHandler) GetCaptchaImage(rw http.ResponseWriter, req *http.Request, vars map[string]string) {
	id := vars["id"]
	buffer, imageErr := webhandler.CaptchaInteractor.GetImageBuffer(id, CaptchaWidth, CaptchaHeight)

	if imageErr != nil {
		webhandler.Responder.InternalServerError(rw)
		return
	}

	rw.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	rw.Header().Set("Pragma", "no-cache")
	rw.Header().Set("Expires", "0")
	rw.Header().Set("Content-Type", "image/png")

	http.ServeContent(rw, req, id+".png", time.Time{}, bytes.NewReader(buffer.Bytes()))
}
