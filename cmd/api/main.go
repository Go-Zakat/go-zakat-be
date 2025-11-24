// @title Auth API
// @version 1.0
// @description API untuk autentikasi (register, login, refresh token, Google OAuth)
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"context"
	"errors"
	"go-zakat/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go-zakat/docs"

	"go-zakat/internal/delivery/http/handler"
	"go-zakat/internal/delivery/http/middleware"
	domainValidator "go-zakat/internal/delivery/http/validator"
	"go-zakat/internal/infrastructure/jwt"
	"go-zakat/internal/infrastructure/oauth"
	"go-zakat/internal/repository/postgres"
	"go-zakat/internal/usecase"

	"go-zakat/pkg/config"
	"go-zakat/pkg/database"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()
	logr := logger.New() // <-- logger init

	// Swagger info
	docs.SwaggerInfo.Title = "Auth API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + cfg.AppPort

	// Database
	dbPool, err := database.NewPostgresPool(cfg.DatabaseURL)
	if err != nil {
		logr.Fatalf("gagal init DB: %v", err)
	}
	defer dbPool.Close()

	val := domainValidator.NewValidator()

	// JWT
	tokenCfg := jwt.TokenConfig{
		AccessSecret:    cfg.JWTAccessSecret,
		RefreshSecret:   cfg.JWTRefreshSecret,
		AccessTokenTTL:  cfg.JWTAccessTTL,
		RefreshTokenTTL: cfg.JWTRefreshTTL,
	}
	tokenSvc := jwt.NewTokenService(tokenCfg)

	// Google
	googleCfg := oauth.GoogleOAuthConfig{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		RedirectURL:  cfg.GoogleRedirectURL,
	}
	googleSvc := oauth.NewGoogleOAuthService(googleCfg)

	userRepo := postgres.NewUserRepository(dbPool, logr)
	authUC := usecase.NewAuthUseCase(userRepo, tokenSvc, googleSvc, val)

	authHandler := handler.NewAuthHandler(authUC)
	authMiddleware := middleware.NewAuthMiddleware(tokenSvc)

	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		for _, o := range cfg.CORSAllowedOrigins {
			if o == origin {
				c.Header("Access-Control-Allow-Origin", origin)
			}
		}
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := router.Group("/auth")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.POST("/refresh", authHandler.Refresh)

		api.GET("/me", authMiddleware.RequireAuth(), authHandler.Me)

		api.GET("/google/login", authHandler.GoogleLogin)
		api.GET("/google/callback", authHandler.GoogleCallback)
		api.POST("/google/mobile/login", authHandler.GoogleMobileLogin)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	go func() {
		logr.Infof("Server berjalan di :%s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logr.Fatalf("Gagal ListenAndServe: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logr.Warn("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logr.Fatalf("Server Shutdown error: %v", err)
	}

	logr.Info("Server exited!")
}
