package main

import (
	"log"
	"net"

	"github.com/aryannr97/data-server/pkg/grpc/ping"
	"google.golang.org/grpc"
)

func main() {
	// Create a TCP listerner
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Error listening on port %v: %v", ":9090", err)
	}

	// Create new GRPC server
	server := grpc.NewServer()

	// Register grpc services
	p := ping.Server{}
	ping.RegisterPingServiceServer(server, &p)

	// Start to listen on GRPC server
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("Error while starting grpc server: %v", err)
	}
}
