package handler

import (
	"net/http"
	"time"

	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/configuration"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var _ Handler = (*handler)(nil)

type (
	Handler http.Handler

	handler struct {
		config configuration.Config
		router http.Handler
	}
)

func NewHandler(
	config configuration.Config,
	healthcheck Healthcheck,
	slack SlackHandler,
) Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Mount("/healthcheck", healthcheck)
	router.Mount("/slack", slack)

	return &handler{
		config: config,
		router: router,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
