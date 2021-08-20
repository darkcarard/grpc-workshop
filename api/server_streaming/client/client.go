package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/api/server_streaming/serverstreamingpb"
	"grpc/common"
	"io"
	"log"
)

const (
	port = 50051
)

func main() {
	common.PrintMessage("gRPC Server Stream client running")

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}

	defer conn.Close()

	c := serverstreamingpb.NewServerStreamingServiceClient(conn)

	//doAttackRequest(1, c)
	doAttackRequest(10, c)
}

func doAttackRequest(damage uint32, c serverstreamingpb.ServerStreamingServiceClient) {
	boss := &serverstreamingpb.Boss{
		Id:     1,
		Health: 5,
	}
	attack := &serverstreamingpb.Attack{
		Damage: damage,
	}
	req := &serverstreamingpb.AttackBossRequest{
		Boss:   boss,
		Attack: attack,
	}

	res, err := c.AttackBoss(context.Background(), req)

	if err != nil {
		log.Fatalf("An error ocurred Attacking the Boss: %v", err)
	}

	var player *serverstreamingpb.Player

	for {
		msg, err := res.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
		}

		boss = msg.GetBoss()
		player = msg.GetPlayer()
	}

	if boss.GetHealth() <= 0 {
		common.PrintMessage("Boss Killed!!!")

		return
	}

	if player.GetHealth() <= 0 {
		common.PrintMessage("Player Killed!!!")

		return
	}
}
