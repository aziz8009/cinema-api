package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ValidateRequest function validates the request body based on the given struct and sends error responses if validation fails.
func ValidateRequest(c echo.Context, req interface{}) error {
	// Bind the request body to the struct
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	// Create a new validator instance
	validate := validator.New()

	// Validate the struct
	if err := validate.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
