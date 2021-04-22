package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"org.spixi/spixi/greet/greetpb"
)

type server struct {}

func (*server) Sum(ctx context.Context, req *greetpb.SumRequest) (*greetpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", req)
	to_add := req.GetAddends()
	
	res := &greetpb.SumResponse{
		Result: to_add.Sum1 + to_add.Sum2,
	}
	return res, nil
}

func main() {
	fmt.Println("This is a sum Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	greetpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve the sum server: %v", err)
	}

}
