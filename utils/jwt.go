package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// getJWTKey retrieves the JWT secret key from the environment variables.
// It loads the environment variables from a .env file.
// If the .env file cannot be loaded or the JWT_KEY environment variable is not set an
// error is created, otherwise the key is returned.
func getJWTKey() (string, error) {
	err := godotenv.Load()

	if err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	jwtKey := os.Getenv("JWT_KEY")

	if jwtKey == "" {
		return "", fmt.Errorf("JWT_KEY environment variable not set")
	}

	return jwtKey, nil
}

// GenerateToken generates a JWT token for a given email and user ID.
// The token is signed using the HMAC SHA256 algorithm and expires after 2 hours.
// It returns the signed token string or any error encountered during the process.
func GenerateToken(email string, userId int64) (string, error) {

	key, err := getJWTKey()

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(key))
}

// VerifyToken verifies the given JWT token and extracts the user ID from it.
// It returns the user ID and any error encountered during the verification process.
// If the token is invalid or the signing method is unexpected, the error is populated.
func VerifyToken(token string) (int64, error) {

	key, err := getJWTKey()

	if err != nil {
		return -1, err
	}

	// Parse the token and validate the signature
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, typeMatch := token.Method.(*jwt.SigningMethodHMAC)

		if !typeMatch {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(key), nil
	})

	if err != nil {
		return -1, errors.New("could not parse token")
	}

	// Check if the token is valid
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return -1, errors.New("invalid token")
	}

	// Extract the claims and ensure they are of the correct type
	claims, typeCheck := parsedToken.Claims.(jwt.MapClaims)

	if !typeCheck {
		return -1, errors.New("invalid token claims")
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil
}
