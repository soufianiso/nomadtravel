package storage

import (
	"context"
	"database/sql"
	"errors"

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
   // Validate pagination parameters
    page := req.GetPage()
    if page <= 0 {
        page = 1
    }

    pageSize := req.GetPageSize()
    if pageSize <= 0 {
        pageSize = 24 // Default page size
    }

    offset := (page - 1) * pageSize

	rows, err := s.db.QueryContext(ctx, 
		`SELECT original_title, original_language, overview, release_date,  adult, poster_path 
		FROM movies 
		ORDER BY original_title 
		LIMIT $1 OFFSET $2`, pageSize, offset) 

	if err != nil{
		return nil, err
	}

    defer rows.Close()

	var movies []*moviespb.Movie

	for rows.Next(){
		var movie moviespb.Movie
		err := rows.Scan(&movie.OriginalTitle,
			&movie.OriginalLanguage,
			&movie.Overview,  
			&movie.ReleaseDate,
			&movie.Adult,
			&movie.PosterPath) 
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(movies) == 0{

		return nil , errors.New("page not been found")

	}
	// Check for errors during iteration


	return  movies, nil
}

