package handlers

import (
	"net/http"

	"github.com/aziz8009/cinema-app/src/usecase/viewership"
	"github.com/labstack/echo/v4"
)

type ViewerShipHandler struct {
	service viewership.ViewerShipService
}

func NewViewerShipHandler(service viewership.ViewerShipService) *ViewerShipHandler {
	return &ViewerShipHandler{service}
}

func (h *ViewerShipHandler) GetAll(c echo.Context) error {

	res, err := h.service.GetAll(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)

}
