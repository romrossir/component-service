package authn

import (
	"fmt"
	"net/http"
)

// MockAuthnService mocks AuthService
type MockAuthService struct {
}

func NewMockAuthnService() *MockAuthService {
	return &MockAuthService{}
}

func (s *MockAuthService) GetUserID(r *http.Request) (string, error) {
	if r == nil {
		return "", fmt.Errorf("http request is nil") // Basic sanity check
	}

	cookie, err := r.Cookie("MOS_AUTH_TOKEN")
	if err != nil {
		if err == http.ErrNoCookie {
			return "", fmt.Errorf("authentication cookie 'MOS_AUTH_TOKEN' not found")
		}
		// For any other error reading cookies
		return "", fmt.Errorf("error reading authentication cookie: %w", err)
	}

	token := cookie.Value
	if token == "" {
		return "", fmt.Errorf("token in 'MOS_AUTH_TOKEN' cookie is empty")
	}

	return token, nil
}
