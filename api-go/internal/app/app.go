package app

import (
	"BookKeep/internal/config"
	"log/slog"
)

func Run() {
	_, err := config.NewConfig("config/local.yaml")
	if err != nil {
		slog.Error("failed to load config", "error", err)
	}

	
}