package main

import "github.com/dgrijalva/jwt-go"
import "time"
import "fmt"

import "os"

var signingKey = os.Getenv("HMAC_KEY")

func getValidationKeyForToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(signingKey), nil
}

func getSignedTokenForUser(username string) (string, error) {
	claims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
		Issuer:    "whatsforlunch",
		Subject:   fmt.Sprint(username),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func validateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, getValidationKeyForToken)

	if claims, ok := token.Claims.(jwt.StandardClaims); ok && token.Valid {
		return claims.Subject, err
	}

	return "", err
}
