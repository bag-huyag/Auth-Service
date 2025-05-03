package main

import (
	"log"
	"net"

	"github.com/bag-huyag/Auth-Service/internal/server"
	authpb "github.com/bag-huyag/Auth-Service/proto"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, server.NewAuthServer())
	reflection.Register(s)

	log.Println("Auth Service running on :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
