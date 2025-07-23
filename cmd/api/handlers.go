package main

import (
	"context"
	"net/http"
	"restful-rds-golang-products/internal/pkg/logger"
	"restful-rds-golang-products/internal/pkg/utils"
	"restful-rds-golang-products/models"
	"time"

	"github.com/aws/aws-sdk-go/aws"
)

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

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	// Prepara la respuesta JSON
	payload := jsonResponse{
		Error:   false,
		Message: "Service is healthy",
		Data:    envelope{"status": "ok", "version": "1.0.0"}, // Puedes añadir más información aquí
	}
	_ = utils.WriteJSON(w, http.StatusOK, payload)
}

func (app *application) confirm(w http.ResponseWriter, r *http.Request) {

	var confirm ConfirmRequest

	err := utils.ReadJSON(w, r, &confirm)
	if err != nil {
		logger.ErrorLog.Printf("Error al leer JSON en CreateProduct: %v", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_, err = cognitoClient.ConfirmSignUp(context.TODO(), &cognito.ConfirmSignUpInput{
		ClientId:         aws.String(ClientID),
		Username:         aws.String(confirm.Email),
		ConfirmationCode: aws.String(confirm.Code),
	})
	if err != nil {
		http.Error(w, "Error confirmando usuario: "+err.Error(), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Confirmacion exitosa",
		Data:    "",
	}

	err = utils.WriteJSON(w, http.StatusCreated, payload)
	if err != nil {
		logger.ErrorLog.Printf("Error al escribir respuesta JSON en CreateProduct: %v", err)
		// En este punto, no hay mucho que se pueda hacer más que loguear el error
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
		logger.ErrorLog.Printf("Error al insertar producto en la base de datos: %v", err)
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
