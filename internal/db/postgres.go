package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection.
// It expects database connection details from environment variables:
// DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE
func InitDB() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "component")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" {
		log.Fatal("Database environment variables (DB_HOST, DB_PORT, DB_USER, DB_NAME) are required.")
	}

	if dbSSLMode == "" {
		dbSSLMode = "disable" // Default SSL mode
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v. Please ensure PostgreSQL is running and accessible, and the connection details are correct.", err)
	}

	log.Println("Successfully connected to the PostgreSQL database!")
}

func GetDB() *sql.DB {
	if DB == nil {
		// This case should ideally not happen if InitDB is called at application start.
		// Consider how to handle this based on your application's lifecycle.
		// For now, we'll log a fatal error.
		log.Fatal("Database connection is not initialized. Call InitDB first.")
	}
	return DB
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
