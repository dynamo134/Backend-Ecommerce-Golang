package handlers

import (
	"net/http"

	aService "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/application/auth"
	uContr "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/contract/authUser"
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

func (ah *AuthHandler) SignUp( ctx *gin.Context) {
	// Implementation for user sign-up
	// This function should handle the user registration logic

	body := uContr.CreateUserRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return 
	}

	res , err := ah.AuthService.SignUp(ctx, &body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	
	ctx.JSON(http.StatusOK, res)

}

func (ah *AuthHandler) SignIn(ctx *gin.Context) {
	// Implementation for user sign-in
	body := uContr.SignInRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// Call the SignIn method of AuthService
	res , err := ah.AuthService.SignIn(ctx, &body)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Set the JWT token in the cookie
	ctx.SetCookie(
		"token",              // name
		res.Token,       // value
		3600,                 // expires in 1 hour
		"/",                  // path
		"localhost",          // domain — use frontend domain in prod
		false,                // secure: true in HTTPS
		true,                 // HttpOnly: yes!
	)

	ctx.JSON(http.StatusOK, res)
}

func (ah *AuthHandler) Logout(ctx *gin.Context) {

	//check if the user is logged in by checking the cookie
	// If the cookie is not found, return an unauthorized error
	_, err := ctx.Cookie("token")
	if err != nil {
		// Cookie not found
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User is not logged in",
		})
		return
	}

	// Clear the cookie
	ctx.SetCookie(
		"token",
		"",
		-1,
		"/",
		"localhost",
		false,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logged out successfully",
	})
}
