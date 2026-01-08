package main

import (
	"context"
	pb "grpc_working/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct {
	pb.UnimplementedCalculateServer
}

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	if in.A == 0 || in.B == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "invalid argument")
	}

	return &pb.AddResponse{Sum: int64(in.A + in.B)}, nil
}

func (s *GreeterServer) HelloNamaste(ctx context.Context, in *pb.GreeterRequest) (*pb.GreeterResponse, error) {
	if in.Name == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "invalid argument")
	}

	var message = "Bonjour, " + in.Name + "||  Namaste, " + in.Name

	return &pb.GreeterResponse{
		Message: message,
	}, nil
}

func (s *GreeterServer) GoodBye(ctx context.Context, in *pb.GoodbyeRequest) (*pb.GoodbyeResponse, error) {
	if in.Name == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "invalid argument")
	}

	var goodBye = in.Name + "Will miss you, sayonara"
	return &pb.GoodbyeResponse{
		Message: goodBye,
	}, nil

}
