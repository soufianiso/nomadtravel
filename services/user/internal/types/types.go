package types

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

 


type CustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}




 







