package main

import (
	"net/http"
	"os"

	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/configuration"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	configuration.Load()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	server := initialize()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error().Err(err).Msg("Failed to exec server.")
		os.Exit(1)
	}
}
