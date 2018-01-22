package interfaces

import (
	"github.com/stretchr/testify/assert"
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

func TestLogger_Log(t *testing.T) {
	logger := new(Logger)
	logErr := logger.Log("")

	assert.NotNil(t, logger)
	assert.Nil(t, logErr)
}
