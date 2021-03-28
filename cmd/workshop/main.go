package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"gitlab.com/Zendden/workshop/internal/handler"
)

/* App entrypoint */
func main() {
	/* Handler */
	h := handler.NewHandler()

	/* Handle routes */
	r := chi.NewRouter()
	r.Get("/", h.Index)

	/* Start server */
	log.Print("Server started")
	err := http.ListenAndServe(":8080", r);
	log.Fatal(err)
}