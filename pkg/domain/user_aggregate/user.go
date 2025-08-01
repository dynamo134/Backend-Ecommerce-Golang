package user_aggregate

import (
	"context"
	"errors"
	"strings"
	"time"
)

type User struct {
	ID             string
	Email          string
	Name           string
	HashedPassword string
	CreatedAt      time.Time
	EmailVerified  bool
}

// NewUser creates a new User instance with basic validation.
func NewUser(ctx context.Context, id, email, username, hashedPwd,gender, phone, address string  ) (*User, error) {
	email = strings.TrimSpace(email)
	username = strings.TrimSpace(username)

	if email == "" || username == "" || hashedPwd == "" {
		return nil, errors.New("email, name, and hashed password are required")
	}

	user := &User{
		ID:             id,
		Email:          email,
		Name:           username,
		HashedPassword: hashedPwd,
		EmailVerified:  false,
		CreatedAt:      time.Now(),
	}
	return user, nil
}
