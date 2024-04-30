package logger

import (
	"log/slog"
)

// TODO: implement logging
func init() {
	logger := slog.Default()
  slog.SetDefault(logger)

	slog.Info("logger initialised")
}
