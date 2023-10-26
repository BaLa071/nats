package main

import (
	"fmt"
	"log"
	"net"
	pro "obj_store/proto"
	"obj_store/service"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pro.RegisterObjGrpcServiceServer(s, &service.Server{})
	fmt.Println("Server is running on : 8080")

	if err := s.Serve(listen); err!=nil{
		log.Fatalf("Failed to serve: %v", err)
	}

}
