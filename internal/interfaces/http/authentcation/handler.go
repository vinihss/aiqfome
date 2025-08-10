package http_interfaces_authentication

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenticationHandler struct {
	controller *AuthenticationController
}

func NewAuthenticationHandler(c *AuthenticationController) *AuthenticationHandler {
	return &AuthenticationHandler{
		controller: c,
	}
}

// Authenticate godoc
// @Summary Create authentication token
// @Description Creates a JWT token for the user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param auth body CreateAuthenticationRequest true "Authentication data"
// @Success 200 {object} AuthenticationResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /authenticate [post]
// @Security BearerAuth
func (h *AuthenticationHandler) Authenticate(c *gin.Context) {

	tokenString, err := createToken("test")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to create token: %w", err).Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
