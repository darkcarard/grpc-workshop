package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/api/unary/unarypb"
	"grpc/common"
	"log"
	"time"
)

const (
	port = 50051
)

func main() {
	common.PrintMessage("gRPC Unary client running")

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}

	defer conn.Close()

	c := unarypb.NewUnaryServiceClient(conn)

	doAttackRequest(c)
}

func doAttackRequest(c unarypb.UnaryServiceClient) {
	boss := &unarypb.Boss{
		Id:     1,
		Health: 5,
	}
	attack := &unarypb.Attack{
		Damage: 1,
	}

	for {
		time.Sleep(time.Second)

		req := &unarypb.AttackBossRequest{
			Boss:   boss,
			Attack: attack,
		}

		res, err := c.AttackBoss(context.Background(), req)

		if err != nil {
			log.Fatalf("An error ocurred Attacking the Boss: %v", err)
		}

		boss = res.GetBoss()

		if boss.GetHealth() <= 0 {
			break
		}

		common.PrintMessage(fmt.Sprintf("Attack result: %+v", boss))
	}

	common.PrintMessage("Boss killed!!!")
}
