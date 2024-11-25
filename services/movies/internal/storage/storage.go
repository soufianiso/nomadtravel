package storage

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	moviespb "github.com/soufianiso/nomadtravel/movies/api/v1/proto/movies"
)

type Storage struct {
	db *sql.DB
}



type Store interface {
	GetMovies(context.Context ,*moviespb.ListMoviesRequest) ([]*moviespb.Movie, error)
}

func NewMoviesStorage(db *sql.DB) *Storage{
	return &Storage{db: db}
}


func(s *Storage) GetMovies(ctx context.Context , req *moviespb.ListMoviesRequest) ([]*moviespb.Movie, error){
	page := req.GetPage()
	pageSize := req.GetPageSize()
	offset := (page - 1) * pageSize

	rows, err := s.db.QueryContext(ctx, 
		`SELECT original_title, original_language, overview, release_date,  adult, poster_path 
		FROM movies 
		ORDER BY original_title 
		LIMIT $1 OFFSET $2`, pageSize, offset) 
	if err != nil{
		return nil, err
	}

	var movies []*moviespb.Movie

	for rows.Next(){
		var movie moviespb.Movie
		if err := rows.Scan(&movie.OriginalTitle, &movie.OriginalLanguage, &movie.Overview,  &movie.ReleaseDate, &movie.Adult, &movie.PosterPath); err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}
	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}


	return  movies, nil
}

