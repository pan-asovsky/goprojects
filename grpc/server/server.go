package main

import (
	"context"
	"google.golang.org/grpc"
	pb "goprojects/grpc/gen/proto"
	"log"
	"net"
)

type testApiServer struct {
	pb.UnimplementedTestAPIServer
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestAPIServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}
