package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/api/bidirectional/bidirectionalpb"
	"grpc/common"
	"io"
	"log"
	"time"
)

const (
	port = 50051
)

func main() {
	common.PrintMessage("gRPC Bidirectional Streaming client running")

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}

	defer conn.Close()

	c := bidirectionalpb.NewBidirectionalServiceClient(conn)

	//doAttackRequest(1, c)
	doAttackRequest(5, c)
}

func doAttackRequest(damage uint32, c bidirectionalpb.BidirectionalServiceClient) {
	attack := &bidirectionalpb.Attack{Damage: damage}
	reqs := []*bidirectionalpb.AttackBossRequest{
		{BossID: 1, Attack: attack},
		{BossID: 1, Attack: attack},
		{BossID: 1, Attack: attack},
		{BossID: 1, Attack: attack},
		{BossID: 1, Attack: attack},
	}

	stream, err := c.AttackBoss(context.Background())

	if err != nil {
		log.Fatalf("Error attacking the boss: %v", err)
	}

	// Create a channel
	waitC := make(chan struct{})

	// go routine for sending requests to the server
	go func() {
		for _, req := range reqs {
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	// go routine for receiving responses from the server
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error reading stream: %v", err)
			}

			if res.GetBoss().GetHealth() <= 0 {
				common.PrintMessage("Boss killed!!!")

				break
			}

			if res.GetPlayer().GetHealth() <= 0 {
				common.PrintMessage("Player killed!!!")

				break
			}

			common.PrintMessage(fmt.Sprintf("Boss attack result: %+v", res.GetBoss()))
			common.PrintMessage(fmt.Sprintf("Player attack result: %+v", res.GetPlayer()))
		}

		// Close the channel
		close(waitC)
	}()

	// Locks channel until it's closed
	<-waitC
}
