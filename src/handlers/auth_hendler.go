package handlers

import (
	"net/http"

	"github.com/aziz8009/cinema-app/src/usecase/auth"
	"github.com/aziz8009/cinema-app/utils"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service auth.AuthService
}

func NewAuthHandler(service auth.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) Login(c echo.Context) error {

	var req auth.LoginRequest

	if err := utils.ValidateRequest(c, &req); err != nil {
		return err
	}

	res, err := h.service.Login(c, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Register(c echo.Context) error {

	var req auth.RegisterRequest

	if err := utils.ValidateRequest(c, &req); err != nil {
		return err
	}

	res, err := h.service.Register(c, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}
