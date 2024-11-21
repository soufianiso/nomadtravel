package storage

import (
	_ "github.com/lib/pq"
	"database/sql"
	"errors"
)

type Storage struct {
	db *sql.DB
}



type Store interface {
	CreateUser(string, string, string) error
	CheckUserByEmail(string) (bool, error)
	GetEmail(string) (string, error)
}

func NewUserStorage(db *sql.DB) *Storage{
	return &Storage{db: db}
}

func(s *Storage) CreateUser(name string , email string, password string) error{
	query := "INSERT INTO users (name, email, password) VALUES ($1,$2,$3)"
	_, err := s.db.Exec(query, name, email, password)
	if err != nil{
		return err
	}

	return nil
}


func (s *Storage) CheckUserByEmail(email string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	err := s.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}



func(s *Storage) GetEmail(email string) (string, error){
	var hashedPassword string

	query := "SELECT password FROM users WHERE email = $1"
	err := s.db.QueryRow(query, email).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// Email not found
			return  "", errors.New("Email not found") 
		}
		// Handle other errors
		return  "", err
	}

	return hashedPassword, nil

}
