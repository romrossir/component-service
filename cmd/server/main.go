package main

import (
	"log"
	"net/http"

	"github.com/romrossi/component-service/internal/component"
	"github.com/romrossi/component-service/internal/db"
)

func main() {
	// Setup DB connection
	db.InitDB()

	// Setup HTTP routing
	handler := component.NewHandler(db.GetDB())
	http.HandleFunc("/components", handler.ComponentsHandler)
	http.HandleFunc("/components/", handler.ComponentHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
