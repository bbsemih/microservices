package main

import (
	"context"
	"log"
	"time"

	pb "github.com/bbsemih/microservices/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked!")
	reqs := []*pb.GreetRequest{
		{FirstName: "Semih"},
		{FirstName: "Berkay"},
		{FirstName: "Öztürk"},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet Response: %v\n", res.Result)
}
