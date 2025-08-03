package server

import "github.com/gin-gonic/gin"

func setupAuthRoutes(r *gin.RouterGroup, h *HTTPServer) {
	// Define auth routes here
	r.POST("/signup", h.Handlers.AuthHandler.SignUp) 
	r.POST("/signin", h.Handlers.AuthHandler.SignIn)
	r.POST("/logout", h.Handlers.AuthHandler.Logout)
}
