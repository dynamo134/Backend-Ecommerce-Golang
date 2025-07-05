package main

import (
	"fmt"
	"log"

	"github.com/dynamo134/Backend-Ecomerce-Golang/config"
)

func main() {
	println("Hello, World!")

	// test DB connection
	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbClient := config.NewMongoClient(cfg)
	fmt.Println("✅ MongoDB connection established successfully.",dbClient)

	
}
