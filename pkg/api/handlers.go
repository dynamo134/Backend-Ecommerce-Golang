package handlers

type Handlers struct {
	AuthHandler *AuthHandler
}

// NewHandlers creates a new handlers instance and initializes it with the given AuthHandler.
// It returns the handlers instance.
func NewHandlers(authHandler *AuthHandler) *Handlers {
	return &Handlers{
		AuthHandler: authHandler,
	}
}