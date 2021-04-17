package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.com/Zendden/workshop/internal/api/joke"
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

	apiClient := joke.NewClient(cfg.JokeURL)

	/* Handler */
	h := handler.NewHandler(apiClient)

	/* Router */
	r := chi.NewRouter()

	/* Handle routes */
	r.Get("/", h.Index)

	path := cfg.Host+":"+cfg.Port

	/* Start server */
	log.Printf("Server started as %s", path)
	err = http.ListenAndServe(path, r)
	log.Fatal(err)
}