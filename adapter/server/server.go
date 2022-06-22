package server

import (
	"net/http"

	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/handler"

	"github.com/rs/zerolog/log"
)

func NewServer(
	config configuration.Config,
	handler handler.Handler,
) *http.Server {
	port := config.ListenPort
	log.Info().Msgf("Listen at :%s", port)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return server
}
