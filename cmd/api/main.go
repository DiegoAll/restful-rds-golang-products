package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"restful-rds-golang-products/database"
	"restful-rds-golang-products/internal/pkg/cognito"
	"restful-rds-golang-products/internal/pkg/logger"
	"restful-rds-golang-products/internal/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	cognitoservice "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
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

// const (
// 	UserPoolID = "us-east-1_ZTzSnlG81"
// 	ClientID   = "2oa45rcrl66qophvccaeesdtl9"
// 	Region     = "us-east-1"
// )

func main() {
	logger.Init()

	appEnv := os.Getenv("APP_ENV")
	logger.InfoLog.Printf("Application Environment (APP_ENV): %s", appEnv)

	var cfg config
	cfg.port = 9090

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "9090"
	}

	var dsn string
	var sdkConfig aws.Config

	userPoolID := os.Getenv("UserPoolID") // Leer desde variable de entorno
	fmt.Println("👀 UserPoolID cargado desde .env:", userPoolID)

	if userPoolID == "" {
		log.Fatal("❌ UserPoolID env var is not defined.")
	}
	clientID := os.Getenv("ClientID") // Leer desde variable de entorno
	if clientID == "" {
		log.Fatal("❌ ClientID env var is not defined.")
	}
	region := os.Getenv("Region") // Leer desde variable de entorno
	if region == "" {
		log.Fatal("❌ Region env var is not defined.")
	}

	switch appEnv {
	case "development_local_db":
		dsn = os.Getenv("POSTGRES_LOCAL_DSN")
		if dsn == "" {
			log.Fatal("❌ POSTGRES_LOCAL_DSN env var is not defined.")
		}
		logger.InfoLog.Printf("🔧 DSN from POSTGRES_LOCAL_DSN: %s", dsn)

	case "development_remote_rds":
		dsn = os.Getenv("POSTGRES_RDS_DSN")
		if dsn == "" {
			log.Fatal("❌ POSTGRES_RDS_DSN env var is not defined.")
		}
		logger.InfoLog.Printf("🔧 DSN from POSTGRES_RDS_DSN: %s", dsn)

	case "production", "aws":
		rdsSecretName := os.Getenv("RDS_SECRET_NAME")
		if rdsSecretName == "" {
			log.Fatal("❌ RDS_SECRET_NAME env var is not defined.")
		}
		logger.InfoLog.Printf("🔒 Fetching secret from Secrets Manager: %s", rdsSecretName)

		var err error
		sdkConfig, err = awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion(region))
		if err != nil {
			log.Fatalf("❌ Failed to load AWS SDK config: %v", err)
		}

		smClient := sm.NewFromConfig(sdkConfig)
		getSecretValueOutput, err := smClient.GetSecretValue(context.TODO(), &sm.GetSecretValueInput{
			SecretId: aws.String(rdsSecretName),
		})
		if err != nil {
			log.Fatalf("❌ Failed to retrieve secret '%s': %v", rdsSecretName, err)
		}

		if getSecretValueOutput.SecretString == nil {
			log.Fatal("❌ Secret string is nil for RDS DSN.")
		}

		logger.InfoLog.Printf("🔍 Secret raw string: %s", *getSecretValueOutput.SecretString)

		var secretMap map[string]string
		err = json.Unmarshal([]byte(*getSecretValueOutput.SecretString), &secretMap)
		if err != nil {
			log.Fatalf("❌ Failed to parse JSON from secret: %v", err)
		}

		var ok bool
		dsn, ok = secretMap["FULL_RDS_DSN"]
		if !ok || dsn == "" {
			log.Fatalf("❌ Key 'FULL_RDS_DSN' not found or empty in secret.")
		}

		logger.InfoLog.Printf("🔑 Extracted DSN from secret: %s", dsn)

	default:
		log.Fatalf("❌ Invalid APP_ENV: %s. Use 'development_local_db', 'development_remote_rds', 'production', or 'aws'", appEnv)
	}

	fmt.Printf("✅ FINAL DSN to connect: %s\n", dsn)

	dbInstance, err := database.ConnectPostgres(dsn)
	if err != nil {
		log.Fatalf("❌ Cannot connect to PostgreSQL: %v", err)
	}
	logger.InfoLog.Println("✅ Connected to PostgreSQL (Ping successful).")

	productRepo := repository.NewPostgresProductRepository(dbInstance.SQL)

	if appEnv != "production" && appEnv != "aws" {
		sdkConfig, err = awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion(region))
		if err != nil {
			log.Fatalf("❌ Failed to load AWS SDK config for Cognito: %v", err)
		}
	}

	cognitoClient := cognitoservice.NewFromConfig(sdkConfig)
	// cognitoAuthService := cognito.NewCognitoAuth(cognitoClient, UserPoolID, ClientID)
	cognitoAuthService := cognito.NewCognitoAuth(cognitoClient, userPoolID, clientID)

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
	app.infoLog.Println("🚀 API listening on port", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
