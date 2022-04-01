// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/configuration"
	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/handler"
	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/server"
	"net/http"
)

// Injectors from inject.go:

func initialize() *http.Server {
	config := configuration.Get()
	healthcheck := handler.NewHealthcheck(config)
	httpHandler := handler.NewHandler(config, healthcheck)
	httpServer := server.NewServer(config, httpHandler)
	return httpServer
}
