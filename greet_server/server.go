package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"shawon1fb/grpc_basic/greet/greetpb/greetpb"
	"strconv"
	"time"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	res := &greetpb.GreetResponse{
		Result: "Hello " + firstName + " " + lastName,
	}
	return res, nil

}

func (*server) GreetManyTime(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimeServer) error {
	fmt.Printf("GreetManyTime function was invked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	for i := 0; i < 10; i++ {
		result := "hello " + firstName + " " + lastName + " number:=> " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func main() {
	fmt.Println("hello server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	tls := false
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	var srv *server = &server{}

	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
