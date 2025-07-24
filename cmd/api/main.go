package main

import (
	"context"
	"encoding/json" // NECESARIO para unmarshal el secreto de Secrets Manager
	"fmt"
	"log"
	"net/http"
	"os"
	"restful-rds-golang-products/database"
	"restful-rds-golang-products/internal/pkg/cognito"
	"restful-rds-golang-products/internal/pkg/logger"
	"restful-rds-golang-products/internal/repository"

	"github.com/aws/aws-sdk-go-v2/aws"              // Correcto: Este es el paquete 'aws' de la V2
	awsconfig "github.com/aws/aws-sdk-go-v2/config" // Correcto: Este es el paquete 'config' de la V2
	cognitoservice "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager" // NECESARIO para Secrets Manager (V2)
)

type config struct {
	port int
}

type application struct {
	config      config
	infoLog     *log.Logger
	errorLog    *log.Logger
	db          *database.DB
	productRepo *repository.ProductsRepository
	cognitoAuth cognito.AuthClient
}

const (
	UserPoolID = "us-east-1_ZTzSnlG81"
	ClientID   = "2oa45rcrl66qophvccaeesdtl9"
	Region     = "us-east-1" // cambia según tu región
)

func main() {
	logger.Init()

	var cfg config
	cfg.port = 9090

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "9090"
	}

	var dsn string
	appEnv := os.Getenv("APP_ENV")

	// Declara sdkConfig aquí, una sola vez, para que esté disponible en todo el main.
	// El tipo debe ser de la V2: "github.com/aws/aws-sdk-go-v2/aws".Config
	var sdkConfig aws.Config

	switch appEnv {
	case "development_local_db":
		dsn = os.Getenv("POSTGRES_LOCAL_DSN")
		if dsn == "" {
			log.Fatal("Error: POSTGRES_LOCAL_DSN environment variable is not defined for 'development_local_db' environment.")
		}
		logger.InfoLog.Printf("Connecting to Local PostgreSQL using DSN: %s", dsn)

	case "development_remote_rds":
		dsn = os.Getenv("POSTGRES_RDS_DSN") // Se espera que venga del docker-compose.yml o env local
		if dsn == "" {
			log.Fatal("Error: POSTGRES_RDS_DSN environment variable is not defined for 'development_remote_rds' environment.")
		}
		logger.InfoLog.Printf("Connecting to Remote RDS for Development using DSN: %s", dsn)

	case "production", "aws": // Puedes usar "production" o "aws" como flag para despliegue en AWS
		// En este caso, el DSN SIEMPRE debe venir de AWS Secrets Manager
		rdsSecretName := os.Getenv("RDS_SECRET_NAME")
		if rdsSecretName == "" {
			log.Fatal("Error: RDS_SECRET_NAME environment variable is not defined for 'production' environment.")
		}

		// Asignamos a la sdkConfig declarada arriba, NO la redeclaramos con :=
		var err error // Declaramos 'err' localmente si no existe, o reusamos si ya existe.
		sdkConfig, err = awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion(Region))
		if err != nil {
			log.Fatalf("unable to load SDK config for secrets manager, %v", err)
		}

		// *** Lógica para obtener el DSN del secreto en Secrets Manager ***
		smClient := sm.NewFromConfig(sdkConfig)

		getSecretValueOutput, err := smClient.GetSecretValue(context.TODO(), &sm.GetSecretValueInput{
			SecretId: aws.String(rdsSecretName), // Se necesita aws.String de la V2
		})
		if err != nil {
			log.Fatalf("Failed to retrieve secret %s: %v", rdsSecretName, err)
		}

		if getSecretValueOutput.SecretString == nil {
			log.Fatal("Error: Secret string is nil for RDS DSN secret.")
		}

		// Deserializa el JSON del secreto (recuerda que lo guardamos como un par Key/Value)
		var secretMap map[string]string
		err = json.Unmarshal([]byte(*getSecretValueOutput.SecretString), &secretMap)
		if err != nil {
			log.Fatalf("Failed to parse secret JSON for RDS DSN: %v", err)
		}

		// Obtiene el DSN de la clave "FULL_RDS_DSN" que definiste en el secreto
		dsn, ok := secretMap["FULL_RDS_DSN"] // <--- ¡Asegúrate de que "FULL_RDS_DSN" sea la clave que usaste!
		if !ok || dsn == "" {
			log.Fatal("Error: 'FULL_RDS_DSN' key not found or empty in the secret.")
		}

		// Para logging, cuidado de no exponer la contraseña.
		// Podrías extraer host y dbname del secretMap si quieres más detalle.
		logger.InfoLog.Printf("Connecting to RDS PostgreSQL from AWS Secrets Manager.")

	default:
		log.Fatalf("Error: APP_ENV environment variable is not defined or has an invalid value: %s. Expected 'development_local_db', 'development_remote_rds', 'production', or 'aws'.", appEnv)
	}

	dbInstance, err := database.ConnectPostgres(dsn)
	if err != nil {
		log.Fatalf("Cannot connect to PostgreSQL: %v", err)
	}
	logger.InfoLog.Println("*** Pinged database successfully! ***") // Agregué este log para confirmar conexión a DB

	productRepo := repository.NewPostgresProductRepository(dbInstance.SQL)

	// Carga sdkConfig si no se ha cargado ya (ej. si no fue el caso "production")
	// La comprobación de Credentials == nil no es la más robusta para saber si sdkConfig es válida,
	// pero es un indicador. Una alternativa es cargarla siempre aquí si no quieres que sea condicional.
	// Sin embargo, si la sdkConfig ya fue asignada en el bloque 'production', ya está lista para Cognito.
	if appEnv != "production" && appEnv != "aws" { // Cargar solo si no se cargó en el caso "production"/"aws"
		sdkConfig, err = awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion(Region))
		if err != nil {
			log.Fatalf("unable to load SDK config for Cognito, %v", err)
		}
	}

	cognitoClient := cognitoservice.NewFromConfig(sdkConfig)
	cognitoAuthService := cognito.NewCognitoAuth(cognitoClient, UserPoolID, ClientID)

	app := &application{
		config:      cfg,
		infoLog:     logger.InfoLog,
		errorLog:    logger.ErrorLog,
		db:          dbInstance,
		productRepo: productRepo,
		cognitoAuth: cognitoAuthService,
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
