package dto

// --- Requests ---

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

// --- Responses ---

type AuthTokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthURLResponse struct {
	AuthURL string `json:"auth_url"`
}

type UserResponse struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	Name      string  `json:"name"`
	Role      string  `json:"role"`
	GoogleID  *string `json:"google_id,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
}

type AuthResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

// UpdateRoleRequest for updating user role
type UpdateRoleRequest struct {
	Role string `json:"role" validate:"required,oneof=admin staf viewer"`
}
