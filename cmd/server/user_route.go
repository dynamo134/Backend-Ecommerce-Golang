package server

import "github.com/gin-gonic/gin"

func setupUserRoutes(r *gin.RouterGroup, h *HTTPServer) {
	// Define user routes here
	//r.GET("/", h.Handlers.GetAllUsers) // Example route to get all users
	r.POST("/signup", h.Handlers.AuthHandler.SignUp) // Example route to create a new user
	r.POST("/signin", h.Handlers.AuthHandler.SignIn) 
	// r.PUT("/:id", h.Handlers.UpdateUser) // Example route to update a user
	// r.DELETE("/:id", h.Handlers.DeleteUser) // Example route to delete a user
}