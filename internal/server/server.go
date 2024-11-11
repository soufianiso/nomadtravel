package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/services"
	pb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/user"
)




func NewGatewayServer(userService pb.UserClient ) http.Handler{
	router := mux.NewRouter()
	router = router.PathPrefix("/api/v1").Subrouter()
	services.SetUser(router, userService)		
	services.SetMovies(router)		
	
	var handler http.Handler
	handler = router

	return handler 

}	





