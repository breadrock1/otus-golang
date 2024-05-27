package grpcserv

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"time"
)

func CustomLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("method %q failed: %s", info.FullMethod, err)
	}

	ip := ""
	if peerAddr, ok := peer.FromContext(ctx); ok {
		ip = peerAddr.Addr.String()
	}

	var userAgent []string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		userAgent = md.Get("user-agent")
	}

	log.WithField("IP", ip).
		WithField("Method", info.FullMethod).
		WithField("User-Agent", userAgent).
		WithField("Latency", time.Since(start)).
		Info("GRPC request processed")

	return resp, err
}
