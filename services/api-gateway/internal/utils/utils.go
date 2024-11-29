package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/soufianiso/nomadtravel/api-gateway/configs"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/types"
)



const (
	ErrNameEmpty      = "Name is empty"
	ErrPasswordEmpty  = "Password is empty"
	ErrNotObjectIDHex = "String is not a valid hex representation of an ObjectId"
)

func Encode(w http.ResponseWriter, r *http.Request, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}
	return nil
}

func Decode(r *http.Request, v any) (error) {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return  fmt.Errorf("decode json: %w", err)
	}
	return  nil
}

func ParseClaims(tokenString string) (error, *types.CustomClaims ) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return configs.Envs.JWTSecret, nil
	})
	if err != nil {
		return err , nil
	}

	claims := &types.CustomClaims{}

	if !token.Valid {
		return fmt.Errorf("invalid token"), nil
	}
	return nil, claims
}

