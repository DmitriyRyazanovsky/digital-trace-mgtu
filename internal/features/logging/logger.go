package logging

import (
	"path"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

type Level int

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

const (
	filename = "./logs/console.log"
)

type logger struct{ log *logrus.Logger }

func (l *logger) Info(msg string) {
	l.log.Info(msg)
}

func (l *logger) Infof(fmt string, args ...interface{}) {
	l.log.Infof(fmt, args...)
}

func (l *logger) InfoWithField(msg string, k string, v interface{}) {
	l.log.WithField(k, v).Info(msg)
}

func (l *logger) InfoWithFields(msg string, fields map[string]interface{}) {
	l.log.WithFields(fields).Info(msg)
}

func (l *logger) Error(msg string) {
	l.log.Error(msg)
}

func (l *logger) Errorf(fmt string, args ...interface{}) {
	l.log.Errorf(fmt, args...)
}

func (l *logger) ErrorWithField(err error, k string, v interface{}) {
	l.log.WithField(k, v).Error(err.Error())
}

func (l *logger) ErrorWithFields(err error, fields map[string]interface{}) {
	l.log.WithFields(fields).Error(err.Error())
}

func (l *logger) Debug(msg string) {
	l.log.Debug(msg)
}

func (l *logger) Debugf(fmt string, args ...interface{}) {
	l.log.Debugf(fmt, args...)
}

type Config struct {
	level  Level
	dir    string
	inFile bool
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetLevel(level Level) {
	c.level = level
}

func (c *Config) SetOutputDir(dir string) {
	c.dir = dir
}

func (c *Config) SetInFile(inFile bool) {
	c.inFile = inFile
}

func logLevelConverter(level Level) logrus.Level {
	switch level {
	case PanicLevel:
		return logrus.PanicLevel
	case FatalLevel:
		return logrus.FatalLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	case WarnLevel:
		return logrus.WarnLevel
	case InfoLevel:
		return logrus.InfoLevel
	case DebugLevel:
		return logrus.DebugLevel
	case TraceLevel:
		return logrus.TraceLevel
	}

	return logrus.InfoLevel
}

func New(c *Config) (Logger, error) {
	logLevel := logLevelConverter(c.level)

	log := logrus.New()
	log.SetLevel(logLevel)

	log.SetOutput(colorable.NewColorableStdout())
	log.SetFormatter(&logrus.TextFormatter{
		PadLevelText:    true,
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if c.inFile {
		rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
			Filename:   path.Join(c.dir, filename),
			MaxSize:    50,
			MaxBackups: 3,
			MaxAge:     28,
			Level:      logLevel,
			Formatter: &logrus.JSONFormatter{
				TimestampFormat: time.RFC822,
			},
		})

		if err != nil {
			return nil, err
		}

		log.AddHook(rotateFileHook)
	}

	return &logger{log}, nil
}
