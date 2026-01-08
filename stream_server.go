package main

import (
	"fmt"
	pb "grpc_server/proto/stream_gen"
	"time"
)

type streamServer struct {
	pb.UnimplementedCalculatorServer
	pb.UnimplementedStreamingServer
}

// Server side streaming example...
func (ss *streamServer) GenerateFibonacci(request *pb.FibonacciRequest, stream pb.Streaming_GenerateFibonacciClient) error {
	n := request.MaxFibNumbers
	a, b := 0, 1

	for range n {
		err := stream.SendMsg(&pb.FibonacciResponse{
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
