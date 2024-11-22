package utils

import (
	"encoding/json"
	"fmt"

	"github.com/aziz8009/cinema-app/utils/constants"
	"github.com/labstack/echo/v4"
)

func AuthData(c echo.Context) (*constants.AuthData, error) {
	// Retrieve the data from the context
	rawAuth := c.Get("user_auth")
	if rawAuth == nil {
		return nil, fmt.Errorf("user_auth not found in context")
	}

	// Check if rawAuth is a map[string]interface{}
	authMap, ok := rawAuth.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("user_auth is not of type map[string]interface{}")
	}

	// Convert map[string]interface{} to constants.AuthData
	var authData constants.AuthData
	jsonBytes, err := json.Marshal(authMap)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user_auth: %v", err)
	}

	err = json.Unmarshal(jsonBytes, &authData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user_auth into AuthData: %v", err)
	}

	return &authData, nil
}
