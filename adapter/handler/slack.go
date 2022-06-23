package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var _ SlackHandler = (*slackHandler)(nil)

type (
	SlackHandler http.Handler

	slackHandler struct {
		router http.Handler
	}
)

func NewSlackHandler() SlackHandler {
	router := chi.NewRouter()

	router.Post("/", post)

	return &slackHandler{
		router: router,
	}
}

func (h *slackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

type (
	SlashCommandRequest struct {
		Token       string `form:"token"`
		Command     string `form:"command"`
		Text        string `form:"text"`
		ResponseUrl string `form:"response_url"`
		TriggerId   string `form:"trigger_id"`
		UserId      string `form:"user_id"`
		UserName    string `form:"user_name"`
		ApiAppId    string `form:"api_app_id"`
	}

	SlashCommandResponse struct {
		ResponseType string                      `json:"response_type"`
		Blocks       []SlashCommandResponseBlock `json:"blocks"`
	}

	SlashCommandResponseBlock struct {
		Type string                        `json:"type"`
		Text SlashCommandResponseBlockText `json:"text"`
	}

	SlashCommandResponseBlockText struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}
)

func post(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	render.JSON(w, r, SlashCommandResponse{
		ResponseType: "in_channel",
		Blocks: []SlashCommandResponseBlock{
			{
				Type: "section",
				Text: SlashCommandResponseBlockText{
					Type: "mrkdwn",
					Text: r.Form.Get("text"),
				},
			},
		},
	})
}
