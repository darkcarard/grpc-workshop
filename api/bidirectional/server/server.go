package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/api/bidirectional/bidirectionalpb"
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

	bidirectionalpb.RegisterBidirectionalServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error listening: %v", err)
	}
}

func (*server) AttackBoss(stream bidirectionalpb.BidirectionalService_AttackBossServer) error {
	boss := &bidirectionalpb.Boss{
		Health: 10,
	}
	bossAttack := &bidirectionalpb.Attack{
		Damage: 2,
	}
	player := &bidirectionalpb.Player{
		Name:   "Player One",
		Health: 5,
	}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error reading client stream: %v", err)
		}

		common.PrintMessage(fmt.Sprintf("Request: %+v", req))

		boss.Id = req.GetBossID()
		boss.Health -= int32(req.GetAttack().GetDamage())
		player.Health -= int32(bossAttack.Damage)

		sendErr := stream.Send(
			&bidirectionalpb.AttackBossResponse{
				Boss:   boss,
				Player: player,
			},
		)

		if sendErr != nil {
			log.Fatalf("Error sending data to client: %v", err)

			return err
		}

		if boss.GetHealth() <= 0 || player.GetHealth() <= 0 {
			return nil
		}
	}
}
