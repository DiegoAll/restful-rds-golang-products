package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"restful-rds-golang-products/database"
	"restful-rds-golang-products/internal/pkg/logger"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *database.DB
	//middleware
}

var (
	clientID    = "89960475367-37h7e26id256t33v5b7p5aho4kbv0gio.apps.googleusercontent.com"
	oauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
		Endpoint:     google.Endpoint,
	}
	verifier *oidc.IDTokenVerifier
)

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

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatalf("No se pudo crear el proveedor OIDC: %v", err)
	}

	verifier = provider.Verifier(&oidc.Config{ClientID: clientID})

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

	app := &application{
		config:   cfg,
		infoLog:  logger.InfoLog,
		errorLog: logger.ErrorLog,
		db:       dbInstance,
		//middleware:         mw,
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
