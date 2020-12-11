package main

import (
	"context"
	"net"

	pb "github.com/mvrilo/presentations/go-protobuf-microservices/generated"
	"google.golang.org/grpc"
)

type messageService struct {
	pb.UnimplementedMessageServiceServer
}

func (m *messageService) GetMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	json := &pb.MessageResponse{Message: "hi picpay"}
	return json, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	msg := &messageService{}
	pb.RegisterMessageServiceServer(srv, msg)
	panic(srv.Serve(lis))
}
