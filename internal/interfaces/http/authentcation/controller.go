package http_interfaces_authentication

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthenticationController struct {
}

func NewAuthenticationController() *AuthenticationController {
	return &AuthenticationController{}
}

func (ctrl *AuthenticationController) CreateAuthentication(req CreateAuthenticationRequest) (AuthenticationResponse, error) {

	token, _ := createToken(req.Name)
	return AuthenticationResponse{
		Token: token,
	}, nil
}
func createToken(username string) (string, error) {
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                         // Subject (user identifier)
		"iss": "aiqfome",                        // Issue
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	tokenString, err := claims.SignedString([]byte("secreto123"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
