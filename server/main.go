package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"pak/proto"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Message(ctx context.Context, req *proto.AReq) (*proto.BResp, error) {
	// Implement your server logic here
	fmt.Printf("Received request: %+v\n", req)

	// Dummy response
	resp := &proto.BResp{
		R: true,
	}

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterMyServiceServer(srv, proto.UnimplementedMyServiceServer{})
	//proto.RegisterMyServiceServer(srv, &server{})

	fmt.Println("Server is listening on port 50051...")
	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
