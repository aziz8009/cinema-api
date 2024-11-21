package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/aziz8009/cinema-app/utils"
	"github.com/aziz8009/cinema-app/utils/constants"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := c.Request().Context()

		token := c.Request().Header.Get("Authorization")
		if token == "" {
			log.Println(ctx, "Empty Authorization Token", token)
			return c.JSON(http.StatusUnauthorized, constants.GetCustomResponse("4010", "empty token", nil, []string{errors.New("empty authorization header").Error()}))
		}

		tokenParts := strings.Split(token, " ")

		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			log.Println(ctx, "Invalid Token Format", token)
			return c.JSON(http.StatusUnauthorized, constants.GetCustomResponse("4010", "invalid token format", nil, []string{errors.New("invalid token format").Error()}))
		}

		tokenStr := tokenParts[1]

		userAuth, err := utils.ParseToken(tokenStr)

		if err != nil {
			log.Println(ctx, "Invalid Token Format", err)
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		c.Set("user_auth", userAuth)

		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	// Need this, because somehow if default error handler use with echo body dump
	// It will be print response error twice
	if c.Get("error-handled") != nil {
		return
	}

	c.Set("error-handled", true)

	resp := constants.DefaultResponse{
		Status:  constants.STATUS_FAILED,
		Message: err.Error(),
		Data:    struct{}{},
		Errors:  make([]string, 0),
	}
	if c.Get("invalid-format") == true {
		resp.Status = constants.STATUS_JSON_VALIDATION_FAILED
		resp.Message = constants.MESSAGE_INVALID_REQUEST_FORMAT
		resp.Errors = []string{err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	} else if c.Get("unauthorized") != nil {
		resp.Status = constants.STATUS_INVALID_AUTHORIZATION
		resp.Message = constants.MESSAGE_UNAUTHORIZED
	} else if c.Get("forbidden") != nil {
		resp.Status = constants.STATUS_FORBIDDEN
		resp.Message = constants.MESSAGE_FORBIDDEN
	} else if strings.Contains(err.Error(), "Error 1062") || strings.Contains(err.Error(), "SQLSTATE 23505") {
		resp.Status = constants.STATUS_MULTIPLE_IDENTIFIER
		resp.Message = constants.MESSAGE_CONFLICT
	} else if strings.Contains(err.Error(), "invalid ownership") || strings.Contains(err.Error(), "bid or pid not provided") || strings.Contains(err.Error(), "pid header not provided") {
		resp.Status = constants.STATUS_INVALID_IDENTIFIER
	}

	c.JSON(http.StatusOK, resp)
}
