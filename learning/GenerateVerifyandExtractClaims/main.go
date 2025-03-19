package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your_secret_key")

// Generate a JWT token
func generateToken() (string, error) {
	claims := jwt.MapClaims{
		"username": "exampleUser",
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // Token expires in 2 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Verify and parse the JWT token
func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func main() {
	// Step 1: Generate a token
	tokenString, err := generateToken()
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	fmt.Println("Generated Token:", tokenString)

	// Step 2: Verify the token
	verifiedToken, err := verifyToken(tokenString)
	if err != nil {
		fmt.Println("Token verification failed:", err)
		return
	}
	fmt.Println("Token successfully verified!")

	// Step 3: Extract claims
	if claims, ok := verifiedToken.Claims.(jwt.MapClaims); ok && verifiedToken.Valid {
		fmt.Println("Username:", claims["username"])
		fmt.Println("Expiration Time:", time.Unix(int64(claims["exp"].(float64)), 0))
	} else {
		fmt.Println("Invalid token claims")
	}
}
