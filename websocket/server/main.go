// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto proto/product_info.proto --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/go/product_info
// Execute go run go/server/main.go

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "server/websocket"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement ecommerce/product_info.
type server struct {
	// orderMap map[string]*pb.Name

	// productMap map[string]*pb.Product
}

// AddProduct implements ecommerce.AddProduct
func (s *server) ServerToClientUnary(ctx context.Context, in *pb.Name) (*pb.Result, error) {
	// out, err := uuid.NewV4()
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	// }
	// in.Id = out.String()
	// if s.productMap == nil {
	// 	s.productMap = make(map[string]*pb.Product)
	// }
	// s.productMap[in.Id] = in
	// log.Printf("Product %v : %v - Added.", in.Id, in.Name)
	// return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()

	a := &pb.Result{Result: "Good"}
	return a, nil
}

func (s *server) ServerToClientStream(stream pb.WebSocket_ServerToClientStreamServer) error {

	a, _ := stream.Recv()
	fmt.Println(a)

	x := &pb.Result{Result: "hihi"}
	stream.Send(x)

	// (pb.WebSocket_ServerToClientStreamClient, error)

	return nil

}

func main() {

	// tcp open
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// rpc function receiver
	// receiver 를 수정해 보는 것도 재밌을 듯
	pb.RegisterWebSocketServer(s, &server{})

	// server open
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
