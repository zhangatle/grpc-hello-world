package main

import (
	context2 "context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-hello-world/proto"
	"log"
)

func main() {
	creeds, err := credentials.NewClientTLSFromFile("../certs/server.crt", "localhost")
	if err != nil {
		log.Printf("Failed to create TLS credentials %v\n", err)
	}
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(creeds))
	defer conn.Close()
	c := pb.NewHelloWorldClient(conn)
	context := context2.Background()
	body := &pb.HelloWorldRequest{
		Referer: "Grpc",
	}
	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
	}
	log.Println(r.Message)
}
