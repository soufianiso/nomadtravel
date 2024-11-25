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
	r.Handle("/movies",handleMoviesByPage(logger, moviesService)).Methods("GET")
	r.Handle("/movie/{id}",handleMovieByPage(logger, moviesService)).Methods("GET")
}


func handleMoviesByPage(log *slog.Logger, movies moviespb.MoviesClient) http.Handler {
	log = log.With("handler","HandleMoviesByPage")

	return http.HandlerFunc(func(w http.ResponseWriter,  r *http.Request){
		log = log.With("requestID",r.Context().Value("requestID"))

		q := r.URL.Query().Get("page")
		idstr , err := strconv.ParseInt(q, 10, 32)
		if err != nil{
			log.Error("the page number incorrect")
			utils.Encode(w,r,500, errors.New("internal errors"))
			return 
		}

		page := int32(idstr)
		ctx, _ := context.WithTimeout(context.Background(),  time.Second * 3)
		res , err := movies.ListMovies(ctx, &moviespb.ListMoviesRequest{
			Page: page,
			PageSize: types.MovieSize,
		} )

		if err != nil{
			log.Error("Failed to retrieve movies", err)
			utils.Encode(w,r, 404, map[string]string{"msg":"no movies"})
			return 
		}

		err = utils.Encode(w, r, 200, res) 
		if err != nil{
			log.Error("Failed to Encode movies",  err)
			return
		}

		log.Info("list movies successfully")
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
