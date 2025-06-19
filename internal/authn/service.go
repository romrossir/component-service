package authn

import "net/http"

// AuthnService defines the interface for authentication operations.
type AuthnService interface {
	// GetUserID retrieves a user ID by inspecting the given HTTP request.
	// It returns the userID if authentication is successful, or an error otherwise.
	GetUserID(r *http.Request) (string, error)
}
