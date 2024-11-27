package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	moviespb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/movies"
	userpb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/user"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/services"
	// "github.com/soufianiso/nomadtravel/api-gateway/internal/utils"
)




func NewGatewayServer(log *slog.Logger, userService userpb.UserClient, moviesService moviespb.MoviesClient) http.Handler{
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	services.SetUser(apiRouter, log, userService)		
	services.SetMovies(apiRouter, log, moviesService)		
	// services.SetAuth(apiRouter, log, moviesService)		
	

	var handler http.Handler

	// Top Level Middlewares
	handler = CORSMiddleware(apiRouter)
	handler = loggingMiddleware(log , handler)
	handler = InjectRequestIDMiddleware(handler)

	return handler 

}	

func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}


func loggingMiddleware(log *slog.Logger, next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		requestID := r.Context().Value("requestID")
		if requestID == nil{

			requestID = "unknown"
		}

		log.Info("Request received", "requestID", requestID ,"method", r.Method, "url", r.URL.String())

		next.ServeHTTP(w, r)

		log.Info("Request processed", "requestID", requestID,  "duration in ms", time.Since(start).Milliseconds())
	})

}

func InjectRequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()

		ctx := context.WithValue(r.Context(), "requestID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}





