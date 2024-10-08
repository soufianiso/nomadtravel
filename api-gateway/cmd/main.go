package main

import (
	"flag"
	"fmt"
	// "net"
	// "log"
	pb "api-gateway/api/proto/user"
	"api-gateway/internal/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// "fmt"
	"api-gateway/configs"
	"log"
	"net/http"
)

var (
	GatewayAddr = flag.String("address of gatewayservice ",configs.Envs.ApiGatewayPort,"this service used to fetch data")
	userAddr = flag.String("address of grcpUserMicro ",configs.Envs.UserMicroservicePort,"this is server grpc server of the user microservice")
)


func main(){
	// Set up a connection to the server.
	conn, err := grpc.NewClient(fmt.Sprintf(":%s",*userAddr), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	userMicroClient := pb.NewUserClient(conn)

	s := server.NewGatewayServer(userMicroClient)

	app := http.Server{
		Addr: fmt.Sprintf(":%s",*GatewayAddr),
		Handler: s,
	}

	log.Printf("api gateway listening to addresss:%s", *GatewayAddr)
	app.ListenAndServe()
}

