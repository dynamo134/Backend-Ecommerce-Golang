package server

import (
	"log"

	"github.com/dynamo134/Backend-Ecomerce-Golang/config"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Engine *gin.Engine
	Config *config.AppConfig
	// Handlers *h.Handlers
}

func NewServer(cfg *config.AppConfig) (*HTTPServer, error) {
	// set the Gin mode based on the environment (In future we can set this in config -> (local, dev)- debugMode, production-> releaseMode)
	gin.SetMode(gin.ReleaseMode)
	// Initialize the Gin engine
	engine := gin.New()

	engine.Use(gin.Logger(), gin.Recovery())

	// Create a new HTTPServer instance
	return &HTTPServer{
		Engine: engine,
		Config: cfg,
		// Handlers: h.NewHandlers(cfg), // Assuming you have a Handlers struct
	},nil

}

func (s *HTTPServer) Run() {
	// Start the server on the configured port
	log.Printf("Starting server on port %s", s.Config.Port)
	if err := s.Engine.Run(":" + s.Config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}