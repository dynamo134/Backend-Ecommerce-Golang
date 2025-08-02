package utils

import (
	"fmt"
	"time"

	"github.com/dynamo134/Backend-Ecommerce-Golang/config"
	"github.com/golang-jwt/jwt"
)

// Create a JWT token using the provided secret key.

func CreateJWT(username, userId string) (string, error) {	
	config, err := config.SetConfig() // Assuming you have a function to get the config
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
	})
	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}



// Token VErification

func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	config, err := config.SetConfig()
	if err != nil {
		return nil, err
	}

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Optional: check signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return []byte(config.JWTSecret), nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("invalid token")
}

