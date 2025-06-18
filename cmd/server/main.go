package main

import (
	"log"
	"net/http"

	"github.com/romrossi/component-api/internal/component"
	"github.com/romrossi/component-api/internal/db"
)

func main() {
	// connStr := "postgres://user:password@localhost/dbname?sslmode=disable"
	connStr := "postgres://postgres:postgres@localhost/component?sslmode=disable"
	database, err := db.Connect(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	handler := component.NewHandler(database)

	http.HandleFunc("/components", handler.ComponentsHandler)
	http.HandleFunc("/components/", handler.ComponentHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
