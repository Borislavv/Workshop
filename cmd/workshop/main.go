package main

import (
	"context"
	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.com/Zendden/workshop/internal/api/joke"
	"gitlab.com/Zendden/workshop/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	h := handler.NewHandler(apiClient, cfg.CustomeJoke)

	/* Router */
	r := chi.NewRouter()

	/* Handle routes */
	r.Get("/", h.Index)

	path := cfg.Host+":"+cfg.Port

	server := &http.Server{
		Addr: path,
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<- quit
		ctx, _ := context.WithTimeout(context.Background(), time.Second * 15)
		done <- server.Shutdown(ctx)
	}()

	/* Start server */
	log.Printf("Server started as %s", path)
	_ = server.ListenAndServe()

	err = <- done

	log.Printf("shutting server down with: %v", err)
}