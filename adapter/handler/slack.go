package handler

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Pranc1ngPegasus/slack-api-practice/adapter/configuration"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/slack-go/slack"
)

var _ SlackHandler = (*slackHandler)(nil)

type (
	SlackHandler http.Handler

	slackHandler struct {
		config configuration.Config
		router http.Handler
	}
)

func NewSlackHandler(
	config configuration.Config,
) SlackHandler {
	router := chi.NewRouter()
	h := &slackHandler{
		config: config,
		router: router,
	}

	router.Use(h.verifySlackToken)

	router.Post("/", post)

	return h
}

func (h *slackHandler) verifySlackToken(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		sv, err := slack.NewSecretsVerifier(r.Header, h.config.Slack.SigningSecret)
		if err != nil {
			errorResponse(w, r, err)
			return
		}

		r.Body = ioutil.NopCloser(io.TeeReader(r.Body, &sv))

		if err := sv.Ensure(); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			errorResponse(w, r, err)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
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
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse(w, r, err)
		return
	}

	render.JSON(w, r, SlashCommandResponse{
		ResponseType: "in_channel",
		Blocks: []SlashCommandResponseBlock{
			{
				Type: "section",
				Text: SlashCommandResponseBlockText{
					Type: "mrkdwn",
					Text: s.Text,
				},
			},
		},
	})
}

func errorResponse(w http.ResponseWriter, r *http.Request, err error) {
	render.JSON(w, r, SlashCommandResponse{
		ResponseType: "ephemeral",
		Blocks: []SlashCommandResponseBlock{
			{
				Type: "section",
				Text: SlashCommandResponseBlockText{
					Type: "mrkdwn",
					Text: err.Error(),
				},
			},
		},
	})
}
