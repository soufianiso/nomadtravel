package service

import (
	"context"
	"log/slog"
	moviespb "github.com/soufianiso/nomadtravel/movies/api/v1/proto/movies"
	"github.com/soufianiso/nomadtravel/movies/internal/storage"
	// "user/internal/types"
)

type MoviesService struct {
	moviespb.UnimplementedMoviesServer
	log *slog.Logger 
	// interface of storage
	store storage.Store
}




func NewMoviesService(log *slog.Logger , store storage.Store) *MoviesService{
	return &MoviesService{
		log : log,
		store: store,

	}
}

func (s *MoviesService) ListMovies(ctx context.Context, req *moviespb.ListMoviesRequest) (*moviespb.ListMoviesResponse, error) {
	movies, err := s.store.GetMovies(ctx, req)
	if err !=  nil{
		s.log.Error("something went wrong in GetMovies query","Err",err)
		
	}
	return &moviespb.ListMoviesResponse{ Movies: movies} , err


}









