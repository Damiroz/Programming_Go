package grpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	pb "tech-ip-sem2/services/auth/pb/proto"
)

type AuthClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(addr string) (*AuthClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &AuthClient{client: pb.NewAuthServiceClient(conn)}, nil
}

func (a *AuthClient) Verify(ctx context.Context, token string) (bool, string, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	resp, err := a.client.Verify(ctx, &pb.VerifyRequest{Token: token})
	if err != nil {
		st, _ := status.FromError(err)
		return false, "", st.Err()
	}

	return resp.Valid, resp.Subject, nil
}
