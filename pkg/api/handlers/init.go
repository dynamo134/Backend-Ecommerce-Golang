package handlers

import (
	aService "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/application/auth"
	aRepo "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/infrastructure/persistence/auth"
	"go.mongodb.org/mongo-driver/mongo"
)

// InitHandlers initializes and wires up all the HTTP handlers with their services.
func InitHandlers(mongoClient *mongo.Client) *Handlers {
	// Initialize the AuthRepository with the provided MongoDB client
	aRepo := aRepo.NewAuthRepository(mongoClient)

	// Initialize the service layer
	authService := aService.NewAuthService(aRepo)

	// Initialize the corresponding handler
	authHandler := NewAuthHandler(authService)

	// Return the main Handlers struct with sub-handlers injected
	return NewHandlers(authHandler)
}
