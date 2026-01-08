package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc_working/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
)

var (
	port = ":50051"
	cert = "cert.pem"
	key  = "key.pem"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	if in.A == 0 || in.B == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "invalid argument")
	}

	return &pb.AddResponse{Sum: int64(in.A + in.B)}, nil
}

func (s *server) HelloNamaste(ctx context.Context, in *pb.GreeterRequest) (*pb.GreeterResponse, error) {
	if in.Name == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "invalid argument")
	}

	var message = "Bonjour, " + in.Name + "||  Namaste, " + in.Name

	return &pb.GreeterResponse{
		Message: message,
	}, nil
}

func (s *server) GoodBye(ctx context.Context, in *pb.GoodbyeRequest) (*pb.GoodbyeResponse, error) {
	if in.Name == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "invalid argument")
	}

	var goodBye = in.Name + "Will miss you, sayonara"
	return &pb.GoodbyeResponse{
		Message: goodBye,
	}, nil

}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening: FAILED", err.Error())
		return
	}
	defer listen.Close()

	credentials, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		fmt.Println("Error loading credentials: FAILED", err.Error())
		return
	}

	// secure TLS implementing server
	grpcServer := grpc.NewServer(grpc.Creds(credentials))

	pb.RegisterCalculateServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &server{})

	fmt.Println("Server started. Listening on :50051")
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("Error listener: FAILED", err.Error())
		log.Fatalln("Server failed...")
		return
	}
}
