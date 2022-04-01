package handler

import (
	"net/http"

	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/configuration"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var _ Healthcheck = (*healthcheck)(nil)

type (
	Healthcheck http.Handler

	healthcheck struct {
		config configuration.Config
		router http.Handler
	}
)

func NewHealthcheck(
	config configuration.Config,
) Healthcheck {
	router := chi.NewRouter()

	router.Get("/", get)

	return &healthcheck{
		config: config,
		router: router,
	}
}

func (h *healthcheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

type (
	HealthcheckResponse struct {
		Message string `json:"message"`
	}
)

func get(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, HealthcheckResponse{
		Message: "ok",
	})
}
