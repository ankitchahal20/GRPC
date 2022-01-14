package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ankit/project/grpc/grpc-request-response/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s server) Add(ctc context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Println("X and Y values are : ", x, " : ", y)
	result := x + y
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	s := &server{}
	listner, err := net.Listen("tcp", ":50501")
	if err != nil {
		log.Fatal(err)
	}
	//Creating a stub instance
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, s)
	grpcServer.Serve(listner)
}
