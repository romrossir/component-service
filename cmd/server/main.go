package main

import (
	"log"
	"net/http"

	"github.com/romrossi/component-service/internal/component"
	"github.com/romrossi/component-service/pkg/db"
)

func main() {
	// Setup DB connection
	sqlDB := db.Connect()

	// Wire up dependencies
	repo := &component.PostgresRepository{DB: sqlDB}
	service := &component.DefaultService{Repo: repo}
	handler := &component.Handler{Service: service}

	// Register HTTP router
	http.Handle("/components", handler)
	http.Handle("/components/", handler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
