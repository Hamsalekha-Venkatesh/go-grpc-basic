package main

import (
	pb "grpc_server/proto/stream_gen"
	"io"
	"log"
	"sync"
	"time"
)

type BiDirectedServer struct {
	pb.UnimplementedStreamingServer
}

func (bidi *BiDirectedServer) ChatNow(chatStream pb.Streaming_ChatNowServer) error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := sendMessages(chatStream)
		if err != nil {
			log.Fatal("Error sending messages", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := receiveMessages(chatStream)
		if err != nil {
			log.Fatal("Error receiving messages", err)
			return
		}
	}()

	wg.Wait()
	return nil
}

func sendMessages(stream pb.Streaming_ChatNowServer) error {
	messages := []string{"Hello", "How're you", " I'm hungry"}
	for _, message := range messages {
		err := stream.SendMsg(&pb.ChatMessage{
			Message:   message,
			Timestamp: time.DateTime,
			PersonId:  99,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func receiveMessages(stream pb.Streaming_ChatNowServer) error {
	for {
		req, err := stream.Recv()
		log.Println("\n Received", req.GetMessage())
		if err == io.EOF {
			log.Println("client closed stream")
			break
		}

		if err != nil {
			return err
		}
	}
	return nil
}
