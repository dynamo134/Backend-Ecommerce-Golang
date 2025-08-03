package server


func SetupPublicRoutes(h *HTTPServer) {
	ecom := h.Engine.Group(BasePath)

	// Define API Groups
	userGroup := ecom.Group("/auth")

	// Top-level resource routes
	setupAuthRoutes(userGroup, h)
}

func SetupPrivateRoutes(h *HTTPServer) {
	ecom := h.Engine.Group(BasePath)

	// Define API Groups
	userGroup := ecom.Group("/users")

	// Top-level resource routes
	setupUserRoutes(userGroup, h)

	// Add more private routes here as needed
}