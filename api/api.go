package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
)

func Setup(host string) (*grpc.ClientConn, error) {
	// Get the directory of the executable
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("failed to get executable path: %v", err)
	}

	// Determine the root directory of the project (assuming the executable is in the root)
	rootDir := filepath.Dir(exePath)

	// Construct the full paths to the TLS certificate and key
	certFile := filepath.Join(rootDir, "tls.crt")
	keyFile := filepath.Join(rootDir, "tls.key")

	// Load the client certificates from disk
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load key pair: %v", err)
	}

	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true, // This will accept any server certificate; not recommended for production
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
