package logger

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf)

	ctx := context.WithValue(context.Background(), RequestIDKey, "12345")
	fields := map[string]interface{}{"key1": "value1"}

	err := logger.Info(ctx, "This is an info message", fields)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "INFO: This is an info message key1=value1 requestID=12345"
	if buf.String()[:len(expected)] != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}

func TestLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf)

	ctx := context.WithValue(context.Background(), UserIDKey, "user123")
	fields := map[string]interface{}{"key2": "value2"}
	err := logger.Error(ctx, errors.New("an error occurred"), fields)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "ERROR: an error occurred key2=value2 userID=user123"
	if buf.String()[:len(expected)] != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}

func TestLogger_Debug_Skipped(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf)

	logger.SetLogLevel(LevelInfo) // Set log level to INFO
	ctx := context.Background()
	fields := map[string]interface{}{"key3": "value3"}

	err := logger.Debug(ctx, "This debug message should be skipped", fields)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if buf.Len() != 0 {
		t.Errorf("expected no output, got %q", buf.String())
	}
}

func TestLogger_Warn(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf)

	ctx := context.Background()
	fields := map[string]interface{}{"key4": "value4"}

	err := logger.Warn(ctx, "This is a warning message", fields)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "WARN: This is a warning message key4=value4"
	if buf.String()[:len(expected)] != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}

func TestLogger_Fatal(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf)

	ctx := context.Background()
	fields := map[string]interface{}{"key5": "value5"}

	err := logger.Fatal(ctx, "This is a fatal message", fields)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "FATAL: This is a fatal message key5=value5"
	if buf.String()[:len(expected)] != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}
