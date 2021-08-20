package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/api/server_streaming/serverstreamingpb"
	"grpc/common"
	"log"
	"net"
	"time"
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

	serverstreamingpb.RegisterServerStreamingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error listening: %v", err)
	}
}

func (*server) AttackBoss(req *serverstreamingpb.AttackBossRequest,
	stream serverstreamingpb.ServerStreamingService_AttackBossServer) error {
	common.PrintMessage(fmt.Sprintf("Request: %+v", req))

	boss := req.GetBoss()
	boss.Health -= int32(req.GetAttack().GetDamage())

	if boss.GetHealth() <= 0 {
		time.Sleep(time.Second)
		stream.Send(
			&serverstreamingpb.AttackBossResponse{
				Boss:   boss,
				Player: &serverstreamingpb.Player{},
			},
		)
	} else {
		player := &serverstreamingpb.Player{
			Name:   "Player 1",
			Health: 10,
		}

		bossAttack := serverstreamingpb.Attack{
			Damage: 2,
		}

		for player.GetHealth() > 0 {
			common.PrintMessage("Attacking the player!!!")
			time.Sleep(time.Second)
			player.Health -= int32(bossAttack.GetDamage())

			stream.Send(
				&serverstreamingpb.AttackBossResponse{
					Boss:   boss,
					Player: player,
				},
			)
		}
	}

	return nil
}
