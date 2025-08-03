package handlers

import (
	uService "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/application/user"
)

type UserHandler struct {
	UserService *uService.UserService
}

func NewUserHandler(userService *uService.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}