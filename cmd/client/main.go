package main

import (
	"context"
	"fmt"

	grpcserver "github.com/Nirss/fibonacci/transport/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	opt := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:8081", opt...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := grpcserver.NewFibonacciServiceClient(conn)
	response, err := client.GetRange(context.Background(), &grpcserver.GetRangeRequest{From: 46, To: 1000})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Result)
}
