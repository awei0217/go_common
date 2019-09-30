package client

import (
	"context"
	"fmt"
	pb "go_common/grpc/helloworld_demo/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

//自定义认证,实现 PerRPCCredentials 接口
type CustomerCredsAuth struct {
}

//获取元数据
func (c CustomerCredsAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "123",
		"appkey": "456",
	}, nil
}

//是否开启传输安全
func (c CustomerCredsAuth) RequireTransportSecurity() bool {
	return false
}

func StartClientToken() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(new(CustomerCredsAuth)))

	conn, err := grpc.Dial("127.0.0.1:8090", opts...)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := pb.NewHelloServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.HelloWorld(ctx, &pb.HelloRequest{})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	fmt.Println(r.Response)
	defer conn.Close()
}
