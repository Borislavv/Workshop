package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.com/Zendden/workshop/internal/config"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"gitlab.com/Zendden/workshop/internal/handler"
)

/* App entrypoint */
func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	/* Handler */
	h := handler.NewHandler()

	/* Router */
	r := chi.NewRouter()

	/* Handle routes */
	r.Get("/", h.Index)

	/* Start server */
	log.Print("Server started")
	err = http.ListenAndServe(":8080", r);
	log.Fatal(err)
}