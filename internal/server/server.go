package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"api-gateway/internal/services"
	pb "api-gateway/api/proto/user"
)




func NewGatewayServer(userService pb.UserClient ) http.Handler{
	router := mux.NewRouter()
	router = router.PathPrefix("/api/v1").Subrouter()
	services.SetUser(router, userService)		
	
	var handler http.Handler
	handler = router

	return handler 

}	





