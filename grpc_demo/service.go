package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/proto"
	"net"
)

type server struct {
	proto.GreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	fmt.Printf(req.Name)
	return &proto.Response{Name: "name"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	//reflection.Register(s)

	defer func() {
		s.Stop()
		listen.Close()
	}()

	fmt.Println("Serving 8001...")
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
