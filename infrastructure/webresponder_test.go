package infrastructure

import (
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestWebResponder_BadRequest(t *testing.T) {
	wr := WebResponder{}
	actual := getActualStatusCode(wr.BadRequest)
	assert.Equal(t, http.StatusBadRequest, actual)
}

func TestWebResponder_Created(t *testing.T) {
	wr := WebResponder{}
	actual := getActualStatusCode(wr.Created)
	assert.Equal(t, http.StatusCreated, actual)
}

func TestWebResponder_Forbidden(t *testing.T) {
	wr := WebResponder{}
	actual := getActualStatusCode(wr.Forbidden)
	assert.Equal(t, http.StatusForbidden, actual)
}

func TestWebResponder_InternalServerError(t *testing.T) {
	wr := WebResponder{}
	actual := getActualStatusCode(wr.InternalServerError)
	assert.Equal(t, http.StatusInternalServerError, actual)
}

func TestWebResponder_NoContent(t *testing.T) {
	wr := WebResponder{}
	actual := getActualStatusCode(wr.NoContent)
	assert.Equal(t, http.StatusNoContent, actual)
}

func TestWebResponder_NotFound(t *testing.T) {
	wr := WebResponder{}
	actual := getActualStatusCode(wr.NotFound)
	assert.Equal(t, http.StatusNotFound, actual)
}

func TestWebResponder_ServiceUnavailable(t *testing.T) {
	wr := WebResponder{}
	actual := getActualStatusCode(wr.ServiceUnavailable)
	assert.Equal(t, http.StatusServiceUnavailable, actual)
}

func TestWebResponder_Success(t *testing.T) {
	wr := WebResponder{}
	actual := getActualStatusCode(wr.Success)
	assert.Equal(t, http.StatusOK, actual)
}

func TestWebResponder_Unauthorized(t *testing.T) {
	wr := WebResponder{}
	actual := getActualStatusCode(wr.Unauthorized)
	assert.Equal(t, http.StatusUnauthorized, actual)
}

func getActualStatusCode(fn func(http.ResponseWriter)) int {
	nr := httptest.NewRecorder()
	fn(nr)
	return nr.Result().StatusCode
}
