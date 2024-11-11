package services

import (
	moviespb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/movies"
	"net/http"
	"github.com/gorilla/mux"	
	"log/slog"
)






func SetMovies(r *mux.Router, logger *slog.Logger, moviesService moviespb.MoviesClient) {
	r.Handle("movies",handleListMovies(logger, moviesService)).Methods("GET")
	r.Handle("movies/{id}",handleGetMovieById(logger, moviesService)).Methods("GET")
}


func handleListMovies(log *slog.Logger, movies moviespb.MoviesClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,  r *http.Request){


	})
}

func handleGetMovieById(log *slog.Logger,  movies moviespb.MoviesClient) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){


	})
}

