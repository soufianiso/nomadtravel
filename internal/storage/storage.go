package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	moviespb "github.com/soufianiso/nomadtravel/movies/api/v1/proto/movies"
)

type Storage struct {
	db *sql.DB
}



type Store interface {
	GetMovies(*moviespb.ListMoviesRequest) error
}

func NewMoviesStorage(db *sql.DB) *Storage{
	return &Storage{db: db}
}


func(s *Storage) GetMovies(movie *moviespb.ListMoviesRequest) error{
	_, err := s.db.Exec("INSERT INTO movies (original_title, overview, release) VALUES ($1,$2,$3)", movie.GetPage()) 

	if err != nil{
		return err
	}
	//logic

	return nil
}

