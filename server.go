package main

import (
	"context"
	"hellogrpc/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Step 4: Create the gRPC server by implementing the HelloServer interface from hello_grpc.pb.go.
type Server struct {
	hello.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func main() {
	// Step 5: Listen on a TCP port.
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Step 6: Start the gRPC server on the listening TCP port.
	s := grpc.NewServer()
	hello.RegisterHelloServer(s, &Server{})
	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
