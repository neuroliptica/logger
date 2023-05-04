// logger: simple centralized logging;
//
// example usage:
//
// localLogger := MakeLogger("local-logger").BindToDefault()
// localLogger.Log("Ein zwei drei")
//	=> [local-logger] Ein zwer drei
//
// str := "some string"
// localLogger.Logf("This is %s", str)
//	=> [local-logger] This is some string

package logger

import (
	"fmt"
	"log"
)

// Global exported channel for synced logging.
var GlobalLogger = make(chan string)

// Initalize infinity loop in separated goroutine for logging in stdout.
func init() {
	go func() {
		for msg := range GlobalLogger {
			log.Println(msg)
		}
	}()

	log.SetFlags(log.Ltime)
	//initLogger := MakeLogger("logger-init").BindToDefault()
	//initLogger.Log("initializing global logger done.")
}

// Default logger. LoggerType is for prefix, Destination is for binded channel.
type Logger struct {
	Destination chan string
	LoggerType  string
}

// Create Logger instance with t prefix, no channel binded.
func MakeLogger(t string) *Logger {
	return &Logger{
		LoggerType: t,
	}
}

// Bind Logger to reading channel.
func (l *Logger) BindToChannel(ch chan string) *Logger {
	l.Destination = GlobalLogger
	return l
}

// Bind Logger to default global channel.
func (l *Logger) BindToDefault() *Logger {
	return l.BindToChannel(GlobalLogger)
}

// Write anything to destination, can be unsafe, if msg isn't a Stringer instance.
func (l *Logger) Log(msg ...interface{}) {
	l.Destination <- fmt.Sprintf("[%s] %3s",
		l.LoggerType, fmt.Sprint(msg...))
}

// Write anything to destination with formating, can be unsafe.
func (l *Logger) Logf(format string, msg ...interface{}) {
	l.Log(fmt.Sprintf(format, msg...))
}
