Hellogrpc is a "hello world" program using gRPC.

gRPC is an open source Remote Procedure Call (RPC) framework by Google. RPC is an inter-process communication (IPC) technique in networked (distributed) computing. Remote procedure call is when a program calls a procedure (subroutine, function) to execute in a different address space, commonly on another computer, as if it were a normal local procedure call. Without the programmer explicitly coding the details for the remote interaction. 

Main usage scenarios
* efficiently connecting polyglot services in microservices style architecture
* connecting mobile devices, browser clients to backend services
* generating efficient client libraries

Core features
* idiomatic client libraries in 11 languages
* highly efficient on wire and with a simple service definition framework
* bi-directional streaming with http/2 based transport
* pluggable auth, tracing, loadbalancing and health checking

Used by: Square, Netflix, CoreOS, Cockroach Labs, Cisco, Juniper

# Components

![](https://grpc.io/img/landing-2.svg)

Service
* method1 (with parameters and return types)
* method2
* ...

Server
* implements service interface
* runs a gRPC server to handle client calls

Client
* has s stub (client) that provides the same methods as the server

# Protocol buffers

By default, gRPC uses Protocol Buffers, Google's open source mechanism for serializing structured data. You define the structure for the data in `.proto` text files. Then you use the protocol buffer compiler `protoc` to generate data access classes in your preferred language(s).

The recommended version to use is proto3.

# Steps to write gRPC program in Go

Search all repo files for "Step ".

Step 0: Create a go.mod file.
```
go mod init hellogrpc
```

Step 1: Install [protocol compiler](https://grpc.io/docs/languages/go/quickstart/#prerequisites).
```
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# if you haven't done this in ~/.bashrc or ~/.zshrc
export PATH="$PATH:$(go env GOPATH)/bin"
```

Step 4: Compile `hello.proto` into Go code (hello_grpc.pb.go, hello.pb.go).
```
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  hello/*.proto
```

Step 7: Run the server.
```
go mod tidy

go run .
```

Step 8: Test the server with [bloomrpc](https://github.com/bloomrpc/bloomrpc) GUI client.

See https://grpc.io/docs/languages/go/basics/ for more.