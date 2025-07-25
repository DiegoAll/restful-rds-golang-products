package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 5               // numero maximo de conexiones permitidas abiertas
const maxIdleDbConn = 5               // numero maximo de conexiones inactivas (ociosas) Abiertas y disponibles para reutilizacion
const maxDbLifeTime = 5 * time.Minute // tiempo antes de que se considere inactiva una conexion

func ConnectPostgres(dsn string) (*DB, error) {

	fmt.Println("📦 Iniciando conexión a la base de datos PostgreSQL...")
	fmt.Printf("🔑 DSN recibido: %s\n", dsn)

	d, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = d
	fmt.Println("✅ Conexión a PostgreSQL establecida y asignada a dbConn")

	return dbConn, nil
}

func testDB(d *sql.DB) error {
	fmt.Println("🔍 Realizando Ping a la base de datos...")

	err := d.Ping()
	if err != nil {
		fmt.Println("Error!", err)
		fmt.Println("❌ Ping fallido:", err)
		return err
	}
	fmt.Println("*** Pinged database succesfully!***")
	return nil
}
