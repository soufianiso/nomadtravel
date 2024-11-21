package server

import (
	"database/sql"
	"log/slog"
	userpb "github.com/soufianiso/nomadtravel/user/api/v1/proto/user"
	"github.com/soufianiso/nomadtravel/user/internal/service"
	"github.com/soufianiso/nomadtravel/user/internal/storage"
	"google.golang.org/grpc"
)


func NewServer(log *slog.Logger , db *sql.DB , grpcServer *grpc.Server) *grpc.Server{
	userStore := storage.NewUserStorage(db)

	userService := service.NewUserService(log, userStore)

	userpb.RegisterUserServer(grpcServer, userService)
	return grpcServer  
}

