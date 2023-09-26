package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(NewContextHandler(slog.LevelInfo))
	s, err := NewServer(logger, 3000)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	if err := s.Run(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
