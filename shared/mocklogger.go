package shared

type MockLogger struct {
	IsCalled      bool
	StoredMessage []string
}

func (logger *MockLogger) Log(msg string) error {
	logger.IsCalled = true
	logger.StoredMessage = append(logger.StoredMessage, msg)
	return nil
}
