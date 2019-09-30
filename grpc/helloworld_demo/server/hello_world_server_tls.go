package server

import (
	"go_common/grpc/helloworld_demo/impl"
	pb "go_common/grpc/helloworld_demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func StartServerTLS() {

	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("E:\\server.pem", "E:\\server.key")
	gRpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloServiceServer(gRpcServer, &impl.HelloServiceServer{})
	gRpcServer.Serve(lis)
}
