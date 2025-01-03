package middlewares

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/utils"
)




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


func LoggingMiddleware(log *slog.Logger, next http.Handler) http.Handler{

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

func InjectIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()

		ctx := context.WithValue(r.Context(), "requestID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}



func AuthMiddleware(log *slog.Logger, next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		token := r.Header.Get("Authorization")
		err, claims := utils.ParseClaims(token)  
		if err != nil {
			log.Error("Failed to signed the token","Details",err)
			w.WriteHeader(http.StatusNotFound)
			return 
		}

		ctx := context.WithValue(r.Context(), "email", claims.Email)

		next.ServeHTTP(w,r.WithContext(ctx))




		

	})
}



