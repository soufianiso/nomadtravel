package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
	// moviespb "github.com/soufianiso/nomadtravel/services/proto/api/v1/proto/"
)

type Storage struct {
	db *sql.DB
}



type Store interface {
	AddItemToWatchlist(int32, int32) error
}

func NewWatchlistStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) AddItemToWatchlist(movieId int32, userId int32) error {
	query := "INSERT INTO watchlists VALUES ($1, $2)" 

	_ , err := s.db.Exec(query , userId, movieId)
	if err != nil {
		return err
	}
	return nil
}
