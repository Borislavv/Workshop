package handler

import (
	"fmt"
	"gitlab.com/Zendden/workshop/internal/api"
	"net/http"
)

type Handler struct {
	jokeClient api.Client
	customeJoke string
}

func NewHandler(jokeClient api.Client, customeJoke string) *Handler {
	return &Handler{
		jokeClient: jokeClient,
		customeJoke: customeJoke,
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if h.customeJoke != "" {
		fmt.Fprint(w, h.customeJoke)
		return
	}

	joke, err := h.jokeClient.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, joke.Joke)
}