package middlewares


import (
	"context"
	"net/http"
	"time"
	"log/slog"
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

func InjectRequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()

		ctx := context.WithValue(r.Context(), "requestID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}





