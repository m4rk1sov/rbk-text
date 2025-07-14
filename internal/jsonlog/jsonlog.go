package jsonlog

import (
	"encoding/json"
	"io"
	"runtime/debug"
	"sync"
	"time"
)

type Level int8

const (
	LevelTrace Level = iota
	LevelInfo
	LevelError
	LevelFatal
	LevelOff
)

// human friendly string for severity
func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

func (l *Logger) PrintTrace(message string, properties map[string]string) {
	_, err := l.print(LevelTrace, message, properties)
	if err != nil {
		return
	}
}

func (l *Logger) PrintInfo(message string, properties map[string]string) {
	_, err := l.print(LevelInfo, message, properties)
	if err != nil {
		return
	}
}

func (l *Logger) PrintError(message string, properties map[string]string) {
	_, err := l.print(LevelError, message, properties)
	if err != nil {
		return
	}
}

func (l *Logger) PrintFatal(message string, properties map[string]string) {
	_, err := l.print(LevelFatal, message, properties)
	if err != nil {
		return
	}
}

func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	// if the level is below the minimum severity, return
	if level < l.minLevel {
		return 0, nil
	}

	// anonymous struct for holding the log entry
	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().Local().Format(time.RFC3339),
		Message:    message,
		Properties: properties,
	}
	// stack for levels above ERROR
	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	// line for holding the log entry text
	var line []byte

	// Marshal log struct text into JSON, if fails, return text error
	line, err := json.Marshal(aux)
	if err != nil {
		line = []byte(LevelError.String() + ": unable to marshal log message: " + err.Error())
	}

	// locking the mutex so no concurrent writes in one
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.out.Write(append(line, '\n'))
}

// implement the Write method to satisfy io.Writer interface
func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(LevelError, string(message), nil)
}
