package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// InitDB initializes the database connection.
// It expects database connection details from environment variables:
// DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE
func Connect() *sql.DB {
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

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("db connect error:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("db ping error:", err)
	}
	return db
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
