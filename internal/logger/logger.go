package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	Namespace string
}

func (l *Logger) Log(message string) {
	now := time.Now().Format(time.RFC3339)
	fmt.Printf("%s::%s::%s::%s\n", now, "INFO", l.Namespace, message)
}

func (l *Logger) LogDebug(message string) {
	now := time.Now().Format(time.RFC3339)
	fmt.Printf("%s::%s::%s::%s\n", now, "DEBUG", l.Namespace, message)
}

func (l *Logger) LogError(message string) {
	now := time.Now().Format(time.RFC3339)
	fmt.Printf("%s::%s::%s::%s\n", now, "ERROR", l.Namespace, message)
}

var RootLogger = Logger{
	Namespace: "ROOT",
}
