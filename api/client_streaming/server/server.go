package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/api/client_streaming/clientstreamingpb"
	"grpc/common"
	"io"
	"log"
	"net"
)

const (
	port = 50051
)

type server struct{}

func main() {
	common.PrintMessage(fmt.Sprintf("gRPC server running in port: %d", port))

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))

	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}

	s := grpc.NewServer()

	clientstreamingpb.RegisterClientStreamingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error listening: %v", err)
	}
}

func (*server) AttackBoss(stream clientstreamingpb.ClientStreamingService_AttackBossServer) error {
	boss := &clientstreamingpb.Boss{Health: 10}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(
				&clientstreamingpb.ClientStreamingResponse{
					Boss: boss,
				},
			)
		}

		if err != nil {
			log.Fatalf("Error reading client stream: %v", err)
		}

		common.PrintMessage(fmt.Sprintf("Request: %+v", req))

		boss.Id = req.BossID
		boss.Health -= int32(req.GetAttack().GetDamage())

		if boss.GetHealth() <= 0 {
			return stream.SendAndClose(
				&clientstreamingpb.ClientStreamingResponse{
					Boss: boss,
				},
			)
		}
	}
}
