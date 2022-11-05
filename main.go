package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kelseyhightower/envconfig"

	"github.com/Looty/go-api/internal"
)

func main() {
	r := chi.NewRouter()

	var cfg internal.Config
	err := envconfig.Process("go-api", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome2"))
	})

	r.Get("/healtz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(os.Getenv("APP_VERSION")))
	})

	port := fmt.Sprintf(":%s", cfg.Port)
	http.ListenAndServe(port, r)
}
