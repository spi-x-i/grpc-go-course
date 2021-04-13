package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"org.spixi/spixi/greet/greetpb"
)

func main() {
	fmt.Println("Hello I'm the sum client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewSumServiceClient(cc)

	doSum(3, 20, c)
}

func doSum(a1 int32, a2 int32, c greetpb.SumServiceClient) {
	fmt.Println("Starting to do a Unary RPC sum...")
	req := &greetpb.SumRequest{
		Addends: &greetpb.Addends{
			Sum1: a1,
			Sum2: a2,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from sum: %v", res.Result)
}