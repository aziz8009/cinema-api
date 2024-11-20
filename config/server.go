package config

import (
	"log"
	"os"

	md "github.com/aziz8009/cinema-app/middleware"
	"github.com/aziz8009/cinema-app/routes"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func StartServer(e *echo.Echo) {

	v := validator.New()
	e.Validator = &md.DataValidator{ValidatorData: v}

	e.HTTPErrorHandler = md.ErrorHandler

	// Connect to the database
	db, err := ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Load routes
	routes.RegisterRoutes(e, db)

	port := ":" + os.Getenv("APP_PORT") // Change as needed

	// Start server
	log.Printf("Server running on http://localhost%s", port)

	e.Logger.Fatal(e.Start(port))

	defer db.Close()
}
