package handlers

import (
	"database/sql"

	"github.com/aziz8009/cinema-app/src/repositories"
	"github.com/aziz8009/cinema-app/src/usecase/auth"
	"github.com/aziz8009/cinema-app/src/usecase/movies"
	"github.com/aziz8009/cinema-app/src/usecase/users"
	"github.com/aziz8009/cinema-app/src/usecase/viewership"
	"github.com/aziz8009/cinema-app/src/usecase/vote"
)

// Handlers struct to hold all handlers
type Handlers struct {
	AuthHandler       *AuthHandler
	UserHandler       *UserHandler
	MovieHandler      *MovieHandler
	ViewerShipHandler *ViewerShipHandler
	VoteHandler       *VoteHandler
}

// InitHandlers initializes all handlers in the application
func InitHandlers(db *sql.DB) *Handlers {
	// Initialize User Handler

	if db == nil {
		panic("db is nil")
	}

	userRepo := repositories.NewUserRepository(db)
	userService := users.NewUserService().SetUserRepo(userRepo).Validate()
	userHandler := NewUserHandler(userService)

	authService := auth.NewAuthService().SetUserRepo(userRepo).Validate()
	authHandler := NewAuthHandler(authService)

	movieRepo := repositories.NewMovieRepository(db)
	movieService := movies.NewMovieService().SetMovieRepo(movieRepo).Validate()
	movieHendler := NewMovieHandler(movieService)

	viewershipRepo := repositories.NewViewerShipRepository(db)
	viewershipService := viewership.NewViewerShipService().SetViewerShipRepo(viewershipRepo).Validate()
	viewershipHandler := NewViewerShipHandler(viewershipService)

	voteRepo := repositories.NewVoteRepository(db)
	voteService := vote.NewVoteService().SetVoteRepo(voteRepo).Validate()
	voteHandler := NewVoteHandler(voteService)

	// Return all handlers as a single struct
	return &Handlers{
		AuthHandler:       authHandler,
		UserHandler:       userHandler,
		MovieHandler:      movieHendler,
		ViewerShipHandler: viewershipHandler,
		VoteHandler:       voteHandler,
	}
}
