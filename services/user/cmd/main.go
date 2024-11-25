package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"net"

	_ "github.com/lib/pq"
	"github.com/soufianiso/nomadtravel/user/configs"
	"github.com/soufianiso/nomadtravel/user/internal/server"
	"google.golang.org/grpc"
)

var (
	serverAdd = flag.String("addrauth", configs.Envs.UserMicroservicePort, "The server port")

	DBUser = flag.String("DBUser", configs.Envs.DBName, "postgres database user")
	DBName = flag.String("DBName", configs.Envs.DBPassword, "postgres database name")
	DBPassword = flag.String("DBPassword", configs.Envs.DBPassword, "postgres database password")
	DBHost = flag.String("DBHost", configs.Envs.DB_HOST, "postgres database host")
	DBPort= flag.String("DBPort", configs.Envs.DB_PORT, "postgres database port")
)

func main(){
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

// grpc listening
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s",*serverAdd))
	defer lis.Close()
	if err != nil{
		slog.Error("", "failed to listen:", err)
	}
	logger.Info("grpc","grpc listening: ", *serverAdd)


// Postgres Database 
	conn := fmt.Sprintf("host=%s port=%s user=%s " +
	    "password=%s dbname=%s sslmode=disable",
	    *DBHost, *DBPort, *DBUser, *DBPassword, *DBName)

	db, err := sql.Open("postgres",conn)	
	defer db.Close()
	if err != nil{
		logger.Error("", "postgres connection", err)
	}

	if err := db.Ping() ; err != nil{
		logger.Error("", "postgres ping", err)
	}
	logger.Info("postgres listening", "listen on" , *DBPort)


	grpcConn := grpc.NewServer()
	s := server.NewServer(logger, db, grpcConn)
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func(){
		defer wg.Done()
		if err := s.Serve(lis); err != nil{
			logger.Error("","failed to server", err)
		}
	}()

	wg.Wait()
}











