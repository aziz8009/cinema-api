package routes

import (
	"database/sql"
	"net/http"

	"github.com/aziz8009/cinema-app/middleware"
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

		admin := v1.Group("/admin", middleware.JWTMiddleware)
		{

			movie := admin.Group("/movies")
			{

				movie.GET("", h.MovieHandler.GetAll)

				movie.POST("", h.MovieHandler.Create)

				movie.GET("/:id", h.MovieHandler.GetDetailById)

				movie.PUT("/:id", h.MovieHandler.Update)

				movie.DELETE("/:id", func(c echo.Context) error {
					return c.String(http.StatusOK, "Remove movie api")
				})

				movie.GET("/most-viewed", h.MovieHandler.GetMostViewed)

				movie.GET("/genres/most-viewed", h.MovieHandler.GetMostViewedByGenre)
			}
		}

		users := v1.Group("/movies", middleware.JWTMiddleware)
		{
			users.GET("/:id", func(c echo.Context) error {
				return c.String(http.StatusOK, "Get detail movies api")
			})

			users.GET("/search", func(c echo.Context) error {
				return c.String(http.StatusOK, "Search Movie movie api")
			})

			users.POST("/:id/view", func(c echo.Context) error {
				return c.String(http.StatusOK, "Add view movies")
			})

			users.POST("/:id/vote", func(c echo.Context) error {
				return c.String(http.StatusOK, "Add vote movie api")
			})

			users.DELETE("/:id/vote", func(c echo.Context) error {
				return c.String(http.StatusOK, "Add unvote movie api")
			})

			users.GET("/user/votes", func(c echo.Context) error {
				return c.String(http.StatusOK, "Get data vote users movies")
			})
		}

	}

}
