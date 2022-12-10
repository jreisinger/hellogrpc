package main

import (
	"hellogrpc/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Step 2: Listen on a TCP port.
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Step 6: Start the gRPC server on a TCP port.
	s := grpc.NewServer()
	hello.RegisterHelloServer(s, &hello.Server{})
	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
