package grpc

import (
	"context"
	"fmt"
	ssov1 "github.com/bezhan2009/AuthProtos/gen/go/sso"
	"google.golang.org/grpc/codes"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	authApi ssov1.AuthClient
	pingApi ssov1.PingServiceClient
}

var client *Client

func New(ctx context.Context,
	addr string,
	timeout time.Duration,
	retriesCount int,
) error {
	const op = "grpc.New"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	conn, err := grpc.DialContext(ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpcretry.UnaryClientInterceptor(retryOpts...)),
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	client = &Client{
		authApi: ssov1.NewAuthClient(conn),
		pingApi: ssov1.NewPingServiceClient(conn),
	}

	return nil
}

func GetClient() *Client {
	return client
}
