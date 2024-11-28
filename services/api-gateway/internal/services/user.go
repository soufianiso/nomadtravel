package services

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	userpb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/user"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/types"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/utils"
)

func SetUser(r *mux.Router, log *slog.Logger, user userpb.UserClient) {
	r.Handle("/register", handleRegister(log, user)).Methods("POST")
	r.Handle("/login", handleLogin(log, user)).Methods("POST")
}

func handleRegister(log *slog.Logger, user userpb.UserClient) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u types.UserRegister
		if err := utils.Decode(r, &u); err != nil {
			log.Error("Failed to Decode json fields", "error", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()


		res, err := user.RegisterUser(ctx, &userpb.RegisterUserRequest{
			Name: u.Name,
			Email:     u.Email,
			Password:  u.Password,
		})
		if err != nil {
			log.Error("message from the microservice: %v", err)
			return 
		}

		log.Info("registration success", "requestID", r.Context().Value("requestID") ,"url", r.URL.String(), "user", res.GetId())
		utils.Encode(w, r, 200, fmt.Sprintf("success", res.GetId()) )

	})
}

func handleLogin(log *slog.Logger, user userpb.UserClient) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u types.UserLogin

		if err := utils.Decode(r, &u) ; err != nil{
			log.Error("Failed to Decode json fields", "error", err)
			utils.Encode(w, r, 200, map[string]string{"Error": "incorrect email or password"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()

		res, err := user.LoginUser(ctx, &userpb.LoginUserRequest{
			Email: u.Email,
			Password: u.Password,
		})

		if err != nil{
			// failed to retrieve the jwt token
			log.Error("message from the microservice: %v", "Error", err)
			utils.Encode(w, r, 200, map[string]string{"Error": "incorrect email or password"})
			return 
		}

		log.Info("Created Token succefully", "requestID", r.Context().Value("requestID") ,"url", r.URL.String(), "email", res.GetEmail(),
			"token", res.GetToken())
		utils.Encode(w, r, 200, map[string]string{"Authorization": res.GetToken()} )



	})
}
