package service

type GoogleOAuthService interface {
	GetAuthURL(state string) string
	ExchangeCode(code string) (accessToken string, err error)
	GetUserInfo(accessToken string) (email, name, googleID string, err error)
	VerifyMobileIDToken(idToken string) (email, name, googleID string, err error)
}

type TokenService interface {
	GenerateAccessToken(userID, role string) (string, error)
	GenerateRefreshToken(userID, role string) (string, error)
	ValidateAccessToken(token string) (userID, role string, err error)
	ValidateRefreshToken(token string) (userID, role string, err error)
}
