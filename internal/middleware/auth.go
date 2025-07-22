package middleware

import (
	"context"
	"errors"
	"net/http"
	"restful-rds-golang-products/internal/pkg/logger"
	"restful-rds-golang-products/internal/pkg/utils"
	"restful-rds-golang-products/internal/repository"
)

// AuthAPIKeyMiddleware es un middleware que valida la API Key.
// Espera la API Key en el encabezado "X-API-Key".
func AuthAPIKeyMiddleware(apiKeyRepo *repository.APIKeyRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")

			if apiKey == "" {
				logger.ErrorLog.Println("API Key no proporcionada en la solicitud")
				utils.ErrorJSON(w, errors.New("API Key no proporcionada"), http.StatusUnauthorized)
				return
			}

			// Validar la API Key en la base de datos
			isValid, err := apiKeyRepo.ValidateAPIKey(context.Background(), apiKey)
			if err != nil {
				logger.ErrorLog.Printf("Error al validar la API Key: %v", err)
				utils.ErrorJSON(w, errors.New("error interno del servidor al validar API Key"), http.StatusInternalServerError)
				return
			}

			if !isValid {
				logger.ErrorLog.Printf("API Key inválida o inactiva: %s", apiKey)
				utils.ErrorJSON(w, errors.New("API Key inválida o inactiva"), http.StatusUnauthorized)
				return
			}

			// Si la API Key es válida, pasa la solicitud al siguiente handler
			next.ServeHTTP(w, r)
		})
	}
}
