package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ankit/project/grpc/grpc-request-response/proto"
	"google.golang.org/grpc"
)

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {

	req := proto.AddRequest{
		X: 100,
		Y: 201,
	}
	res, err := client.Add(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	output := res.GetResult()
	fmt.Println("Output : ", output)

}

func main() {
	clientConn, err := grpc.Dial("localhost:50501", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	// creating a proxy Instance
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	// Request && Response
	// SOme function Call here
	doRequestResponse(ctx, client)
}
