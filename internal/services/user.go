package services

import (
	pb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/user"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/types"
	"context"
	"fmt"
	"log"
	"net/http"
	// "sync"
	"time"

	"github.com/soufianiso/nomadtravel/api-gateway/internal/utils"
	"github.com/gorilla/mux"
)


func SetUser(r *mux.Router, user pb.UserClient) {	
	r.Handle("/register",handleRegister(user)).Methods("POST")
}

func handleRegister(user pb.UserClient) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u types.User
		if err := utils.Decode(r, u) ; err != nil{
			fmt.Println("error")
			return 
		}
		log.Println("the step of verification works")
		 
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		
		res , err := user.RegisterUser(ctx, &pb.RegisterUserRequest{
			FirstName: u.FirstName ,
			LastName: u.LastName, 
			Email: u.Email, 
			Password: u.Password,
		})
		if err != nil{
			fmt.Println(err)
		}
		utils.Encode(w,r ,200, res.GetId())
	})
}
