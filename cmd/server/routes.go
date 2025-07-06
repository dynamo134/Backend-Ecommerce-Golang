package server


func SetupRoutes(h *HTTPServer) {
	ecom := h.Engine.Group(BasePath)

	// Define API Groups
	userGroup := ecom.Group("/users")

	// Top-level resource routes
	//setupUserRoutes(userGroup, h)
}