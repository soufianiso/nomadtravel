package types

import (
)

type MoviePageNumber int
const MovieSize = 24

type UserRegister struct {
	Name string			`json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type Movie struct {
	ID               int    `json:"id"`
	OriginalTitle    string `json:"original_title"`
	OriginalLanguage string `json:"original_language"`
	Overview         string `json:"overview"`
	ReleaseDate      string `json:"release_date"`
	Adult            bool   `json:"adult"`
	PosterPath       string `json:"poster_path"`
}

type UserLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
