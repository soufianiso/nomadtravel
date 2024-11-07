package service

import (
	"context"
	"log/slog"
	pb "user/api/proto/user"
	"user/internal/storage"
	// "user/internal/types"
)

type Userservice struct {
	pb.UnimplementedUserServer
	log *slog.Logger 
	// interface of storage
	store storage.Store
	userClient pb.UserClient
}




func NewUserService(log *slog.Logger , store storage.Store, userClient pb.UserClient ) *Userservice{
	return &Userservice{
		log : log,
		store: store,
		userClient: userClient,

	}
}

func (s *Userservice) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	err := s.store.CreateUser(req)
	if err !=  nil{
		s.log.Info("RegisterUser", "Failed to retrieve user by Email (%s)", req.GetFirstName())
		s.log.Error("RegisterUser", "the error", err)
	}

	s.log.Info("RegisterUser", "Create the succesfully", req.GetFirstName())
	return  &pb.RegisterUserResponse{Id: req.GetFirstName()}, nil

}









