package main

import (
	"context"
	"fmt"
	"log"

	authpb "github.com/bag-huyag/Auth-Service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	res, err := client.Register(context.Background(), &authpb.RegisterRequest{
		Email:    "test@mail.com",
		Password: "123",
	})
	fmt.Printf("Register Response: %v\nError: %v\n\n", res, err)

	// Test Login
	loginRes, err := client.Login(context.Background(), &authpb.LoginRequest{
		Email:    "test@mail.com",
		Password: "123",
	})
	fmt.Printf("Login Response: %v\nError: %v\n", loginRes, err)
}
