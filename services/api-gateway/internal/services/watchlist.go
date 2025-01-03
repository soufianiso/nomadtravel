package services

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	moviespb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/movies"

	// "github.com/soufianiso/nomadtravel/api-gateway/internal/types"
	m "github.com/soufianiso/nomadtravel/api-gateway/internal/middlewares"
	// "github.com/soufianiso/nomadtravel/api-gateway/internal/utils"
)

func SetWatchlist(r *mux.Router, logger *slog.Logger, moviesService moviespb.MoviesClient) {
	r.Handle("/watchlist", m.AuthMiddleware(logger, handleAddToWatchlist(logger))).Methods("POST")
}

func handleAddToWatchlist(logger *slog.Logger) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		

	})

}



