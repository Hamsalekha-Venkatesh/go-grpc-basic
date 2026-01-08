package main

import (
	"fmt"
	// pb "grpc_server/proto/gen"
	pbstream "grpc_server/proto/stream_gen"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listen.Close()

	//************* TRADITIONAL gPRC Server ***************************
	//credentials, err := credentials.NewServerTLSFromFile(cert, key)
	//if err != nil {
	//	fmt.Println("Error loading credentials: FAILED", err.Error())
	//	return
	//}

	// secure TLS implementing server
	//grpcServer := grpc.NewServer(grpc.Creds(credentials))
	//
	//pb.RegisterCalculateServer(grpcServer, &server{})
	//pb.RegisterGreeterServer(grpcServer, &GreeterServer{})
	//
	//fmt.Println("Traditional Server started. Listening on :50051")
	//err = grpcServer.Serve(listen)
	//if err != nil {
	//	fmt.Println("Error listener: FAILED", err.Error())
	//	log.Fatalln("Server failed...")
	//	return
	//}
	//************* TRADITIONAL gPRC Server ***************************

	//************* STREAMING gPRC Server ***************************

	streamingServer := grpc.NewServer()
	fmt.Println("Streaming Server started. Listening on :50052")

	//pbstream.RegisterCalculatorServer(streamingServer, &streamServer{})
	pbstream.RegisterStreamingServer(streamingServer, &BiDirectedServer{})

	if err := streamingServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	//************* STREAMING gPRC Server ***************************

}
