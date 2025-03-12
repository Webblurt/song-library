package utils

import (
	"log"
	"os"
	"strconv"
)

type Logger struct {
	*log.Logger
	debugEnabled bool
}

func bgColor(logType string) string {
	switch logType {
	case "error":
		return "\033[41m" //red
	case "info":
		return "\033[44m" //blue
	case "debug":
		return "\033[42m" //green
	case "warn":
		return "\033[43m" //yellow
	case "success":
		return "\033[45m" // magenta
	case "trace":
		return "\033[46m" // cyan
	default:
		return "\033[0m" //reset
	}
}

func NewLogger(prefix, debug string) *Logger {
	debugEnabled, err := strconv.ParseBool(debug)
	if err != nil {
		debugEnabled = false
	}
	return &Logger{
		Logger:       log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile),
		debugEnabled: debugEnabled,
	}
}

func (l *Logger) Info(v ...interface{}) {
	color := bgColor("info")
	reset := bgColor("reset")
	l.SetPrefix(color + "INFO: " + reset)
	l.Println(v...)
}

func (l *Logger) Debug(v ...interface{}) {
	if !l.debugEnabled {
		return
	}
	color := bgColor("debug")
	reset := bgColor("reset")
	l.SetPrefix(color + "DEBUG: " + reset)
	l.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	color := bgColor("warn")
	reset := bgColor("reset")
	l.SetPrefix(color + "WARN: " + reset)
	l.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	color := bgColor("error")
	reset := bgColor("reset")
	l.SetPrefix(color + "ERROR: " + reset)
	l.Println(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	color := bgColor("error")
	reset := bgColor("reset")
	l.SetPrefix(color + "FATAL: " + reset)
	l.Println(v...)
	os.Exit(1)
}

func (l *Logger) Trace(v ...interface{}) {
	color := bgColor("trace")
	reset := bgColor("reset")
	l.SetPrefix(color + "TRACE: " + reset)
	l.Println(v...)
}

func (l *Logger) Success(v ...interface{}) {
	color := bgColor("success")
	reset := bgColor("reset")
	l.SetPrefix(color + "SUCCESS: " + reset)
	l.Println(v...)
}

func (l *Logger) Custom(logType string, v ...interface{}) {
	color := bgColor(logType)
	reset := bgColor("reset")
	l.SetPrefix(color + logType + ": " + reset)
	l.Println(v...)
}
