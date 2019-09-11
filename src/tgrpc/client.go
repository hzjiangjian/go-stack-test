package tgrpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:8081"
	defaultName = "world"
)

func RunClient() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewMyServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Heartbeat(ctx, &HeartbeatRequest{Id: "0"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r)

	r1, err1 := c.Handler1(ctx, &Handler1Request{Id: "0", Left:1, Right:1})
	if err1 != nil {
		log.Fatalf("could not greet: %v", err1)
	}
	log.Printf("Greeting: %v", r1)
}

