package client

import (
	"context"
	"fmt"
	pb "go_common/grpc/helloworld_demo/proto"
	"go_common/grpc/helloworld_demo/register_center"
	"log"
	"strconv"

	"google.golang.org/grpc/balancer/roundrobin"

	"google.golang.org/grpc"
)

func StartClientRegister() {

	//连接etcd,得到名命名空间
	schema, err := register_center.GenerateAndRegisterEtcdResolver("127.0.0.1:2379", "HelloService")
	if err != nil {
		log.Fatal("init etcd resolver err:", err.Error())
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:///HelloService", schema), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建客户端存根对象
	c := pb.NewHelloServiceClient(conn)
	//客户端发起调用，返回一个流
	r, err := c.HelloWorldClientAndServerStream(context.Background(), grpc.EmptyCallOption{})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	//用流给服务端发送消息
	for i := 0; i < 10; i++ {
		r.Send(&pb.HelloRequest{Request: "my is golang gRpc client " + strconv.Itoa(i)})
	}
	//流关闭
	r.CloseSend()

	//接受服务端返回的消息
	for {
		res, err := r.Recv()
		if err != nil && err.Error() == "EOF" {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
			break
		}
		log.Printf("result:%v", res.Response)
	}
	defer conn.Close()
}
