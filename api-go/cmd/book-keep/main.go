package main

import (
	bookkeep "BookKeep/internal/app/book-keep"
	"log/slog"
	"os"
)

func main() {
	s := bookkeep.NewBookKeep()

	if err := s.Start(); err != nil {
		slog.Error("failed to start book keep", "error", err)
		os.Exit(1)
	}
}
