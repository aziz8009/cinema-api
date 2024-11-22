package utils

import (
	"encoding/json"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Data any `json:"data"`
	jwt.StandardClaims
}

func GenerateToken(data any) (string, error) {

	var secretKey = os.Getenv("JWT_SECRET_KEY")
	var secretIv = os.Getenv("JWT_SECRET_IV")
	var exp = time.Now().Add(7 * 24 * time.Hour).Unix()

	bt, err := json.Marshal(data)

	if err != nil {
		return "", nil
	}

	claims := jwt.MapClaims{
		"data": EncryptAES256CBC(string(bt), secretKey, secretIv),
		"exp":  exp, // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseToken parses a JWT and returns the claims
func ParseToken(tokenString string) (claims *Claims, err error) {
	var secretKey = os.Getenv("JWT_SECRET_KEY")

	var secretIv = os.Getenv("JWT_SECRET_IV")

	secret := []byte(secretKey)
	// var secretIv = os.Getenv("JWT_SECRET_IV")

	claims = &Claims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	encryptedString, err := DecryptAES256CBC(claims.Data.(string), secretKey, secretIv) // * should be string
	if err != nil {
		return nil, err
	}

	var data any
	err = json.Unmarshal([]byte(encryptedString), &data)
	if err != nil {
		return nil, err
	}

	claims.Data = data

	return
}
