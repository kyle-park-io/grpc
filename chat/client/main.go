package main

import (
	"log"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "client/chat"
)

const (
	port    = ":50051"
	address = "localhost:50051"
)

var (
	r      pb.ChatClient
	ctx    context.Context
	cancel context.CancelFunc
)

func sendTest(wg *sync.WaitGroup) {
	defer wg.Done()

	v, err := r.SendMsgTest(ctx, &pb.TestMsg{Id: 1})
	if err != nil {
		log.Fatalf("could not send msg: %v", err)
	}
	log.Printf("send msg successfully: %s", v.Value)
}

func send(wg *sync.WaitGroup) {
	defer wg.Done()

	now := time.Now()
	eventTime := timestamppb.New(now)

	v, err := r.SendMsg(ctx, &pb.ChatMsg{Id: 2, UserId: "kyle", Text: "hello world!", EventTime: eventTime})
	if err != nil {
		log.Fatalf("could not send msg: %v", err)
	}
	log.Printf("send msg successfully: %s", v.Value)
}

func getTest(wg *sync.WaitGroup) {
	defer wg.Done()

	v, err := r.GetMsgTest(ctx, &pb.GetTestMsg{Id: 1})
	if err != nil {
		log.Fatalf("could not send msg: %v", err)
	}
	log.Printf("send msg successfully: %v", v)
}

func get(wg *sync.WaitGroup) {
	defer wg.Done()

	v, err := r.GetMsg(ctx, &pb.GetChatMsg{Id: 2, UserId: "kyle"})
	if err != nil {
		log.Fatalf("could not send msg: %v", err)
	}
	log.Printf("send msg successfully: %v", v)
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer conn.Close()
	r = pb.NewChatClient(conn)

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// send
	var wg1 sync.WaitGroup
	requestFuncs := []func(*sync.WaitGroup){
		sendTest,
		send,
	}
	for _, reqFunc := range requestFuncs {
		wg1.Add(1)
		go reqFunc(&wg1)
	}
	wg1.Wait()

	log.Println("First group of tasks completed. Starting second group of tasks...")
	// get
	var wg2 sync.WaitGroup
	requestFuncs = []func(*sync.WaitGroup){
		getTest,
		get,
	}
	for _, reqFunc := range requestFuncs {
		wg2.Add(1)
		go reqFunc(&wg2)
	}
	wg2.Wait()
}
