package auth

import (
	"context"
	"time"

	uAgg "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/domain/user_aggregate"
)

type userModel struct {
	ID             string `bson:"_id,omitempty"`
	Email          string `bson:"email"`
	Name           string `bson:"name"`
	HashedPassword string `bson:"hashed_password"`
	CreatedAt      int64  `bson:"created_at"`
}

func newUserModel(ctx context.Context, ua *uAgg.User) *userModel {
	return &userModel{
		ID:             ua.ID,
		Email:          ua.Email,
		Name:           ua.Name,
		HashedPassword: ua.HashedPassword,
		CreatedAt:      ua.CreatedAt.Unix(),
	}
}

func (um *userModel) toDomain() *uAgg.User {
	return &uAgg.User{
		ID:             um.ID,
		Email:          um.Email,
		Name:           um.Name,
		HashedPassword: um.HashedPassword,
		CreatedAt:      time.Unix(um.CreatedAt, 0),
	}
}