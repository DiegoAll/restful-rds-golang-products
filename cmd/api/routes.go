package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Route("/v1", func(r chi.Router) {

		// Rutas públicas de la V1

		// r.Get("/health", app.systemController.HealthCheck)
		// r.Post("/signup", app.authController.RegisterUser)
		// r.Post("/callback", app.authController.Callback)
		// r.Post("/login", app.authController.AuthenticateUser)

		// r.Post("/input", app.IngestData)

		// Receives an IP address and returns and enriched JSON.
		//r.Get("/enrichment", app.Enrichment)

		// Query log
		//r.Get("/statistics", app.Statistics)
		// *csv

		// r.Get("/health", app.healthCheck)
		//r.Get("/health", app.healthCheck)

	})

	return mux

}
