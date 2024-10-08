package services

import (
	pb "api-gateway/api/proto/user"
	"api-gateway/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func SetUser(r *mux.Router, user pb.UserClient) {	
	r.Handle("/register",handleRegister(user)).Methods("POST")
}

func handleRegister(user pb.UserClient) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u types.User
		json.NewDecoder(r.Body).Decode(&u)

		log.Println("step ov verification works")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res , err := user.RegisterUser(ctx, &pb.RegisterUserRequest{FirstName: u.FirstName ,LastName: u.LastName, Email: u.Email, Password: u.Password})
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("succeess")
		json.NewEncoder(w).Encode(res.GetId())
	
	})
}
