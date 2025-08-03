package handlers

import (
	aService "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/application/auth"
	uService "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/application/user"
	aRepo "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/infrastructure/persistence/auth"
	uRepo "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/infrastructure/persistence/user"
	"go.mongodb.org/mongo-driver/mongo"
)

// InitHandlers initializes and wires up all the HTTP handlers with their services.
func InitHandlers(mongoClient *mongo.Client) *Handlers {
	// Initialize the AuthRepository with the provided MongoDB client
	aRepo := aRepo.NewAuthRepository(mongoClient)
	// Initialize the UserRepository with the provided MongoDB client
	uRepo := uRepo.NewUserRepository(mongoClient)

	// Initialize the service layer
	authService := aService.NewAuthService(aRepo)
	// Initialize the UserService with the UserRepository
	userService := uService.NewUserService(uRepo)

	// Initialize the corresponding handler
	authHandler := NewAuthHandler(authService)
	// Initialize the UserHandler with the UserService
	userHandler := NewUserHandler(userService)

	// Return the main Handlers struct with sub-handlers injected
	return NewHandlers(authHandler, userHandler)
}
