package log

import (
	"sync"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type (
	// LoggerInterface ...
	LoggerInterface interface {
		Info(args ...interface{})
		Infof(format string, args ...interface{})
		Error(args ...interface{})
		Errorf(format string, args ...interface{})
	}

	// LoggerStruct ...
	LoggerStruct struct {
		LoggerInterface
		prefix string
	}
)

var (
	logger   *logrus.Logger
	once     sync.Once
	instance *LoggerStruct
)

// Info is a logrus log message at level info on the standard logger
func (q *LoggerStruct) Info(args ...interface{}) {
	q.Info(args...)
}

// Infof is a logrus log message at level infof on the standard logger
func (q *LoggerStruct) Infof(format string, args ...interface{}) {
	q.Infof(format, args...)
}

// Error is a logrus log message at level error on the standard logger
func (q *LoggerStruct) Error(args ...interface{}) {
	q.Error(args...)
}

// Errorf is a logrus log message at level errorf on the standard logger
func (q *LoggerStruct) Errorf(format string, args ...interface{}) {
	q.Errorf(format, args...)
}

func NewLog() *LoggerStruct {
	once.Do(func() {
		logger = logrus.New()
		logger.Formatter = &prefixed.TextFormatter{
			FullTimestamp: true,
		}

		instance = &LoggerStruct{
			prefix: "SERVICE_NAME",
		}
	})

	return instance
}
