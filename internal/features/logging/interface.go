package logging

type Logger interface {
	Info(msg string)
	Infof(fmt string, args ...interface{})
	InfoWithField(msg string, k string, v interface{})
	InfoWithFields(msg string, fields map[string]interface{})
	Error(msg string)
	Errorf(fmt string, args ...interface{})
	ErrorWithField(err error, k string, v interface{})
	ErrorWithFields(err error, fields map[string]interface{})
	Debug(msg string)
	Debugf(fmt string, args ...interface{})
}
