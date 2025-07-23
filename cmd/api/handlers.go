package main

import (
	"context"
	"errors"
	"net/http"
	"restful-rds-golang-products/internal/pkg/logger"
	"restful-rds-golang-products/internal/pkg/utils"
	"restful-rds-golang-products/models"
	"time"
	// Importar tipos de Cognito
	// cognitoTypes "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type envelope map[string]interface{}

type ConfirmRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type ConfirmSignUpRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	// Prepara la respuesta JSON
	payload := jsonResponse{
		Error:   false,
		Message: "Service is healthy",
		Data:    envelope{"status": "ok", "version": "1.0.0"}, // Puedes añadir más información aquí
	}
	_ = utils.WriteJSON(w, http.StatusOK, payload)
}

// signup
func (app *application) SignUp(w http.ResponseWriter, r *http.Request) {
	var input SignUpRequest

	err := utils.ReadJSON(w, r, &input)
	if err != nil {
		logger.ErrorLog.Printf("Error reading JSON for signup: %v", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if input.Email == "" || input.Password == "" {
		utils.ErrorJSON(w, errors.New("correo electrónico y contraseña son requeridos"), http.StatusBadRequest)
		return
	}

	err = app.cognitoAuth.SignUp(r.Context(), input.Email, input.Password)
	if err != nil {
		logger.ErrorLog.Printf("Error during Cognito SignUp for %s: %v", input.Email, err)
		utils.ErrorJSON(w, err)
		return
	}

	payload := utils.JSONResponse{
		Error:   false,
		Message: "Registro exitoso. Por favor, confirma tu cuenta con el código enviado a tu correo.",
	}

	err = utils.WriteJSON(w, http.StatusAccepted, payload) // 202 Accepted indica que la solicitud fue aceptada para procesamiento
	if err != nil {
		logger.ErrorLog.Printf("Error writing JSON response for signup: %v", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}

func (app *application) ConfirmSignUp(w http.ResponseWriter, r *http.Request) {
	var confirm ConfirmSignUpRequest

	err := utils.ReadJSON(w, r, &confirm)
	if err != nil {
		logger.ErrorLog.Printf("Error reading JSON for confirm signup: %v", err)
		utils.ErrorJSON(w, err)
		return
	}

	if confirm.Email == "" || confirm.Code == "" {
		utils.ErrorJSON(w, errors.New("correo electrónico y código de confirmación son requeridos"), http.StatusBadRequest)
		return
	}

	err = app.cognitoAuth.ConfirmSignUp(r.Context(), confirm.Email, confirm.Code)
	if err != nil {
		logger.ErrorLog.Printf("Error during Cognito ConfirmSignUp for %s: %v", confirm.Email, err)
		utils.ErrorJSON(w, err)
		return
	}

	payload := utils.JSONResponse{
		Error:   false,
		Message: "Confirmación de cuenta exitosa. Ya puedes iniciar sesión.",
	}

	err = utils.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		logger.ErrorLog.Printf("Error writing JSON response for confirm signup: %v", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
	}
}

// Login
func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	var login LoginRequest

	err := utils.ReadJSON(w, r, &login)
	if err != nil {
		logger.ErrorLog.Printf("Error al leer JSON en Login: %v", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if login.Email == "" || login.Password == "" {
		utils.ErrorJSON(w, errors.New("correo electrónico y contraseña son requeridos"), http.StatusBadRequest)
		return
	}

	authResult, err := app.cognitoAuth.SignIn(r.Context(), login.Email, login.Password)
	if err != nil {
		logger.ErrorLog.Printf("Error during Cognito SignIn for %s: %v", login.Email, err)
		utils.ErrorJSON(w, err)
		return
	}

	payload := utils.JSONResponse{
		Error:   false,
		Message: "Inicio de sesión exitoso.",
		Data:    authResult, // Devuelve los tokens de Cognito
	}

	err = utils.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		logger.ErrorLog.Printf("Error writing JSON response for signin: %v", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
	}

}

func (app *application) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := utils.ReadJSON(w, r, &product)
	if err != nil {
		logger.ErrorLog.Printf("Error al leer JSON en CreateProduct: %v", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Asignar timestamps (aunque el repositorio también podría hacerlo, es bueno tenerlo aquí para el modelo)
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second) // Contexto con timeout
	defer cancel()

	err = app.productRepo.InsertProduct(ctx, &product)
	if err != nil {
		logger.ErrorLog.Printf("Error inserting product: %v", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Producto creado exitosamente",
		Data:    product,
	}

	err = utils.WriteJSON(w, http.StatusCreated, payload)
	if err != nil {
		logger.ErrorLog.Printf("Error al escribir respuesta JSON en CreateProduct: %v", err)
		// En este punto, no hay mucho que se pueda hacer más que loguear el error
	}
}

// List Product
func (app *application) AllProducts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second) // Contexto con timeout
	defer cancel()

	products, err := app.productRepo.GetAllProducts(ctx)
	if err != nil {
		logger.ErrorLog.Printf("Error al obtener todos los productos de la base de datos: %v", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Productos obtenidos exitosamente",
		Data:    products,
	}

	err = utils.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		logger.ErrorLog.Printf("Error al escribir respuesta JSON en AllProducts: %v", err)
		// En este punto, no hay mucho que se pueda hacer más que loguear el error
	}
}
