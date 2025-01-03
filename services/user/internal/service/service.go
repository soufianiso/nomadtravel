package service

import (
	"context"
	"log/slog"

	userpb "github.com/soufianiso/nomadtravel/user/api/v1/proto/user"
	"github.com/soufianiso/nomadtravel/user/internal/storage"
	// "user/internal/types"
)

type Userservice struct {
	userpb.UnimplementedUserServer
	log *slog.Logger
	// interface of storage
	store storage.Store
}

func NewUserService(log *slog.Logger, store storage.Store) *Userservice {
	return &Userservice{
		log:   log,
		store: store,
	}
}

func (s *Userservice) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	exists, err := s.store.CheckUserByEmail(req.GetEmail())
	if err != nil {
		s.log.Error("problem in databsase","Details",err)
		return nil, err
	}
	if exists {
		s.log.Error("email already exists","Details",err)
		return nil, err
	}

	hashed, err := hashPassword(req.GetPassword())

	if err != nil {

		s.log.Error("Failed to hash the password","Details",err)
	}
	err = s.store.CreateUser(req.Name, req.Email, hashed)
	if err != nil {
		s.log.Error("Failed to Create the user","Details",err)
		return nil, err
	}

	s.log.Info("RegisterUser", "Create the succesfully", req.GetName())
	return &userpb.RegisterUserResponse{Id: req.GetName()}, nil

}

func (s *Userservice) LoginUser(ctx context.Context, req *userpb.LoginUserRequest) (*userpb.LoginUserResponse, error) {
	email := req.GetEmail()

	hashedpassword, err := s.store.GetEmail(email)
	if err != nil {
		s.log.Error("failed to verify the email","Details",err)
		return nil, err
	}
	s.log.Info("found the email and its hashed password", "email", email,
		"hashedpassword", hashedpassword)

	err = validatePassword(hashedpassword, req.GetPassword())
	if err != nil {
		s.log.Error("failed to compared the hashed password against plain password","Details",err)
		return nil, err
	}

	s.log.Info("successed in comparing the hashedpassword against plain", "email", email)

	token, err := CreateToken(email)
	if err != nil {
		s.log.Error("Failed to Create a token","Details",err)
		return nil, err
	}

	s.log.Info("Create Token succesfully", "email", email)
	return &userpb.LoginUserResponse{
		Email: email,
		Token: token,
	}, nil
}
