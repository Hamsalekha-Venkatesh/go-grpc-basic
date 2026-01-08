package main

import (
	pb "grpc_working/proto/stream_gen"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type streamClient struct {
	pb.UnimplementedStreamingServer
}

// Client side streaming example...
func (sc *streamClient) SendClickStream(stream pb.Streaming_SendClickStreamServer) error {

	// Aggregation map adId -> Counters
	type perAdIdCounters struct {
		ipset        int32
		deviceClicks int32
	}

	aggMap := make(map[string]*perAdIdCounters) // Map[ad_id | counters]

	for { // keep receiving the request until EOF...
		req, err := stream.Recv()
		if err == io.EOF {
			break //client finished streaming ....
		}
		if err != nil {
			return status.Errorf(codes.Internal, "recv error: %v", err)
		}

		if _, ok := aggMap[req.AdId]; !ok { // if the map doesn't have adId..initialize this zero...
			aggMap[req.AdId] = &perAdIdCounters{
				ipset:        0,
				deviceClicks: 0,
			}
		}
		aggMap[req.AdId].deviceClicks++
		aggMap[req.AdId].ipset++
	}

	// return response for
	for adId, clickCounter := range aggMap {
		response := &pb.ClickResponse{
			AdId: adId,
			ClickCount: &pb.PerAdIdClickRequest{
				TotalDeviceClicks: clickCounter.deviceClicks,
				TotalIpClicks:     clickCounter.ipset,
			},
		}
		return stream.SendAndClose(response)
	}

	return nil
}
