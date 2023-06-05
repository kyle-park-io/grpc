package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "client/websocket"
)

const (
	address = "localhost:50051"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWebSocketClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	r, err := c.ServerToClientUnary(ctx, &pb.Name{Description: "Kyle"})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}

	fmt.Println(r)

	ex, err := c.ServerToClientStream(ctx)
	a := &pb.Name{Description: "Stream"}
	ex.Send(a)

	aa, _ := ex.Recv()
	fmt.Println(aa)
	for {

		// c.ServerToClientStream()

	}

	// log.Printf("Product ID: %s added successfully", r.Value)

}
