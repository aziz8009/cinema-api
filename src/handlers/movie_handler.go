package handlers

import (
	"net/http"

	"github.com/aziz8009/cinema-app/src/usecase/movies"
	"github.com/labstack/echo/v4"
)

type MovieHandler struct {
	service movies.MovieService
}

func NewMovieHandler(service movies.MovieService) *MovieHandler {
	return &MovieHandler{service}
}

func (h *MovieHandler) GetAll(c echo.Context) error {

	res, err := h.service.GetAll(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)

}

func (h *MovieHandler) GetDetailById(c echo.Context) error {

	res, err := h.service.GetDetailById(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)

}

func (h *MovieHandler) Create(c echo.Context) error {

	res, err := h.service.Create(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)

}

func (h *MovieHandler) Update(c echo.Context) error {

	res, err := h.service.Update(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)

}

func (h *MovieHandler) GetMostViewed(c echo.Context) error {

	res, err := h.service.GetMostViewed(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)

}

func (h *MovieHandler) GetMostViewedByGenre(c echo.Context) error {

	res, err := h.service.GetMostViewedByGenre(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)

}
