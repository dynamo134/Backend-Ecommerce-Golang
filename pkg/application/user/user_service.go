package user

import (
	uRepo "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/infrastructure/persistence/user"
)

// UserService provides methods to interact with user data.
type UserService struct {
	uRepo *uRepo.UserRepository
}

func NewUserService(uRepo *uRepo.UserRepository) *UserService {
	return &UserService{
		uRepo: uRepo,
	}
}