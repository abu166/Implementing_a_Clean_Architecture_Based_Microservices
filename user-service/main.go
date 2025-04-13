package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "Implementing_a_Clean_Architecture_Based_Microservices/user-service/config"
    "Implementing_a_Clean_Architecture_Based_Microservices/user-service/database"
    "Implementing_a_Clean_Architecture_Based_Microservices/user-service/proto"
    "Implementing_a_Clean_Architecture_Based_Microservices/user-service/services"
)

// Embed UnimplementedUserServiceServer to satisfy the UserServiceServer interface
type UserServiceServer struct {
    proto.UnimplementedUserServiceServer
}

// RegisterUser implements the RegisterUser RPC method
func (s *UserServiceServer) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
    user, err := services.RegisterUser(req.Username, req.Password, req.Email)
    if err != nil {
        return nil, err
    }

    return &proto.RegisterUserResponse{
        Id:       int32(user.ID),
        Username: user.Username,
        Email:    user.Email,
    }, nil
}

// AuthenticateUser implements the AuthenticateUser RPC method
func (s *UserServiceServer) AuthenticateUser(ctx context.Context, req *proto.AuthenticateUserRequest) (*proto.AuthenticateUserResponse, error) {
    token, err := services.AuthenticateUser(req.Username, req.Password)
    if err != nil {
        return nil, err
    }

    return &proto.AuthenticateUserResponse{
        Token: token,
    }, nil
}

// GetUserProfile implements the GetUserProfile RPC method
func (s *UserServiceServer) GetUserProfile(ctx context.Context, req *proto.GetUserProfileRequest) (*proto.GetUserProfileResponse, error) {
    user, err := services.GetUserProfile(req.Id)
    if err != nil {
        return nil, err
    }

    return &proto.GetUserProfileResponse{
        Id:       int32(user.ID),
        Username: user.Username,
        Email:    user.Email,
    }, nil
}

func main() {
    config.LoadConfig()
    database.InitDB()

    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()

    // Register the UserServiceServer
    proto.RegisterUserServiceServer(s, &UserServiceServer{})

    // Enable gRPC reflection
    reflection.Register(s) // Add this line

    log.Println("User Service running on port 50053")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}