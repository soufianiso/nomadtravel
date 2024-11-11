package server

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	moviespb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/movies"
	userpb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/user"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/services"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/utils"
)




func NewGatewayServer(log *slog.Logger, userService userpb.UserClient, moviesService moviespb.MoviesClient) http.Handler{
	router := mux.NewRouter()
	router = router.PathPrefix("/api/v1").Subrouter()

	services.SetUser(router, log, userService)		
	services.SetMovies(router, log, moviesService)		
	
	var handler http.Handler

	// Top Level Middlewares
	handler = utils.CORSMiddleware(router)

	return handler 

}	





