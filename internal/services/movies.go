package services

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	moviespb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/movies"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/types"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/utils"
)

func SetMovies(r *mux.Router, logger *slog.Logger, moviesService moviespb.MoviesClient) {
	// r.Handle("movies",handleListMovies(logger, moviesService)).Methods("GET")
	r.Handle("/movies/{id}",handleMoviesByPage(logger, moviesService)).Methods("GET")
	r.Handle("/movie/{id}",handleMovieByPage(logger, moviesService)).Methods("GET")
}


func handleMoviesByPage(log *slog.Logger, movies moviespb.MoviesClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,  r *http.Request){
		vars := mux.Vars(r)
		id := vars["id"]
		idstr , err := strconv.ParseInt(id, 10, 32)
		if err != nil{
			log.Error("the id page of movies is not an integer","Error",err)
			utils.Encode(w,r,500, errors.New("internal errors"))
		}

		page := int32(idstr)
		ctx, _ := context.WithTimeout(context.Background(),  time.Second * 3)
		res , err := movies.ListMovies(ctx, &moviespb.ListMoviesRequest{
			Page: page,
			PageSize: types.MovieSize,
		} )

		if err != nil{
			log.Error("Failed to retireve movies from MoviesMicroService", "Error", err)
		}

		if err := utils.Encode(w, r, 200, res ) ; err != nil{
			log.Error("Failed to Encode movies", "Error", err)
		}

	})
}

// func handleMoviesByPage(log *slog.Logger,  movies moviespb.MoviesClient) http.Handler{
// 	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){


// 	})
// }


func handleMovieByPage(log *slog.Logger,  movies moviespb.MoviesClient) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){


	})
}
