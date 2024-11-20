package handlers

import (
	"database/sql"

	"github.com/aziz8009/cinema-app/src/repositories"
	"github.com/aziz8009/cinema-app/src/usecase/auth"
)

// Handlers struct to hold all handlers
type Handlers struct {
	AuthHandler *AuthHandler
}

// InitHandlers initializes all handlers in the application
func InitHandlers(db *sql.DB) *Handlers {
	// Initialize User Handler

	if db == nil {
		panic("db is nil")
	}

	userRepo := repositories.NewUserRepository(db)

	authService := auth.NewAuthService().SetUserRepo(userRepo).Validate()
	authHandler := NewAuthHandler(authService)

	// Return all handlers as a single struct
	return &Handlers{
		AuthHandler: authHandler,
	}
}
