package services

// import (
// 	"os"

// 	"github.com/joho/godotenv"
// 	"net/http"
// 	"github.com/soufianiso/nomadtravel/api-gateway/internal/types"
// 	"github.com/soufianiso/nomadtravel/api-gateway/internal/utils"
// 	"github.com/gorilla/mux"
// 	"log/slog"
// )

// func SetAuth(r *mux.Router, logger *slog.Logger, userService *userpb.UserClient) {
// 	// r.Handle("movies",handleListMovies(logger, moviesService)).Methods("GET")
// 	r.Handle("/login",handleMoviesByPage(logger, userService)).Methods("GET")
// 	r.Handle("/register",handleMovieByPage(logger, userService)).Methods("GET")
// }

// func handleLogin(logger *slog.Logger,  userService) http.Handler{
// 	godotenv.Load()
// 	jwtsecret := os.Getenv("jwtsecret")
	
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var user types.LoginUser

// 		if err := utils.Decode(r, &user) ; err != nil{
// 			logger.Printf("Failed to decode request body:%v",err)
// 			return
// 		}

// 		if err := utils.Validate(&user); err != nil{
// 			logger.Printf("validation failed: %s", err)
// 			utils.Encode(w, r, http.StatusBadRequest, utils.ApiError{ Error: "email or password is incorrect" })
// 			return 
// 		}

// 		u , err := storage.GetUserByEmail(user.Email)
// 		if err !=  nil{
// 			utils.Encode(w, r, http.StatusBadRequest, utils.ApiError{ Error: "email or password is incorrect" })
// 			logger.Printf("Failed to retrieve user by Email (%s)",user.Email)
// 			return 
// 		}

// 		if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
// 			logger.Printf("Password mismatch for the user: (%s)",user.Email)
// 			utils.Encode(w, r, http.StatusBadRequest, utils.ApiError{ Error: "email or password is incorrect" })
// 			return 
// 		}
		
// 		tokenString, err := auth.Createjwt(user.Email, jwtsecret)
// 		if err != nil{
// 			logger.Printf("failed to create JWT token: %v",err)
// 			return 
// 		}

// 		err = utils.Encode(w, r, http.StatusCreated, map[string]string{"Authorization": tokenString}) 
// 		if err != nil{
// 			logger.Printf("failed to encode the response: %v",err)
// 		}
// 	})
// }
