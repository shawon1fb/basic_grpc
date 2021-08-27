package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"shawon1fb/grpc_basic/greet/greetpb/greetpb"
	"strconv"
	"time"
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
	//doServerStreaming(c)
	doClintStreaming(c)

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

func doClintStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a client streaming RPC...")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v\n", err)
	}

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "shawon",
				LastName:  "fb",
			},
		},

		//&greetpb.LongGreetRequest{
		//	Greeting: &greetpb.Greeting{
		//		FirstName: "2",
		//		LastName:  "fb",
		//	},
		//},
		//&greetpb.LongGreetRequest{
		//	Greeting: &greetpb.Greeting{
		//		FirstName: "3",
		//		LastName:  "fb",
		//	},
		//},
	}

	for i := 0; i < 100; i++ {
		requests = append(requests, &greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: strconv.Itoa(i),
				LastName:  "fb",
			},
		})
	}

	/// iterate the slice and send individual requests
	for _, req := range requests {
		fmt.Printf("Sending req:%v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("LongGreetResponse : %v\n", res)

}
