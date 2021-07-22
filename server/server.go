package main

import (
	"flag"
	"fmt"
	"github.com/sakiib/grpc-gateway-template/gateway"
	pb "github.com/sakiib/grpc-gateway-template/gen/go/helloworld"
	"github.com/sakiib/grpc-gateway-template/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 0, "server port")
	flag.Parse()

	log.Printf("starting the server on port: %d", *port)

	grpcServer := grpc.NewServer()
	greeterServer := service.NewGreeter()

	pb.RegisterGreeterServer(grpcServer, greeterServer)

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen with: %s", err.Error())
	}

	go func() {
		log.Fatalf("grpc server error: %s", grpcServer.Serve(listener))
	}()

	log.Fatalf("gateway error: %s", gateway.Run("dns:///"+fmt.Sprintf("0.0.0.0:%d", *port)))
}
