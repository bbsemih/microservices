package main

import (
	"io"
	"log"

	pb "github.com/bbsemih/microservices/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max was invoked with a streaming request!")
	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
			}
		}
	}
}
