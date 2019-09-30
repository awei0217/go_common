package server

import (
	"context"
	"go_common/grpc/helloworld_demo/impl"
	pb "go_common/grpc/helloworld_demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServerIntercept() {

	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//创建grpc拦截器
	serverIntercept1 := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//做自己想做的事
		err = do(ctx)
		if err != nil {
			return
		}
		// 校验通过后继续处理请求
		return handler(ctx, req)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(serverIntercept1))
	//将拦截器添加进去
	gRpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloServiceServer(gRpcServer, &impl.HelloServiceServer{})
	gRpcServer.Serve(lis)
}

/**
自定义方法，做自己想做的事，比如记录请求入参，出参
*/
func do(ctx context.Context) error {
	return nil
}

//创建拦截器链（思想是递归）
func InterceptChain(intercepts ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	//获取拦截器的长度
	l := len(intercepts)
	//如下我们返回一个拦截器
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//在这个拦截器中，我们做一些操作
		//构造一个链
		chain := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(currentCtx, currentReq, info, currentHandler)
			}
		}
		//声明一个handler
		chainHandler := handler
		for i := l - 1; i >= 0; i-- {
			//递归一层一层调用
			chainHandler = chain(intercepts[i], chainHandler)
		}
		//返回结果
		return chainHandler(ctx, req)
	}
}
