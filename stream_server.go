package main

import (
	"fmt"
	pb "grpc_working/proto/stream_gen"
	"time"
)

type streamServer struct {
	pb.UnimplementedCalculatorServer
}

func (ss *streamServer) GenerateFibonacci(request *pb.FibonacciRequest, stream pb.Calculator_GenerateFibonacciServer) error {
	n := request.MaxFibNumbers
	a, b := 0, 1

	for range n {
		err := stream.Send(&pb.FibonacciResponse{
			Number: int32(a + b),
		})
		if err != nil {
			fmt.Println("Error sending stream:", err)
			return err
		}

		a, b = b, a+b
		time.Sleep(1 * time.Second)
	}
	return nil // no error
}
