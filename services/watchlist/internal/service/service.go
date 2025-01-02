package service

import (
	"context"
	"log/slog"

	"github.com/soufianiso/nomadtravel/services/watchlist/api/v1/proto/watchlist"
	"github.com/soufianiso/nomadtravel/services/watchlist/internal/storage"
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
		s.log.Error("can't retrieve movies","Details",err)
		return nil, err
	}

	s.log.Info("retrieve movies page succefully","page", req.GetPage())
	return &moviespb.ListMoviesResponse{ Movies: movies} , err


}









