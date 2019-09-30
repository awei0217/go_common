package server

import (
	"context"
	"errors"
	"fmt"
	"go_common/grpc/helloworld_demo/impl"
	pb "go_common/grpc/helloworld_demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// 服务端方法实现中调用check方法校验
func check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("没有token细腻些")
	}

	fmt.Println(md, ok)
	var (
		appid  string
		appkey string
	)
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}
	if appid != "123" || appkey != "456" {
		return grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}
	fmt.Println(appid, appkey)
	return nil
}
func StartServerToken() {

	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(gRpcServer, &impl.HelloServiceServer{})
	gRpcServer.Serve(lis)
}
