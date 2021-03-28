package ping

import (
	"context"
	"log"
)

// Server handles server status
type Server struct{}

// Ping serves grpc running status
func (p *Server) Ping(ctx context.Context, in *PingRequest) (*PingResponse, error) {
	log.Println("Server status requested")
	return &PingResponse{Status: "UP"}, nil
}

func (p *Server) mustEmbedUnimplementedPingServiceServer() {}
