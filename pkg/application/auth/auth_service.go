package auth

import (
	aRepo "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/infrastructure/persistence/auth"
)

type AuthService struct {
	authRepo *aRepo.AuthRepository
}