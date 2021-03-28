package main

import (
	"context"
	"log"

	"github.com/aryannr97/data-server/pkg/grpc/ping"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error establishing connection with grpc server: %v", err)
	}

	defer conn.Close()

	client := ping.NewPingServiceClient(conn)

	res, err := client.Ping(context.Background(), &ping.PingRequest{Ping: "status"})
	if err != nil {
		log.Printf("Error returned: %v", err)
	}

	log.Printf("Server Status: %v", res.Status)
}
