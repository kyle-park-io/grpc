package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "server/chat"
	"server/telegram"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	port = ":50051"
)

type server struct {
	testMsgMap map[string]*pb.TestMsg
	msgMap     map[string]*pb.ChatMsg
}

// send
func (s *server) SendMsgTest(ctx context.Context, in *pb.TestMsg) (*wrapperspb.StringValue, error) {
	log.Printf("receive test msg! : %v", in)

	if s.testMsgMap == nil {
		s.testMsgMap = make(map[string]*pb.TestMsg)
	}
	s.testMsgMap[strconv.FormatUint(in.Id, 10)] = in
	log.Printf("testMsgMap: %v", s.testMsgMap)

	return wrapperspb.String("ok"), nil
}

func (s *server) SendMsg(ctx context.Context, in *pb.ChatMsg) (*wrapperspb.StringValue, error) {
	log.Printf("receive msg! : %v", in)

	if s.msgMap == nil {
		s.msgMap = make(map[string]*pb.ChatMsg)
	}
	s.msgMap[strconv.FormatUint(in.Id, 10)] = in
	log.Printf("msgMap: %v", s.msgMap)

	err := telegram.SendMessage()
	if err != nil {
		log.Print(err)
	}

	return wrapperspb.String("ok"), nil
}

// get
func (s *server) GetMsgTest(ctx context.Context, in *pb.GetTestMsg) (*pb.TestMsg, error) {
	log.Printf("try go get test msg! : %v", in)

	fmt.Println(s.testMsgMap["1"])

	msg, exists := s.testMsgMap[strconv.FormatUint(in.Id, 10)]
	fmt.Println(msg)
	fmt.Println(exists)
	if exists && msg != nil {
		log.Printf("msg %v - Retrieved.", msg)
		return msg, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "msg does not exist.", in)
}

func (s *server) GetMsg(ctx context.Context, in *pb.GetChatMsg) (*pb.ChatMsg, error) {
	log.Printf("try go get msg! : %v", in)

	msg, exists := s.msgMap[strconv.FormatUint(in.Id, 10)]
	if exists && msg != nil {
		log.Printf("msg %v - Retrieved.", msg)
		return msg, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "msg does not exist.", in)
}

func main() {
	// tcp open
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})

	// server open
	log.Printf("Server is running on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
