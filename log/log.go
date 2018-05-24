package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	log.SetLevel(logrus.InfoLevel)
	log.Out = os.Stdout
}

// SetLevel : inits log level
func SetLevel(debug bool) {
	l := logrus.InfoLevel
	if debug {
		l = logrus.DebugLevel
	}
	log.SetLevel(l)
}

// Info wrap
func Info(format string, a ...interface{}) {
	log.Info(sFormat(format, a...))
}

// Debug wrap
func Debug(format string, a ...interface{}) {
	log.Debug(sFormat(format, a))
}

// Warn wrap
func Warn(format string, a ...interface{}) {
	log.Warn(sFormat(format, a))
}

// Error wrap
func Error(format string, a ...interface{}) {
	log.Error(sFormat(format, a))
}

// Fatal wrap
func Fatal(format string, a ...interface{}) {
	log.Fatal(sFormat(format, a))
}

// Panic wrap
func Panic(format string, a ...interface{}) {
	log.Panic(sFormat(format, a))
}

// format helper
func sFormat(format string, a ...interface{}) string {
	if len(a) > 0 {
		return fmt.Sprintf(format, a...)
	} else {
		return format
	}
}
