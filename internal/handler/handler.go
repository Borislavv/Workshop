package handler

import (
	"fmt"
	"gitlab.com/Zendden/workshop/internal/api"
	"net/http"
)

type Handler struct {
	jokeClient api.Client
}

func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		jokeClient: jokeClient,
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	joke, err := h.jokeClient.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, joke.Joke)
}