package server

import (
	"context"

	authpb "github.com/bag-huyag/Auth-Service/proto"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
	users map[string]string
}

func NewAuthServer() *AuthServer {
	return &AuthServer{
		users: make(map[string]string),
	}
}

func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.AuthResponse, error) {
	if _, exists := s.users[req.Email]; exists {
		return &authpb.AuthResponse{Message: "User already exists"}, nil
	}
	s.users[req.Email] = req.Password
	token := "dummy-token-for-" + req.Email
	return &authpb.AuthResponse{Token: token, Message: "Registered successfully"}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.AuthResponse, error) {
	password, exists := s.users[req.Email]
	if !exists || password != req.Password {
		return &authpb.AuthResponse{Message: "Invalid credentials"}, nil
	}
	token := "dummy-token-for-" + req.Email
	return &authpb.AuthResponse{Token: token, Message: "Login successful"}, nil
}
