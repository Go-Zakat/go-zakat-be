package middleware

import (
	"net/http"
	"strings"

	"go-zakat/internal/domain/service"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware menyimpan dependencies untuk validasi JWT
type AuthMiddleware struct {
	tokenSvc service.TokenService
}

func NewAuthMiddleware(tokenSvc service.TokenService) *AuthMiddleware {
	return &AuthMiddleware{tokenSvc: tokenSvc}
}

// RequireAuth adalah middleware yang mengecek Authorization: Bearer <token>
func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "Authorization header kosong",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "Format Authorization harus: Bearer <token>",
			})
			return
		}

		tokenStr := parts[1]
		userID, err := m.tokenSvc.ValidateAccessToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "token tidak valid atau expired",
			})
			return
		}

		// Simpan userID ke context supaya handler bisa pakai
		c.Set("user_id", userID)

		c.Next()
	}
}
