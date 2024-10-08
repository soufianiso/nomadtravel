package server

import (
	"database/sql"
	"log/slog"
	pb "user/api/proto/user"
	"user/internal/service"
	"user/internal/storage"

	"google.golang.org/grpc"
)


func NewServer(log *slog.Logger , db *sql.DB, authClient pb.UserClient ) *grpc.Server{
	s := grpc.NewServer()
	userStore := storage.NewUserStorage(db)
	userservice := service.NewUserService(log, userStore, authClient)

	pb.RegisterUserServer(s, userservice)
	return s  
}

