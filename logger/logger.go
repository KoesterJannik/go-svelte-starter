package logger

import (
	"fmt"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() {
	var handler slog.Handler
	val := os.Getenv("ENV")
	if val == "" {
		panic("Environment variable ENV not set.")
	}
	fmt.Println("ENV:", val)
	if os.Getenv("ENV") == "dev" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		handler = slog.NewJSONHandler(file, nil)
	}
	Logger = slog.New(handler)
	Logger.Info("logger initialized")
}
