package main

import (
	"log"
	"net"

	datastore "github.com/aryannr97/data-server/internal/datastore/mongodb"
	"github.com/aryannr97/data-server/pkg/grpc/ping"
	"github.com/aryannr97/data-server/pkg/grpc/user"
	"github.com/aryannr97/data-server/pkg/repository"
	"google.golang.org/grpc"
)

func main() {
	store, err := datastore.Init()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
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

	userStore := repository.MongoUserRepository{DB: store}
	u := user.Server{UserStore: &userStore}
	user.RegisterUserServiceServer(server, &u)

	// Start to listen on GRPC server
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("Error while starting grpc server: %v", err)
	}
}
