package service

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/soufianiso/nomadtravel/user/configs"
	"github.com/soufianiso/nomadtravel/user/internal/types"
)



func CreateToken(email string) (string, error){
	// This secret key I should retrieve from .env
	secretKey := []byte(configs.Envs.JWTSecret)
	exp := time.Now().Add(1 * time.Hour)
	claims := &types.CustomClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
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
