package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

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


// func Createjwt(email string, secret string) (string, error){
// 	signingKey := []byte("secret")	
// 	claims := jwt.MapClaims{
// 		"Email": email,
// 		"exp": time.Now().Add(time.Hour * 72).Unix(),
// 	}

// 	// Create the token with claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Sign the token with your secret key
// 	tokenString, err := token.SignedString(signingKey)

// 	if err != nil{
// 		return "" ,err
// 	}

// 	return tokenString, nil
// }


