package main

import (
	"context"
	"fmt"
	"net"

	"serverpc/dummy/dummypb" // Importa el paquete generado

	"google.golang.org/grpc"
)

type DummyServiceServer struct{}

func (s DummyServiceServer) Hello(f context.Context, x *dummypb.DummyRequest) (*dummypb.DummyResponse, error) {

	resp := &dummypb.DummyResponse{
		Name: x.Name,
	}
	return resp, nil
}

func main() {

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	server := grpc.NewServer()
	dummypb.RegisterDummyServiceServer(server, DummyServiceServer{})

	fmt.Println("Server started on port 50051")
	if err := server.Serve(listen); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
