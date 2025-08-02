package auth

import (
	"context"
	"fmt"

	uAgg "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/domain/user_aggregate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	client *mongo.Client
}
// NewAuthRepository creates a new instance of AuthRepository with the provided MongoDB client.
func NewAuthRepository(client *mongo.Client) *AuthRepository {
	return &AuthRepository{
		client: client,
	}
}

// Returns the user collection handle
func (r *AuthRepository) userCollection(ctx context.Context) *mongo.Collection {
	return r.client.Database("Ecommerce").Collection("users")
}

// Create inserts a new user into MongoDB
func (r *AuthRepository) CreateUser(ctx context.Context, ua *uAgg.User) error {
	u := newUserModel(ctx, ua)

	id, err := r.userCollection(ctx).InsertOne(ctx, u)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	fmt.Sprintf("✅ Created user id: %v", id.InsertedID)
	return nil
}

func (r *AuthRepository) GetUserByUsername(ctx context.Context, username string) (*uAgg.User, error) {
	var userModel userModel
	filter := bson.M{"name": username}
	err := r.userCollection(ctx).FindOne(ctx, filter).Decode(&userModel)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}

	return userModel.toDomain(), nil
}