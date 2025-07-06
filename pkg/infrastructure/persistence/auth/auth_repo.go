package auth

import "go.mongodb.org/mongo-driver/mongo"

type AuthRepository struct {
	client *mongo.Client
}