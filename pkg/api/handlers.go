package handlers

type handlers struct {
	AuthHandler *AuthHandler
}

// NewHandlers creates a new handlers instance and initializes it with the given AuthHandler.
// It returns the handlers instance.
func NewHandlers(authHandler *AuthHandler) *handlers {
	return &handlers{
		AuthHandler: authHandler,
	}
}