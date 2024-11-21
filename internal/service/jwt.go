package service

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/soufianiso/nomadtravel/user/configs"
)



func CreateToken(email string) (string, error){
	// This secret key I should retrieve from .env
	secretKey := []byte(configs.Envs.JWTSecret)
	expired := []byte(string(configs.Envs.JWTExpirationInSeconds))

	claims := jwt.MapClaims{
		"Email": email,
		"exp": expired, 
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
