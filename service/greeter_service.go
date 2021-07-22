package service

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/sakiib/grpc-gateway-template/gen/go/helloworld"
)

type Greeter struct {
	pb.UnimplementedGreeterServer
}

func NewGreeter() *Greeter {
	return &Greeter{pb.UnimplementedGreeterServer{}}
}

func (g *Greeter) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := in.GetName()
	if name == "" {
		return nil, errors.New("failed to get user name")
	}

	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %s", name),
	}, nil
}
