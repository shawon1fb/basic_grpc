package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"shawon1fb/grpc_basic/greet/greetpb/greetpb"
)

func main() {
	fmt.Println("Client started")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)

	//doUnary(c)
	doServerStreaming(c)

}
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "shahanul",
			LastName:  "haque",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a serverStreaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "shahanul",
			LastName:  "haque",
		},
	}
	resultStream, err := c.GreetManyTime(context.Background(), req)
	if err != nil {
		log.Fatalf("errr while calling GreetManyTime RPC: %v", err)
	}

	for {

		msg, err := resultStream.Recv()
		if err == io.EOF {
			/// stream ended successfully
			break
		}
		if err != nil {
			log.Fatalf("error while readig stream : %v", err)
		}
		log.Printf("Response from GreetMany: %v", msg.GetResult())
	}
}
