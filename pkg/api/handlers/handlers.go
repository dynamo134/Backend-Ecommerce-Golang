package handlers

type Handlers struct {
	AuthHandler *AuthHandler
	UserHandler *UserHandler 
}

// NewHandlers creates a new handlers instance and initializes it with the given AuthHandler.
// It returns the handlers instance.
func NewHandlers(
	authHandler *AuthHandler,
	userHandler *UserHandler,
	) *Handlers {
	return &Handlers{
		AuthHandler: authHandler,
		UserHandler: userHandler,
	}
}