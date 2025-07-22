package main

import (
	"net/http"

	// Importa el middleware de Chi (sin alias, se usará como 'middleware')
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	// Importa tu middleware personalizado con un alias para evitar conflicto de nombres
	customMiddleware "restful-rds-golang-products/internal/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer) // Usa el paquete 'middleware' de Chi
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-API-Key"}, // Añadir X-API-Key
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Route("/v1", func(r chi.Router) {

		// Rutas públicas de la V1
		r.Get("/health", app.healthCheck)
		// La ruta AllProducts es pública y no requiere API Key
		r.Get("/products/all", app.AllProducts)

		r.Route("/products", func(r chi.Router) {
			// Aplica tu middleware de API Key personalizado a las rutas que lo requieran
			// Se usa el alias 'customMiddleware' para referirse a tu paquete.
			r.Use(customMiddleware.AuthAPIKeyMiddleware(app.apiKeyRepo))
			r.Post("/", app.CreateProduct) // Esta ruta ahora requiere API Key
			// r.Get("/", app.AllProducts) // Si esta ruta también requiriera autenticación, se dejaría aquí
		})

		// La línea 'mux.Post("/products", app.CreateProduct)' duplicada ha sido eliminada.
		// Las rutas restantes que no requieren autenticación se mantienen públicas.
		// mux.Get("/products/get/{id}", app.GetProduct)
		// mux.Put("/products/update/{id}", app.UpdateProduct)

	})

	return mux
}
