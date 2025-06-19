package main

import (
	"log"
	"net/http"

	"github.com/romrossi/component-service/internal/authn"
	"github.com/romrossi/component-service/internal/component"
	"github.com/romrossi/component-service/internal/middleware"
	"github.com/romrossi/component-service/pkg/db"
)

func main() {
	// Setup DB connection
	sqlDB := db.Connect()

	// Wire up dependencies for authentication service and middleware
	authnService := authn.NewMockAuthnService()
	authnMiddleware := middleware.AuthnMiddleware(authnService)

	// Wire up dependencies for component service
	repo := &component.PostgresRepository{DB: sqlDB}
	service := &component.DefaultService{Repo: repo}
	cachedService := component.NewCachedService(service)
	handler := &component.Handler{Service: cachedService}

	// Register HTTP router
	http.Handle("/components", authnMiddleware(handler))
	http.Handle("/components/", authnMiddleware(handler))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
