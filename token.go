package gojwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type Token interface {
	Generate(payload map[string]any, key string) (string, error)
	Verify(token, key string) bool
	Parse(token, key string) (map[string]any, error)
}

type jwtToken struct {
}

func (j *jwtToken) Generate(payload map[string]any, key string) (string, error) {
	// Create a new token with the HMAC signing method
	token := jwt.New(jwt.SigningMethodHS512)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range payload {
		claims[k] = v
	}

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtToken) Verify(token, key string) bool {
	// Verify the JWT token
	claims, err := j.Parse(token, key)

	return claims != nil && err == nil
}

func (j *jwtToken) Parse(token, key string) (map[string]any, error) {
	// Parse and verify the JWT token
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if !tk.Valid {
		return nil, errors.New("token invalid")
	}

	return tk.Claims.(jwt.MapClaims), nil
}

func New() Token {
	return &jwtToken{}
}
