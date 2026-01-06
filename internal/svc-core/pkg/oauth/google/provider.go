package google

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/konsultin/project-goes-here/dto"
)

const (
	tokenInfoURL = "https://oauth2.googleapis.com/tokeninfo?id_token="
)

// Provider implements Google OAuth authentication
type Provider struct {
	clientID   string
	httpClient *http.Client
}

// NewProvider creates a new Google OAuth provider
func NewProvider(clientID string) *Provider {
	return &Provider{
		clientID: clientID,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// UserInfo represents user info from Google
type UserInfo struct {
	ProviderId    string // Google user ID (sub)
	Email         string
	Name          string
	Picture       string
	EmailVerified bool
}

// TokenInfo represents the response from Google's tokeninfo endpoint
type TokenInfo struct {
	Iss           string `json:"iss"`            // Issuer
	Azp           string `json:"azp"`            // Authorized party
	Aud           string `json:"aud"`            // Audience (client ID)
	Sub           string `json:"sub"`            // Subject (Google user ID)
	Email         string `json:"email"`          // User's email
	EmailVerified string `json:"email_verified"` // "true" or "false"
	Name          string `json:"name"`           // User's full name
	Picture       string `json:"picture"`        // Profile picture URL
	GivenName     string `json:"given_name"`     // First name
	FamilyName    string `json:"family_name"`    // Last name
	Iat           string `json:"iat"`            // Issued at
	Exp           string `json:"exp"`            // Expiration
	Error         string `json:"error"`          // Error message if any
}

// VerifyToken verifies Google ID token and returns user info
func (p *Provider) VerifyToken(ctx context.Context, idToken string) (*UserInfo, error) {
	// Call Google's tokeninfo endpoint
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tokenInfoURL+idToken, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var tokenInfo TokenInfo
	if err := json.Unmarshal(body, &tokenInfo); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Check for error in response
	if tokenInfo.Error != "" {
		return nil, fmt.Errorf("invalid token: %s", tokenInfo.Error)
	}

	// Verify audience matches our client ID
	if tokenInfo.Aud != p.clientID {
		return nil, fmt.Errorf("invalid audience: expected %s, got %s", p.clientID, tokenInfo.Aud)
	}

	// Verify issuer
	if tokenInfo.Iss != "accounts.google.com" && tokenInfo.Iss != "https://accounts.google.com" {
		return nil, fmt.Errorf("invalid issuer: %s", tokenInfo.Iss)
	}

	return &UserInfo{
		ProviderId:    tokenInfo.Sub,
		Email:         tokenInfo.Email,
		Name:          tokenInfo.Name,
		Picture:       tokenInfo.Picture,
		EmailVerified: tokenInfo.EmailVerified == "true",
	}, nil
}

// GetProviderName returns the provider name
func (p *Provider) GetProviderName() string {
	return "google"
}

// GetProviderId returns the provider enum value
func (p *Provider) GetProviderId() dto.AuthProvider_Enum {
	return dto.AuthProvider_GOOGLE
}
