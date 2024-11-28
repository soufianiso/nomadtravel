package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	moviespb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/movies"
	userpb "github.com/soufianiso/nomadtravel/api-gateway/api/v1/proto/user"
	"github.com/soufianiso/nomadtravel/api-gateway/internal/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"net/http"

	"github.com/soufianiso/nomadtravel/api-gateway/configs"
)

sasa
var (
	//this is in jwt branche
	gatewayAddr = flag.String("gatewayAddr", configs.Envs.ApiGatewayPort, "Address for the API Gateway service")
	userAddr    = flag.String("userAddr", configs.Envs.UserMicroservicePort, "Address for the User microservice")
	moviesAddr  = flag.String("moviesAddr", configs.Envs.MoviesMicroservicePort, "Address for the Movies microservice")
)

func main(){
	flag.Parse()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	userConn, err := grpc.NewClient(fmt.Sprintf(":%s",*userAddr), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Can't init gRPC client", "Client", "User", "Error",err)
	}
	defer userConn.Close()

	moviesConn, err := grpc.NewClient(fmt.Sprintf(":%s",*moviesAddr), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Can't init gRPC client", "Client", "Movies", "Error",err)
	}
	defer moviesConn.Close()

	// Create gRPC Clients
	userMicroClient := userpb.NewUserClient(userConn)
	moviesMicroClient := moviespb.NewMoviesClient(moviesConn)
	
	// Initialize the Api Gateway Server
	s := server.NewGatewayServer(log, userMicroClient, moviesMicroClient)

	// HTTP server setup
	app := http.Server{
		Addr: fmt.Sprintf(":%s",*gatewayAddr),
		Handler: s,
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Start HTTP server 
	go func() {
		defer wg.Done()
		log.Info("HTTP server starting", "address", app.Addr)
		if err := app.ListenAndServe() ; err != nil{
		}
	}()

	// Graceful shutdown 
	go func(){
		defer wg.Done()
		<-ctx.Done()
		log.Info("shutting down server...", "address", app.Addr)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Gracefully shut down the server
		if err := app.Shutdown(shutdownCtx); err != nil {
			log.Error("can't shutdown the server", "address", app.Addr)
		}

		log.Info("the Server Shute Down Successfully", "address", app.Addr)
	}()

	wg.Wait()
}

