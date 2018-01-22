package captcha

import (
	"bytes"
	"github.com/lovohh/lovohhwebapi/shared"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

// =-=-=-=-=-=-=-=-=-=-=-=-=-= MOCKS

type MockStore struct {
	memory map[string][]byte
}

func (store *MockStore) Set(id string, digits []byte) {
	store.memory[id] = digits
}
func (store *MockStore) Get(id string, clear bool) []byte {
	return store.memory[id]
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= SETUP

func TestMain(m *testing.M) {
	// setup
	retCode := m.Run()
	// teardown
	os.Exit(retCode)
}

func logger_captcha() (*shared.MockLogger, *CaptchaInteractor) {
	logger := new(shared.MockLogger)
	captcha := InitCaptchaHandler(logger)

	return logger, captcha
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= ACCESSORY

func byteArr2Str(byteArr []byte) string {
	var buffer bytes.Buffer

	for idx := range byteArr {
		buffer.WriteString(strconv.Itoa(int(byteArr[idx])))
	}

	return buffer.String()
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-= TESTS

func TestInitCaptchaHandler(t *testing.T) {
	logger, captcha := logger_captcha() // objective
	captcha.Logger.Log("")

	assert.NotNil(t, captcha)
	assert.NotNil(t, captcha.Store)
	assert.True(t, logger.IsCalled)
}

func TestCaptchaInteractor_New(t *testing.T) {
	logger, captcha := logger_captcha()
	id := captcha.New() // objective
	result := captcha.Store.Get(id, false)

	assert.NotNil(t, result)
	assert.Len(t, logger.StoredMessage, 1)
	assert.Equal(t, "Generating new CAPTCHA", logger.StoredMessage[0])
}

func TestCaptchaInteractor_Reload(t *testing.T) {
	logger, captcha := logger_captcha()
	id := captcha.New()
	firstResult := captcha.Store.Get(id, false)
	isReloaded := captcha.Reload(id) // objective
	secondResult := captcha.Store.Get(id, false)
	isNotEqual := !bytes.Equal(firstResult, secondResult)

	assert.True(t, isReloaded && isNotEqual)
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "Reloading some predefined CAPTCHA", logger.StoredMessage[1])
}

func TestCaptchaInteractor_Reload_InvalidID(t *testing.T) {
	logger, captcha := logger_captcha()
	isReloaded := captcha.Reload("1") // objective

	assert.False(t, isReloaded)
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "Reloading some predefined CAPTCHA", logger.StoredMessage[0])
	assert.Equal(t, "\t - Invalid CAPTCHA id", logger.StoredMessage[1])
}

func TestCaptchaInteractor_GetImageBuffer(t *testing.T) {
	logger, captcha := logger_captcha()
	id := captcha.New()
	imageBuffer, getErr := captcha.GetImageBuffer(id, 80, 45) // objective

	assert.True(t, len(imageBuffer.Bytes()) > 0)
	assert.Nil(t, getErr)
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "Retrieving CAPTCHA image", logger.StoredMessage[1])
}

func TestCaptchaInteractor_GetImageBuffer_InvalidID(t *testing.T) {
	logger, captcha := logger_captcha()
	imageBuffer, getErr := captcha.GetImageBuffer("1", 80, 45) // objective

	assert.NotNil(t, getErr)
	assert.Nil(t, imageBuffer)
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "Retrieving CAPTCHA image", logger.StoredMessage[0])
	assert.Equal(t, "\t - Invalid CAPTCHA id", logger.StoredMessage[1])
}

func TestCaptchaInteractor_VerifyString(t *testing.T) {
	logger, captcha := logger_captcha()
	captcha.Store = &MockStore{make(map[string][]byte)}
	id := captcha.New()
	solutionDigits := captcha.Store.Get(id, false)
	expected := byteArr2Str(solutionDigits)
	isValid := captcha.VerifyString(id, expected) // objective

	assert.True(t, isValid)
	assert.Len(t, logger.StoredMessage, 3)
	assert.Equal(t, "Attempting to verify CAPTCHA", logger.StoredMessage[1])
	assert.Equal(t, "\t- Successfully verified CAPTCHA", logger.StoredMessage[2])
}

func TestCaptchaInteractor_VerifyString_NonNumericGivenDigits(t *testing.T) {
	logger, captcha := logger_captcha()
	isValid := captcha.VerifyString("1", "") // objective

	assert.False(t, isValid)
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "\t- Empty given-digits", logger.StoredMessage[1])
}

func TestCaptchaInteractor_VerifyString_InvalidGivenDigits(t *testing.T) {
	logger, captcha := logger_captcha()
	isValid := captcha.VerifyString("1", "abc") // objective

	assert.False(t, isValid)
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "\t- Invalid given-digits", logger.StoredMessage[1])
}

func TestCaptchaInteractor_VerifyString_InvalidID(t *testing.T) {
	logger, captcha := logger_captcha()
	isValid := captcha.VerifyString("1", "123") // objective

	assert.False(t, isValid)
	assert.Len(t, logger.StoredMessage, 2)
	assert.Equal(t, "\t- Invalid CAPTCHA id", logger.StoredMessage[1])
}

func TestCaptchaInteractor_VerifyString_InvalidDigits(t *testing.T) {
	logger, captcha := logger_captcha()
	id := captcha.New()
	isValid := captcha.VerifyString(id, "123456") // objective

	assert.False(t, isValid)
	assert.Len(t, logger.StoredMessage, 3)
	assert.Equal(t, "\t- Failed verify CAPTCHA", logger.StoredMessage[2])
}
