package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
	serverAddress := "localhost:8000"
	ctx := context.Background()
	srv := grpc.NewServer()
	router := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	msg := &messageService{}
	pb.RegisterMessageServiceServer(srv, msg)

	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		panic(err)
	}

	go func(srv *grpc.Server, listener net.Listener) {
		panic(srv.Serve(lis))
	}(srv, lis)

	if err = pb.RegisterMessageServiceHandlerFromEndpoint(ctx, router, serverAddress, opts); err != nil {
		panic(err)
	}

	println("starting http at 8001")
	panic(http.ListenAndServe("localhost:8001", router))
}
