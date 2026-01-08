package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc_working/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	if in.A == 0 || in.B == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "invalid argument")
	}

	return &pb.AddResponse{Sum: int64(in.A + in.B)}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening: FAILED", err.Error())
		return
	}
	defer listen.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterCalculateServer(grpcServer, &server{})

	fmt.Println("Server started. Listening on :50051")
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("Error listener: FAILED", err.Error())
		log.Fatalln("Server failed...")
		return
	}
}
