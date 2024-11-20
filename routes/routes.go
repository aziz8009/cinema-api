package routes

import (
	"database/sql"
	"net/http"

	"github.com/aziz8009/cinema-app/src/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {

	h := handlers.InitHandlers(db)

	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "API CINEMA APP")
	})

	v1 := e.Group("/v1")
	{
		v1.POST("/login", h.AuthHandler.Login)
		v1.POST("/register", h.AuthHandler.Register)
	}

}
