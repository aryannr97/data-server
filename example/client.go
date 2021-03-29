package main

import (
	"context"
	"log"

	"github.com/aryannr97/data-server/pkg/grpc/ping"
	"github.com/aryannr97/data-server/pkg/grpc/user"
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
	log.Printf("Server Status: %v Error: %v", res, err)

	userClient := user.NewUserServiceClient(conn)

	res1, err := userClient.AddUser(context.Background(), &user.UserAddRequest{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john@mail.com",
		Phone:     "4567891245",
		Country:   "US",
	})
	log.Printf("AddUser Response: %v Error: %v", res1, err)

	res2, err := userClient.GetUser(context.Background(), &user.UserGetReuqest{Id: res1.Id})
	log.Printf("GetUser Response: %v Error: %v", res2, err)

	_, err = userClient.UpdateUser(context.Background(), &user.UserUpdateRequest{
		Id:        res1.Id,
		Firstname: "Mike",
		Lastname:  "Doe",
		Email:     "mike@mail.com",
		Phone:     "4567891245",
		Country:   "US",
	})
	log.Printf("UpdateUser Error: %v", err)

	res2, err = userClient.GetUser(context.Background(), &user.UserGetReuqest{Id: res1.Id})
	log.Printf("GetUser after update Response: %v Error: %v", res2, err)
	_, err = userClient.DeleteUser(context.Background(), &user.UserDeleteRequest{Id: res1.Id})
	log.Printf("DeleteUser Error: %v", err)
	res2, err = userClient.GetUser(context.Background(), &user.UserGetReuqest{Id: res1.Id})
	log.Printf("GetUser after delete Response: %v Error: %v", res2, err)
}
