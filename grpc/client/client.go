package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "goprojects/grpc/gen/proto"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestAPIClient(conn)
	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello, world!"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.Msg)
}
