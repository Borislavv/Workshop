package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"gitlab.com/Zendden/workshop/internal/handler"
)

func main() {
	h := handler.NewHandler()

	r := chi.NewRouter()
	r.Get("/", h.Index)

	err := http.ListenAndServe(":8080", r);
	log.Fatal(err)
}