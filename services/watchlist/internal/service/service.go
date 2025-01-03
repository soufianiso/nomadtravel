package service

import (
	"context"
	"log/slog"

	"github.com/soufianiso/nomadtravel/services/watchlist/api/v1/proto/watchlist"
	"github.com/soufianiso/nomadtravel/services/watchlist/internal/storage"
	// "user/internal/types"
)

type WatchlistService struct {
	watchlist.UnimplementedWatchlistServer
	log *slog.Logger 
	store storage.Store
}

func NewWachlistService(log *slog.Logger , store storage.Store) *WatchlistService{
	return &WatchlistService{
		log : log,
		store: store,

	}
}

func (s *WatchlistService) AddToWatchlist(
	ctx context.Context,
	req *watchlist.AddToWatchlistRequest,
	) (*watchlist.AddToWatchlistResponse, error) {

	userId := req.GetUserId()	
	movieId := req.GetMovieId()
	
	err := s.store.AddItemToWatchlist(userId, movieId)
	if err != nil{
		return nil, err
	}

	return  nil, nil
}








