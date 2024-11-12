package server

import (
	"database/sql"
	"log/slog"

	moviespb "github.com/soufianiso/nomadtravel/movies/api/v1/proto/movies"
	"github.com/soufianiso/nomadtravel/movies/internal/service"
	"github.com/soufianiso/nomadtravel/movies/internal/storage"
	"google.golang.org/grpc"
)

func NewServer(log *slog.Logger ,db *sql.DB, grpcServer *grpc.Server) *grpc.Server{
	// initialize new storage
	moviesStore := storage.NewMoviesStorage(db)

	// initiliaze new movies service 
	moviesService := service.NewUserService(log, moviesStore)

	moviespb.RegisterMoviesServer(grpcServer, moviesService)

	return grpcServer
}

