package tgrpc

import (
"context"
"log"
"net"

"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Heartbeat(ctx context.Context, in *HeartbeatRequest) (*HeartbeatResponse, error) {
	log.Printf("Received: %v", in.Id)
	return &HeartbeatResponse{Code:0, Message: "ok", Data:&HeartbeatResponse_DataStruct{Id:in.Id}}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) Handler1(ctx context.Context, in *Handler1Request) (*Handler1Response, error) {
	log.Printf("Received: %v", in.Id)
	return &Handler1Response{Code:0, Message:"ok", Data:&Handler1Response_DataStruct{Id:in.Id, Result:in.Left+in.Right}}, nil
}

func RunServer(){
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterMyServerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}