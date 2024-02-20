package main

import (
	"context"
	"fmt"
	"log"

	"pak/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewMyServiceClient(conn)

	req := &proto.AReq{
		G: true,
		K: &proto.ToNest{
			NewField: 42,
			CoolEnum: &proto.SomeEnum{},
		},
		AlsoEnum: &proto.SomeEnum{},
	}

	resp, err := client.Message(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Message: %v", err)
	}

	fmt.Printf("Received response: %+v\n", resp)
}
