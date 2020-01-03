package server

import (
	"go_common/grpc/helloworld_demo/impl"
	pb "go_common/grpc/helloworld_demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServer() {
	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//创建一个grpc服务器对象
	gRpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(gRpcServer, &impl.HelloServiceServer{})
	//开启服务端
	gRpcServer.Serve(lis)
}
