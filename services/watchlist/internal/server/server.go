package server

import (
	"database/sql"
	"log/slog"

	watchlistpb "github.com/soufianiso/nomadtravel/watchlist/api/watchlist"
	"github.com/soufianiso/nomadtravel/watchlist/internal/service"
	"github.com/soufianiso/nomadtravel/watchlist/internal/storage"
	"google.golang.org/grpc"
)

func NewServer(log *slog.Logger ,db *sql.DB, grpcServer *grpc.Server) *grpc.Server{
	// initialize new storage
	moviesStore := storage.NewMoviesStorage(db)

	// initiliaze new movies service 
	moviesService := service.NewMoviesService(log, moviesStore)

	// watchlistpb.RegisterWachlitServer here

	return grpcServer
}

