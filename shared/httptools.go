package shared

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

func GetDefaultRecorderAndRequest() (*httptest.ResponseRecorder, *http.Request) {
	return GetRecorderAndModdedRequest("GET", "/", "")
}

func GetRecorderAndModdedRequest(method, target, jsonStr string) (*httptest.ResponseRecorder, *http.Request) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(method, target, strings.NewReader(jsonStr))

	return rr, req
}
