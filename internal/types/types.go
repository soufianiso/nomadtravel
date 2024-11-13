package types

import (
	"time"
)

type MoviePageNumber int
const MovieSize = 20

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Movie struct {
	ID        int       `json:"id"`
	OriginalTitle string    `json:"original_title"`
	OriginalLanguage  string    `json:"original_Language"`
	Overview     string    `json:"overview"`
	Release_date  string    `json:"release_date"`
	Adult  string    `json:"adult"`
}
