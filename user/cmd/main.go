package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"log/slog"

	// "log/slog"
	"net"
	"os"
	"user/configs"
	"user/internal/client"
	"user/internal/server"

	_ "github.com/lib/pq"
)

var (
	serverAdd = flag.String("addrauth", configs.Envs.UserMicroservicePort, "The server port")
	authAdd = flag.String("authAdd", configs.Envs.AuthMicroservicePort, "Auth Microservice port")

	DBUser = flag.String("DBUser", configs.Envs.DBName, "postgres database user")
	DBName = flag.String("DBName", configs.Envs.DBPassword, "postgres database name")
	DBPassword = flag.String("DBPassword", configs.Envs.DBPassword, "postgres database password")
	DBHost = flag.String("DBHost", configs.Envs.DB_HOST, "postgres database host")
	DBPort= flag.String("DBPort", configs.Envs.DB_PORT, "postgres database port")
)

func main(){
	flag.Parse()

// logging
	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, handlerOpts))
	slog.SetDefault(logger)

// grpc listening
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s",*serverAdd))
	defer lis.Close()
	if err != nil{
		slog.Error("", "failed to listen:", err)
	}
	slog.Info("grpc listening", "Listen on",*serverAdd)

// Postgres Database 
	conn := fmt.Sprintf("host=%s port=%s user=%s " +
	    "password=%s dbname=%s sslmode=disable",
	    *DBHost, *DBPort, *DBUser, *DBPassword, *DBName)

	db, err := sql.Open("postgres",conn)	
	defer db.Close()
	if err != nil{
		slog.Error("", "postgres connection", err)
	}

	if err := db.Ping() ; err != nil{
		slog.Error("", "postgres ping", err)
	}
	slog.Info("postgres listening", "listen on" , *DBPort)

	// Grpc Auth client
	authClient := client.NewAuthClient(string(*authAdd))

	// Grpc User Server
	s := server.NewServer(logger, db, authClient)
	if err := s.Serve(lis); err != nil{
		log.Fatalf("failed to server %v ", err)
	}









}
