package pkg

import (
	"log/slog"
	"os"
)

func LoggerInit() {
	env := os.Getenv("APP_ENV")
	logLevel := os.Getenv("APP_LOGLEVEL")

	level := slog.LevelInfo
	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	case "fatal":
		level = slog.LevelError
	}

	var handler slog.Handler
	if env == "local" {
		err := os.MkdirAll("../logs", os.ModePerm)
		if err != nil {
			panic("Failed to create log directory")
		}

		file, err := os.OpenFile("../logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("Failed to log to file, using default stderr")
		}
		handler = slog.NewTextHandler(file, &slog.HandlerOptions{
			Level: level,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Info("Logger initialized")
}
