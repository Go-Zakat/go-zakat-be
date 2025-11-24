package dto

// RegisterRequest merepresentasikan JSON body untuk /auth/register
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// GoogleMobileLoginRequest Untuk mobile Google login (pakai id_token)
type GoogleMobileLoginRequest struct {
	IDToken string `json:"id_token" binding:"required"`
}
