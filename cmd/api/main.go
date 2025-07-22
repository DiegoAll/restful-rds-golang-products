package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restful-rds-golang-products/database"
	"restful-rds-golang-products/internal/pkg/logger"
	"restful-rds-golang-products/internal/repository"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *database.DB
	// middleware *middleware.Middleware
	productRepo *repository.ProductsRepository
	apiKeyRepo  *repository.APIKeyRepository
}

func main() {

	logger.Init()

	var cfg config
	cfg.port = 9090

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "9090"
	}

	// dsn := os.Getenv("POSTGRES_DSN")
	// if dsn == "" {
	// 	// Actualizamos el mensaje de error
	// 	log.Fatal("Error: La variable de entorno POSTGRES_DSN no está definida.")
	// }
	// log.Printf("DEBUG: POSTGRES_DSN obtenido del entorno: %s", dsn)

	var dsn string
	rdsEnabled := os.Getenv("RDS_ENABLED")

	if rdsEnabled == "true" {
		// Si RDS_ENABLED es "true", usa la cadena de conexión de RDS
		dsn = os.Getenv("POSTGRES_RDS_DSN")
		if dsn == "" {
			log.Fatal("Error: POSTGRES_RDS_DSN environment variable is not defined when RDS_ENABLED is 'true'.")
		}
		logger.InfoLog.Printf("Connecting to RDS PostgreSQL using DSN (host and dbname shown for debug): %s", dsn)
	} else {
		// De lo contrario, usa la cadena de conexión local
		dsn = os.Getenv("POSTGRES_LOCAL_DSN")
		if dsn == "" {
			// Si no se define ninguna de las dos, es un error fatal
			log.Fatal("Error: POSTGRES_LOCAL_DSN environment variable is not defined when RDS_ENABLED is not 'true'.")
		}
		logger.InfoLog.Printf("Connecting to Local PostgreSQL using DSN (host and dbname shown for debug): %s", dsn)
	}

	dbInstance, err := database.ConnectPostgres(dsn)
	if err != nil {
		log.Fatalf("Cannot connect to PostgreSQL: %v", err)
	}

	// Inicializar los repositorios
	productRepo := repository.NewPostgresProductRepository(dbInstance.SQL)
	apiKeyRepo := repository.NewAPIKeyRepository(dbInstance.SQL)

	app := &application{
		config:      cfg,
		infoLog:     logger.InfoLog,
		errorLog:    logger.ErrorLog,
		db:          dbInstance,
		productRepo: productRepo,
		apiKeyRepo:  apiKeyRepo,
		// middleware: mw,
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)

	// Aca esta el tema del TLS handler
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
