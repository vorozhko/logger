package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	// create a logger that writes to console
	consoleLogger := NewLogger(os.Stdout)
	consoleLogger.SetLogLevel(LevelDebug)

	// create a logger that writes to file
	file, err := os.Create("app.log")
	if err != nil {
		fmt.Print("Failed to create log file:", err)
	}
	defer file.Close()
	fileLogger := NewLogger(file)

	ctx := context.WithValue(context.Background(), RequestIDKey, "12345")
	ctx = context.WithValue(ctx, UserIDKey, "67890")

	consoleLogger.Info(ctx, "User logged in", map[string]interface{}{"svc": "auth"})
	fileLogger.Error(ctx, fmt.Errorf("database connection failed"), map[string]interface{}{"svc": "db"})
	fileLogger.Debug(ctx, "Some debug output", nil)
}
