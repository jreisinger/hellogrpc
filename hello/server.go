package hello

import "context"

// Step 5: Create the gRPC server by implementing the HelloServer interface from hello_grpc.pb.go.
type Server struct {
	UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello, " + in.GetName()}, nil
}
