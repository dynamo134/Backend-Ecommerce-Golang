package server

import "github.com/gin-gonic/gin"

func setupUserRoutes(r *gin.RouterGroup, h *HTTPServer) {
	// Define user routes here
	//r.GET("/", h.Handlers.GetAllUsers) // Example route to get all users
	r.POST("/", h.Handlers.AuthHandler.SignUp) // Example route to create a new user
	// r.GET("/:id", h.Handlers.GetUserByID) // Example route to get a user by ID
	// r.PUT("/:id", h.Handlers.UpdateUser) // Example route to update a user
	// r.DELETE("/:id", h.Handlers.DeleteUser) // Example route to delete a user
}