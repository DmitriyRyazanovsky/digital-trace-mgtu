package logging

type mockLogger struct{}

func (l *mockLogger) Info(msg string) {}

func (l *mockLogger) Infof(fmt string, args ...interface{}) {}

func (l *mockLogger) InfoWithField(msg string, k string, v interface{}) {}

func (l *mockLogger) InfoWithFields(msg string, fields map[string]interface{}) {}

func (l *mockLogger) Error(msg string) {}

func (l *mockLogger) Errorf(fmt string, args ...interface{}) {}

func (l *mockLogger) ErrorWithField(err error, k string, v interface{}) {}

func (l *mockLogger) ErrorWithFields(err error, fields map[string]interface{}) {}

func (l *mockLogger) Debug(msg string) {}

func (l *mockLogger) Debugf(fmt string, args ...interface{}) {}

func NewMock() Logger {
	return &mockLogger{}
}
