package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenConfig menyimpan konfigurasi yang dibutuhkan untuk bikin dan validasi JWT
type TokenConfig struct {
	AccessSecret    string        // secret key untuk access token
	RefreshSecret   string        // secret key untuk refresh token
	AccessTokenTTL  time.Duration // berapa lama access token berlaku
	RefreshTokenTTL time.Duration // berapa lama refresh token berlaku
}

// TokenService mengimplementasikan interface TokenService (domain/service)
type TokenService struct {
	cfg TokenConfig
}

// CustomClaims adalah payload tambahan dalam JWT kita
type CustomClaims struct {
	UserID string `json:"user_id"` // ID user yang terkait token
	Role   string `json:"role"`    // Role user (admin, staf, viewer)
	jwt.RegisteredClaims
}

// NewTokenService membuat instance baru tokenService
func NewTokenService(cfg TokenConfig) *TokenService {
	return &TokenService{cfg: cfg}
}

// GenerateAccessToken membuat JWT access token dengan expired pendek (mis: 15 menit)
func (s *TokenService) GenerateAccessToken(userID, role string) (string, error) {
	now := time.Now()

	claims := &CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.cfg.AccessTokenTTL)),
		},
	}

	// Kita pakai HMAC dengan secret (HS256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Kenapa pakai secret? Supaya hanya server kita yang bisa bikin token valid
	return token.SignedString([]byte(s.cfg.AccessSecret))
}

// GenerateRefreshToken membuat token dengan masa berlaku lebih lama (mis: 7 hari)
func (s *TokenService) GenerateRefreshToken(userID, role string) (string, error) {
	now := time.Now()

	claims := &CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.cfg.RefreshTokenTTL)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.RefreshSecret))
}

// ValidateToken general, dipakai oleh ValidateAccessToken & ValidateRefreshToken
func (s *TokenService) ValidateToken(tokenStr string, secret string) (string, string, error) {
	// Parse dan validasi token
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Pastikan algorithm-nya HS256 (atau yang kita harapkan)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metode signing tidak valid")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return "", "", errors.New("token tidak valid")
	}

	return claims.UserID, claims.Role, nil
}

func (s *TokenService) ValidateAccessToken(token string) (string, string, error) {
	return s.ValidateToken(token, s.cfg.AccessSecret)
}

func (s *TokenService) ValidateRefreshToken(token string) (string, string, error) {
	return s.ValidateToken(token, s.cfg.RefreshSecret)
}
