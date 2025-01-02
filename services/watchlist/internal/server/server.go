package server

import (
	"database/sql"
	"log/slog"

	watchlistpb "github.com/soufianiso/nomadtravel/services/watchlist/api/v1/proto/watchlist"
	"github.com/soufianiso/nomadtravel/services/watchlist/internal/service"
	"github.com/soufianiso/nomadtravel/services/watchlist/internal/storage"
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

