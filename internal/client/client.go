package client

import (
	"log"
	pb "user/api/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func NewAuthClient(addr string) pb.UserClient{
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	s := pb.NewUserClient(conn)
	return s
}
