package main

import (
	"fmt"
	"log"

	"github.com/dynamo134/Backend-Ecomerce-Golang/cmd/server"
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

	// Initialize the server with the configuration
	s , err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to start server, err: %v. Shutting down.", err)
	}
	//setup the routes
	s.SetupRoutes()
	
	// Running the server
	s.Run()
}
