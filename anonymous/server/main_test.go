package main

import (
	"context"
	"net"
	"sync"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "server/chat"
)

func TestConnect(t *testing.T) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	// run server
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})

	// open server goroutine
	errChan := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	// run client
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewChatServiceClient(conn)

	req := &pb.ConnectEvent{UserId: "kyle", Timestamp: timestamppb.Now()}
	res, err := client.Connect(context.Background(), req)
	if err != nil {
		t.Fatalf("could not greet: %v", err)
	}

	s.GracefulStop()
	wg.Wait()
	select {
	case err := <-errChan:
		t.Fatalf("server error: %v", err)
	default:
		t.Logf("%v", res)
	}
}
