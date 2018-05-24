package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
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
	log.Info(fmt.Sprintf(format, a))
}

// Debug wrap
func Debug(format string, a ...interface{}) {
	log.Debug(fmt.Sprintf(format, a))
}

// Warn wrap
func Warn(format string, a ...interface{}) {
	log.Warn(fmt.Sprintf(format, a))
}

// Error wrap
func Error(format string, a ...interface{}) {
	log.Error(fmt.Sprintf(format, a))
}

// Fatal wrap
func Fatal(format string, a ...interface{}) {
	log.Fatal(fmt.Sprintf(format, a))
}

// Panic wrap
func Panic(format string, a ...interface{}) {
	log.Panic(fmt.Sprintf(format, a))
}
