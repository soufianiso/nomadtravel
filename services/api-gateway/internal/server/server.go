package server

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	moviespb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/movies"
	userpb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/user"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/services"
	m "github.com/soufianiso/nomadtravel/api-gateway/internal/middlewares"
	// "github.com/soufianiso/nomadtravel/api-gateway/internal/utils"
)




func NewGatewayServer(log *slog.Logger, userService userpb.UserClient, moviesService moviespb.MoviesClient) http.Handler{
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	services.SetUser(apiRouter, log, userService)		
	services.SetMovies(apiRouter, log, moviesService)		
	services.SetWatchlist(apiRouter, log, moviesService)		
	

	var handler http.Handler

	// Top Level Middlewares
	handler = m.CORSMiddleware(apiRouter)
	handler = m.LoggingMiddleware(log , handler)
	handler = m.InjectIDMiddleware(handler)

	return handler 

}	

