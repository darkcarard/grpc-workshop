package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/api/client_streaming/clientstreamingpb"
	"grpc/common"
	"log"
	"time"
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

	c := clientstreamingpb.NewClientStreamingServiceClient(conn)

	//doAttackRequest(1, c)
	doAttackRequest(5, c)
}

func doAttackRequest(damage uint32, c clientstreamingpb.ClientStreamingServiceClient) {
	attack := &clientstreamingpb.Attack{Damage: damage}
	reqs := []*clientstreamingpb.ClientStreamingRequest{
		{BossID: 1, Attack: attack},
		{BossID: 1, Attack: attack},
		{BossID: 1, Attack: attack},
		{BossID: 1, Attack: attack},
		{BossID: 1, Attack: attack},
	}
	stream, err := c.AttackBoss(context.Background())

	if err != nil {
		log.Fatalf("Error calling ComputeAverage: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response from Server: %v", err)
	}

	bossHealth := res.GetBoss().GetHealth()

	if bossHealth <= 0 {
		common.PrintMessage("Boss killed!!!")

		return
	}

	common.PrintMessage(fmt.Sprintf("Boss attack result: %+v", res.GetBoss()))
}
