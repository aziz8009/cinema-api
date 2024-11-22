package handlers

import (
	"net/http"

	"github.com/aziz8009/cinema-app/src/usecase/vote"
	"github.com/labstack/echo/v4"
)

type VoteHandler struct {
	service vote.VoteService
}

func NewVoteHandler(service vote.VoteService) *VoteHandler {
	return &VoteHandler{service}
}

func (h *VoteHandler) GetAll(c echo.Context) error {

	res, err := h.service.GetAll(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)

}
