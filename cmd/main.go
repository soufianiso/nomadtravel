package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"github.com/soufianiso/nomadtravel/movies/configs"
	"github.com/soufianiso/nomadtravel/movies/internal/server"
	"net"
	"os"
	"sync"

	// "time"

	// "movies/internal/server"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var (
	serverAddr  = flag.String("moviesAddr", configs.Envs.MoviesMicroservicePort, "Address for the Movies microservice")
	userAddr    = flag.String("userAddr", configs.Envs.UserMicroservicePort, "Address for the User microservice")
	postgresHost =  flag.String("postgresAddr",configs.Envs.DB_HOST, "postgres host")
	postgresName =  flag.String("postgresName",configs.Envs.DBName, "postgres name")
	postgresPort =  flag.String("postgresPort",configs.Envs.DB_PORT, "postgres port")
	postgresUser =  flag.String("postgresUser",configs.Envs.DBUser, "postgres User")
	postgresPassword =  flag.String("postgesPassword",configs.Envs.DBPassword, "postgres Password")

)

func main(){
	
	flag.Parse()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))


	// litening on port 
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s",*serverAddr))
	if err != nil{
		log.Error("error to open connection maybe the port is being used","Error",err)
	}
	

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		*postgresHost, *postgresPort, *postgresUser, *postgresPassword, *postgresName)

	db, err := sql.Open("postgres",conn)
	if err != nil{
		log.Error("can't connect to database s","Error",err)
		os.Exit(1)
	}

	if err = db.Ping() ;  err != nil{
		log.Error("can't ping postgres",)
	}

	log.Info("postgres connected successfully on port ", "success", fmt.Sprintf("%s",*postgresPort))

	grpcConn := grpc.NewServer()
	s := server.NewServer(log, db, grpcConn) 

 	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		defer wg.Done()
		log.Info("server listening on port","server Address",*serverAddr)
		if err := s.Serve(lis); err != nil{
			log.Error("Failed to serve","Error",err)
		}
	}()
	wg.Wait()

}












