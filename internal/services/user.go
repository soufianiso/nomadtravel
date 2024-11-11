package services

import (
	"context"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"time"

	userpb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/user"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/types"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/utils"
)

func SetUser(r *mux.Router, log *slog.Logger, user userpb.UserClient) {
	r.Handle("/register", handleRegister(log, user)).Methods("POST")
}

func handleRegister(log *slog.Logger, user userpb.UserClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u types.User
		if err := utils.Decode(r, u); err != nil {
			log.Error("Failed to Decode json fields", "error", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := user.RegisterUser(ctx, &userpb.RegisterUserRequest{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Password:  u.Password,
		})
		if err != nil {
			log.Error("error in the calling microsercvcie this is the error: %v", err)
		}
		utils.Encode(w, r, 200, res.GetId())
	})
}
