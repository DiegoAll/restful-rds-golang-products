package repository

import (
	"context"
	"database/sql"
	"restful-rds-golang-products/internal/pkg/logger"
	"time"
)

// APIKeyRepository maneja las operaciones de base de datos para las API Keys.
type APIKeyRepository struct {
	DB *sql.DB
}

// NewAPIKeyRepository crea una nueva instancia de APIKeyRepository.
func NewAPIKeyRepository(db *sql.DB) *APIKeyRepository {
	return &APIKeyRepository{DB: db}
}

// ValidateAPIKey verifica si una API Key es válida y está activa en la base de datos.
func (r *APIKeyRepository) ValidateAPIKey(ctx context.Context, apiKey string) (bool, error) {
	query := `
        SELECT expires_at, active
        FROM api_keys
        WHERE key = $1`

	var expiresAt sql.NullTime
	var active bool

	err := r.DB.QueryRowContext(ctx, query, apiKey).Scan(&expiresAt, &active)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.DebugLog.Printf("API Key no encontrada: %s", apiKey)
			return false, nil // La API Key no existe
		}
		logger.ErrorLog.Printf("Error de base de datos al validar API Key: %v", err)
		return false, err
	}

	// Verificar si la API Key está activa y no ha expirado
	if !active {
		logger.DebugLog.Printf("API Key inactiva: %s", apiKey)
		return false, nil
	}

	if expiresAt.Valid && expiresAt.Time.Before(time.Now()) {
		logger.DebugLog.Printf("API Key expirada: %s", apiKey)
		return false, nil // La API Key ha expirado
	}

	logger.InfoLog.Printf("API Key validada exitosamente: %s", apiKey)
	return true, nil
}
