package handlers

import (
	aService "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/application/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService   *aService.AuthService
}

// NewAuthHandler returns a new instance of AuthHandler with the given AuthService.
func NewAuthHandler(authService *aService.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (ah *AuthHandler) SignUp( c *gin.Context) {
	// Implementation for user sign-up
}

func (ah *AuthHandler) SignIn(c *gin.Context) {
	// Implementation for user sign-in
}

func (ah *AuthHandler) Logout(c *gin.Context) {
	// Implementation for user sign-out
}