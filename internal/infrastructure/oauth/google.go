package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GoogleOAuthConfig menyimpan konfigurasi yang dibutuhkan untuk Google OAuth
type GoogleOAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

// GoogleOAuthService mengimplementasikan GoogleOAuthService
type GoogleOAuthService struct {
	oauthConfig *oauth2.Config
	httpClient  *http.Client
}

// NewGoogleOAuthService membuat instance baru googleOAuthService
func NewGoogleOAuthService(cfg GoogleOAuthConfig) *GoogleOAuthService {
	oauthConfig := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return &GoogleOAuthService{
		oauthConfig: oauthConfig,
		httpClient:  http.DefaultClient,
	}
}

// GetAuthURL mengembalikan URL yang harus dikunjungi user untuk login dengan Google
func (s *GoogleOAuthService) GetAuthURL(state string) string {
	// state dipakai untuk mencegah CSRF (harus dicek saat callback)
	return s.oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

// ExchangeCode menukar "code" dari callback jadi access token
func (s *GoogleOAuthService) ExchangeCode(code string) (string, error) {
	ctx := context.Background()
	token, err := s.oauthConfig.Exchange(ctx, code)
	if err != nil {
		return "", fmt.Errorf("gagal exchange code: %w", err)
	}

	return token.AccessToken, nil
}

// GoogleUserInfo struktur data basic dari Google API
type GoogleUserInfo struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Sub      string `json:"sub"` // unique Google ID
	Picture  string `json:"picture"`
	Verified bool   `json:"email_verified"`
}

// GetUserInfo mengambil info user dari Google menggunakan access token
func (s *GoogleOAuthService) GetUserInfo(accessToken string) (string, string, string, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		return "", "", "", err
	}

	// Set Authorization header: Bearer <accessToken>
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", "", fmt.Errorf("status code tidak OK: %d", resp.StatusCode)
	}

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return "", "", "", err
	}

	return userInfo.Email, userInfo.Name, userInfo.Sub, nil
}

func (s *GoogleOAuthService) VerifyMobileIDToken(idToken string) (email, name, googleID string, err error) {
	tokenInfoURL := "https://oauth2.googleapis.com/tokeninfo?id_token=" + idToken

	resp, err := http.Get(tokenInfoURL)
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", "", "", errors.New("invalid id_token")
	}

	var payload struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Sub   string `json:"sub"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return "", "", "", err
	}

	return payload.Email, payload.Name, payload.Sub, nil
}
