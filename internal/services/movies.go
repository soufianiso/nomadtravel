package services

import (
	"net/http"
	"github.com/gorilla/mux"	
	"log/slog"
)






func SetMovies(r *mux.Router, logger *slog.Logger) {
	r.Handle("movies",listMoviesHandler(logger)).Methods("GET")
	r.Handle("movies/{id}",listMovieByIdHandler(logger)).Methods("GET")
}


func listMoviesHandler(log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,  r *http.Request){


	})
}


func listMovieByIdHandler(log *slog.Logger) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){


	})
}

