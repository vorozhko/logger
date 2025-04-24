package main

import (
	"context"
	"fmt"
	"io"
)

var currentLogLevel LogLevel = LevelInfo

type contextKey string

const (
	RequestIDKey contextKey = "requestID"
	UserIDKey    contextKey = "UserID"
)

type LogLevel int

const (
	LevelDebug LogLevel = 0 // debug
	LevelInfo  LogLevel = 1 // info
	LevelWarn  LogLevel = 2 // warn
	LevelError LogLevel = 3 // error
	LevelFatal LogLevel = 4 // fatal

)

type Logging interface {
	Info(ctx context.Context, message string, fields map[string]interface{}) error
	Error(ctx context.Context, err error, fields map[string]interface{}) error
	Debug(ctx context.Context, message string, fields map[string]interface{}) error
	Warn(ctx context.Context, message string, fields map[string]interface{}) error
	Fatal(ctx context.Context, message string, fields map[string]interface{}) error
}

type Logger struct {
	destination io.Writer
}

func NewLogger(destination io.Writer) *Logger {
	return &Logger{destination: destination}
}

func (l Logger) SetLogLevel(level LogLevel) {
	currentLogLevel = level
}

func logLevelToString(level LogLevel) string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	default:
		return "UNKNOWN"
	}
}

func (l *Logger) log(logLevel LogLevel, ctx context.Context, message string, fields map[string]interface{}) error {
	if logLevel < currentLogLevel {
		return nil // Skip logging if the level is below the current log level
	}
	logMessage := fmt.Sprintf("%s: %s", logLevelToString(logLevel), message)

	// Add fields (structured logging)
	for k, v := range fields {
		logMessage += fmt.Sprintf(" %s=%v", k, v)
	}

	// Extract metadata from context
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		logMessage += fmt.Sprintf(" requestID=%v", requestID)
	}
	if userID := ctx.Value(UserIDKey); userID != nil {
		logMessage += fmt.Sprintf(" userID=%v", userID)
	}

	// End the log line
	_, err := fmt.Fprintln(l.destination, logMessage)

	// no errors for now
	return err
}

func (l Logger) Info(ctx context.Context, message string, fields map[string]interface{}) error {
	return l.log(LevelInfo, ctx, message, fields)
}

func (l Logger) Warn(ctx context.Context, message string, fields map[string]interface{}) error {
	return l.log(LevelWarn, ctx, message, fields)
}

func (l Logger) Fatal(ctx context.Context, message string, fields map[string]interface{}) error {
	return l.log(LevelFatal, ctx, message, fields)
}

func (l Logger) Debug(ctx context.Context, message string, fields map[string]interface{}) error {
	return l.log(LevelDebug, ctx, message, fields)
}

func (l Logger) Error(ctx context.Context, err error, fields map[string]interface{}) error {
	return l.log(LevelError, ctx, err.Error(), fields)
}
