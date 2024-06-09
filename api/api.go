package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
)

func Setup(host string) (*grpc.ClientConn, error) {
	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{},
		InsecureSkipVerify: true,
	})

	// Dial the gRPC server with the TLS credentials
	c, err := grpc.Dial(host, grpc.WithTransportCredentials(creds))
	if err != nil {
		return c, fmt.Errorf("error dialing git service host: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	for {
		s := c.GetState()
		if s == connectivity.Ready {
			break
		}
		if !c.WaitForStateChange(ctx, s) {
			if ctx.Err() == context.DeadlineExceeded {
				return c, fmt.Errorf("timeout: failed to connect to git service host within specified time")
			}
		}
	}

	return c, nil
}
