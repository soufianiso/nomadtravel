package service

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/soufianiso/nomadtravel/user/configs"
	"time"
)



func CreateToken(email string) (string, error){
	// This secret key I should retrieve from .env
	secretKey := []byte(configs.Envs.JWTSecret)
	expired := configs.Envs.JWTExpirationInSeconds

	claims := jwt.MapClaims{
		"Email": email,
		"exp": time.Now().Add(time.Duration(expired) * time.Second).Unix(), // Unix timestamp for expiration
	}
	

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil{
		return "", err
	}

	return  tokenString , nil

}
