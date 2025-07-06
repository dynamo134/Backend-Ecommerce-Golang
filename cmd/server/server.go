package server

import (
	"log"

	"github.com/dynamo134/Backend-Ecomerce-Golang/config"
	"github.com/gin-gonic/gin"
)

var (

	BasePath 		= "/api/v1"
)

type HTTPServer struct {
	Engine *gin.Engine
	Config *config.AppConfig
	// Handlers *h.Handlers
}

// NewServer initializes a new HTTPServer instance with a Gin engine.
// It sets the Gin mode to release mode and applies default middleware 
// for logging and recovery. The server is configured with the provided 
// AppConfig, which includes settings like the port number. This function 
// returns the HTTPServer instance and an error if any initialization issues occur.

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

// Run starts the server on the configured port and logs an error if the server fails to start.
func (s *HTTPServer) Run() {
	// Start the server on the configured port
	log.Printf("Starting server on port %s", s.Config.Port)
	if err := s.Engine.Run(":" + s.Config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (h *HTTPServer) SetupRoutes() {
	SetupRoutes(h)
}