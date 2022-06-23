//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/handler"
	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/server"

	"github.com/google/wire"
)

func initialize() *http.Server {
	wire.Build(
		configuration.Get,

		handler.NewHealthcheck,
		handler.NewSlackHandler,
		handler.NewHandler,

		server.NewServer,
	)

	return nil
}
