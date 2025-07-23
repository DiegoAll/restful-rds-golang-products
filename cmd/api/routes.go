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
		// MIDDLEWARE APLICADO
		r.Get("/products", app.AllProducts) // Endpoint: GET /v1/products/

		r.Group(func(r chi.Router) {
			r.Use(customMiddleware.AuthAPIKeyMiddleware(app.apiKeyRepo))

			// Rutas de productos PROTEGIDAS por la API Key
			// La ruta base de este grupo es /v1/, así que /products/ aquí se convierte en /v1/products/
			r.Post("/products", app.CreateProduct) // Endpoint: POST /v1/products/
			// Si tuvieras otras rutas protegidas como GET /products/{id} o PUT /products/{id}, irían aquí:
			// r.Get("/products/{id}", app.GetProduct)
			// r.Put("/products/{id}", app.UpdateProduct)
		})

		// La línea 'mux.Post("/products", app.CreateProduct)' duplicada ha sido eliminada.
		// Las rutas restantes que no requieren autenticación se mantienen públicas.
		// mux.Get("/products/get/{id}", app.GetProduct)
		// mux.Put("/products/update/{id}", app.UpdateProduct)

	})

	return mux
}
