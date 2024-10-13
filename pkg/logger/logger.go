package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func LogInit() {
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	handler := slog.NewTextHandler(file, opts)
	Logger = slog.New(handler)
}
