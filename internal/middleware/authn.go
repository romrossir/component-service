package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/romrossi/component-service/internal/authn"
)

// userContextKey definition remains the same.
type userContextKeyType string

const userContextKey userContextKeyType = "userID"

func AuthnMiddleware(authnService authn.AuthnService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Directly call AuthService with the request.
			// Token extraction is now the responsibility of the AuthService implementation.
			userID, err := authnService.GetUserID(r)
			if err != nil {

				// Respond with a generic unauthorized message, or use err.Error() if it's safe for client exposure.
				// For now, using a generic message.
				http.Error(w, "Unauthorized: Authentication failed", http.StatusUnauthorized)
				return
			}

			// If GetUserID returns an empty userID without an error, it's still an auth failure.
			// This check should ideally be enforced by AuthService implementations (i.e., return an error if userID is empty).
			// However, as a safeguard in the middleware:
			if userID == "" {
				log.Printf("Authentication failed: AuthService returned empty userID without error")
				http.Error(w, "Unauthorized: Authentication failed", http.StatusUnauthorized)
				return
			}

			// Add userID to context for downstream handlers
			ctx := context.WithValue(r.Context(), userContextKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserIDFromContext retrieves the userID from the request context.
// Returns an empty string if userID is not found.
func GetUserIDFromContext(ctx context.Context) string {
	userID, ok := ctx.Value(userContextKey).(string)
	if !ok {
		return ""
	}
	return userID
}
