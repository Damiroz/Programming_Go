package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"tech-ip-sem2/services/auth/internal/service"
	pb "tech-ip-sem2/services/auth/pb/proto"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	svc *service.AuthService
}

func NewAuthServer() *AuthServer {
	return &AuthServer{
		svc: service.NewAuthService(),
	}
}

func (s *AuthServer) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	log.Println("Verify called with token:", req.Token)

	valid, subject := s.svc.ValidateToken(req.Token)
	log.Println("ValidateToken result:", valid, subject)

	if !valid {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	return &pb.VerifyResponse{
		Valid:   true,
		Subject: subject,
	}, nil
}

func RunGRPC(addr string) error {
	lis, err := net.Listen("tcp4", addr)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, NewAuthServer())

	log.Println("Auth gRPC server running on", addr)
	return server.Serve(lis)
}
