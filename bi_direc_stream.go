package main

import (
	"fmt"
	pb "grpc_server/proto/stream_gen"
	"io"
	"log"
	"time"
)

type BiDirectedServer struct {
	pb.UnimplementedStreamingServer
	pb.UnimplementedCalculatorServer
}

func (bidi *BiDirectedServer) ChatNow(stream pb.Streaming_ChatNowServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Received: %v", req.GetMessage())
		// ⚠️ Blocking input (OK for demo only)
		var reply string
		fmt.Print("Enter reply: ")
		_, _ = fmt.Scanln(&reply)

		err = stream.SendMsg(pb.ChatMessage{
			PersonId:  100,
			Message:   reply,
			Timestamp: time.Now().Format(time.RFC3339),
		})
		if err != nil {
			return err
		}
	}
}
