package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/api/unary/unarypb"
	"grpc/common"
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

	unarypb.RegisterUnaryServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error listening: %v", err)
	}
}

func (*server) AttackBoss(ctx context.Context, req *unarypb.AttackBossRequest) (*unarypb.AttackBossResponse, error) {
	common.PrintMessage(fmt.Sprintf("   Request: %+v", req))

	boss := req.GetBoss()
	attack := req.GetAttack()
	boss.Health -= int32(attack.GetDamage())

	res := &unarypb.AttackBossResponse{
		Boss: boss,
	}

	return res, nil
}
