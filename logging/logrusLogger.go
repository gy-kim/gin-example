package logging

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type logrusLogger struct{}

func (l logrusLogger) Debug(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	logrus.Debug(msg)
}

func (l logrusLogger) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}
