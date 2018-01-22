package infrastructure

import (
	"encoding/json"
	"github.com/lovohh/lovohhwebapi/shared"
	"github.com/lovohh/lovohhwebapi/usecases"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

// =-=-=-=-=-=-=-=-=-=-=-=-=-= SETUP

func TestMain(m *testing.M) {
	// setup
	retCode := m.Run()
	// teardown
	os.Exit(retCode)
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= TESTS

func TestJSONWebResponder_Success(t *testing.T) {
	responder := new(JSONWebResponder)
	rr, _ := shared.GetDefaultRecorderAndRequest()
	respMap := usecases.M{"Hello": "World"}
	responder.Success(rr, respMap) // objective

	var result usecases.M
	json.Unmarshal(rr.Body.Bytes(), &result)

	assert.Equal(t, respMap, result)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

func TestJSONWebResponder_Created(t *testing.T) {
	responder := new(JSONWebResponder)
	rr, _ := shared.GetDefaultRecorderAndRequest()
	responder.Created(rr, nil)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestJSONWebResponder_BadRequest(t *testing.T) {
	responder := new(JSONWebResponder)
	rr, _ := shared.GetDefaultRecorderAndRequest()
	responder.BadRequest(rr, "some failure")

	var result usecases.M
	json.Unmarshal(rr.Body.Bytes(), &result)

	assert.Equal(t, usecases.M{"error": "some failure"}, result)
	assert.Equal(t, http.StatusBadRequest, rr.Code)

}
