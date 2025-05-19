package main

import (
	"fmt"
	"sample/api/object/mygrpc"
	"sample/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("", "")
	if err != nil {
		fmt.Println("Failed to create tls credentials", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	listner, backend := server.OpenConnection(":5566")
	mygrpc.RegisterMyServiceServer(grpcServer, backend)
	if grpcServer.Serve(listner); err != nil {
		fmt.Println("Failed to server", err)
	}
}
