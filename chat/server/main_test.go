package main

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	pb "server/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSendMsg(t *testing.T) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})

	errChan := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := s.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewChatClient(conn)

	req := &pb.ChatMsg{Id: 7, UserId: "kyle-park", Text: "Hello World!", EventTime: timestamppb.New(time.Now())}
	res, err := client.SendMsg(context.Background(), req)
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
