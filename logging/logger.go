package logging

// Logger is logging interface
type Logger interface {
	Debug(message string, args ...interface{})
	Debugf(format string, args ...interface{})
	Info(message string, args ...interface{})
	Infof(format string, args ...interface{})
	Warn(message string, args ...interface{})
	Warnf(format string, args ...interface{})
	Error(message string, args ...interface{})
	Errorf(format string, args ...interface{})
}
