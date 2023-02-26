package main

import (
	"time"

	"github.com/golang-jwt/jwt"
)

//==================================================================================================

var jwtSecret = []byte("some_secret_key")

//----------------------------------------------------------------------------------------------

// function to create a new JWT token
func generateToken(id string) (string, error) {

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//----------------------------------------------------------------------------------------------

// function to verify the JWT token
func verifyToken(tokenString string) (string, error) {

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, new(jwt.ValidationError)
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	// Validate the token and return the user ID
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(string)
		return id, nil
	}

	return "", err
}
