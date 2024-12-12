package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "server/chat"
)

const (
	port    = ":50051"
	address = "localhost:50051"
)

type server struct{}

func (s *server) Connect(ctx context.Context, in *pb.ConnectEvent) (*pb.ConnectEvent, error) {
	e := pb.ConnectEvent{UserId: in.UserId, IsConnected: true, Timestamp: timestamppb.Now()}
	return &e, nil
}

func runServer() {
	// tcp open
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})

	// server open
	log.Printf("Server is running on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	runServer()
}
