package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"movies/configs"
	"net"
	"os"
	"sync"

	// "time"

	// "movies/internal/server"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var (


	postgresHost =  flag.String("postgres onnection",configs.Envs.DB_HOST, "postgres host")
	postgresName =  flag.String("postgres nnection",configs.Envs.DBName, "postgres name")
	postgresPort =  flag.String("postgres nection",configs.Envs.DB_PORT, "postgres port")
	postgresUser =  flag.String("postgres ion",configs.Envs.DBUser, "postgres User")
	postgresPassword =  flag.String("postges connection",configs.Envs.DBPassword, "postgres Password")
	serverPort =  flag.String("server port",configs.Envs.MoviesMicroservicePort, "Movie microservice port")

)

func main(){
	
	flag.Parse()
	log := slog.Default()


	// litening on port 
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s",*serverPort))
	if err != nil{
		log.Error("ss",err)
	}

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		*postgresHost, *postgresPort, *postgresUser, *postgresPassword, *postgresName)

	db, err := sql.Open("postgres",conn)
	if err != nil{
		log.Error("", "can't connect to database s",err)
		os.Exit(1)
	}

	if err = db.Ping() ;  err != nil{
		fmt.Println("can't ping ")
		os.Exit(1)
	}


	
	// not yet created just a pseodocode
	// server its the struct that has the unimplemented proto buf pb.Unimplemented
	server := server.NewServer()

	grpcConn := grpc.NewServer()

	pb.RegisterMoviesServer(grpcConn,server)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		if err := grpcConn.Serve(lis) ; err != nil{
			log.Error("","can't serve",err)
		}
	}()
	


	// Now I need to established a grpc connection

	
	// app := server.NewServer(log,db)
}












