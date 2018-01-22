package infrastructure

import (
	gmux "github.com/gorilla/mux"
	"github.com/lovohh/lovohhwebapi/interfaces"
	"net/http"
)

type VarsHandler func(w http.ResponseWriter, r *http.Request, vars map[string]string)

func (vh VarsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := gmux.Vars(r)
	vh(w, r, vars)
}

func GetRouterWithRoutes(webservice *interfaces.WebserviceHandler) *gmux.Router {
	router := gmux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/email", func(rw http.ResponseWriter, req *http.Request) {
		webservice.SendEmail(rw, req)
	}).Methods("POST")

	captchaRouter := apiRouter.PathPrefix("/captcha").Subrouter()

	captchaRouter.HandleFunc("", func(rw http.ResponseWriter, req *http.Request) {
		webservice.NewCaptcha(rw)
	}).Methods("GET")

	captchaRouter.HandleFunc("", func(rw http.ResponseWriter, req *http.Request) {
		webservice.ReloadCaptcha(rw, req)
	}).Methods("PUT")

	captchaRouter.Handle("/img/{id:[a-zA-Z0-9]+}",
		VarsHandler(func(rw http.ResponseWriter, req *http.Request, vars map[string]string) {
			webservice.GetCaptchaImage(rw, req, vars)
		})).Methods("GET")

	return router
}
