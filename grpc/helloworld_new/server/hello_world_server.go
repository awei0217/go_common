package server

import (
	"context"
	pb "go_common/grpc/helloworld_new/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type HelloServiceServer struct {
}

func (*HelloServiceServer) HelloWorld(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("%v", req.Request)
	return &pb.HelloResponse{Response: "hello my is gRpcServer"}, nil
}
func (*HelloServiceServer) HelloWorldServerStream(req *pb.HelloRequest, srv pb.HelloService_HelloWorldServerStreamServer) error {
	log.Printf("%v", req.Request)
	srv.Send(&pb.HelloResponse{Response: "hello my is gRpcServer stream"})
	return nil
}
func (*HelloServiceServer) HelloWorldClientStream(srv pb.HelloService_HelloWorldClientStreamServer) error {
	for {
		req, err := srv.Recv()
		if err != nil && err.Error() == "EOF" {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
			break
		} else {
			log.Printf("%v", req.Request)
		}
	}
	srv.SendAndClose(&pb.HelloResponse{Response: "hello my is gRpcServer"})
	return nil
}
func (*HelloServiceServer) HelloWorldClientAndServerStream(srv pb.HelloService_HelloWorldClientAndServerStreamServer) error {
	for {
		req, err := srv.Recv()
		if err != nil && err.Error() == "EOF" {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
			break
		} else {
			log.Printf("%v", req.Request)
			time.Sleep(1 * time.Millisecond)
			srv.Send(&pb.HelloResponse{Response: "hello my is gRpcServer stream"})
		}
	}
	return nil
}

func StartServer() {

	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gRpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(gRpcServer, &HelloServiceServer{})
	gRpcServer.Serve(lis)
}
