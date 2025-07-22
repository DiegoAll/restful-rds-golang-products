package main

import (
	"context"
	"fmt"
	"net/http"
	"restful-rds-golang-products/internal/pkg/utils"

	"golang.org/x/oauth2"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type envelope map[string]interface{}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	// Prepara la respuesta JSON
	payload := jsonResponse{
		Error:   false,
		Message: "Service is healthy",
		Data:    envelope{"status": "ok", "version": "1.0.0"}, // Puedes añadir más información aquí
	}
	_ = utils.WriteJSON(w, http.StatusOK, payload)
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (app *application) Callback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	code := r.URL.Query().Get("code")
	oauth2Token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "No se pudo intercambiar el código por token", http.StatusInternalServerError)
		return
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No se encontró el id_token", http.StatusInternalServerError)
		return
	}

	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	fmt.Println("✅ ID Token:", rawIDToken)
	fmt.Println("✅ Access Token:", oauth2Token.AccessToken)

	var claims struct {
		Email string `json:"email"`
	}

	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, "No se pudieron obtener los claims", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Autenticación exitosa. Bienvenido: %s", claims.Email)
}

func (app *application) Protected(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Falta el header Authorization", http.StatusUnauthorized)
		return
	}

	// El formato del header debe ser: "Bearer <token>"
	const prefix = "Bearer "
	if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
		http.Error(w, "Formato del token inválido", http.StatusUnauthorized)
		return
	}

	rawIDToken := authHeader[len(prefix):]

	ctx := r.Context()
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	var claims struct {
		Email string `json:"email"`
	}

	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, "No se pudieron obtener los claims", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Acceso concedido al usuario: %s", claims.Email)

}
