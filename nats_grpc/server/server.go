package main

import (
	"fmt"
	"log"
	"net"

	pro "nats_grpc/proto"
	"nats_grpc/service"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pro.RegisterFileTransferServiceServer(s, &service.Server{})

	fmt.Println("Server is running on :50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
