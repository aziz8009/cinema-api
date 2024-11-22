package handlers

import (
	"net/http"

	"github.com/aziz8009/cinema-app/src/usecase/users"
	"github.com/aziz8009/cinema-app/utils"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service users.UserService
}

func NewUserHandler(service users.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.service.GetAllUsers(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req users.UserReq

	if err := utils.ValidateRequest(c, &req); err != nil {
		return err
	}

	res, err := h.service.CreateUser(c, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}
