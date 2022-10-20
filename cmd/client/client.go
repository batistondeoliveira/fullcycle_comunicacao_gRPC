package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/batistondeoliveira/fullcycle_comunicacao_gRPC/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerboso(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}
	fmt.Println(res)
}

func AddUserVerboso(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}
	responseStream, err := client.AddUserVerboso(context.Background(), req)
	if err != nil {
		log.Fatalf("Could no make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could no receive the msg: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "e1",
			Name:  "Eliel 1",
			Email: "eli1@eli.com",
		},
		&pb.User{
			Id:    "e2",
			Name:  "Eliel 2",
			Email: "eli2@eli.com",
		},
		&pb.User{
			Id:    "e3",
			Name:  "Eliel 3",
			Email: "eli3@eli.com",
		},
		&pb.User{
			Id:    "e4",
			Name:  "Eliel 4",
			Email: "eli4@eli.com",
		},
		&pb.User{
			Id:    "e5",
			Name:  "Eliel 5",
			Email: "eli5@eli.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}
	fmt.Println(res)
}
