package client

import (
	"context"
	"fmt"
	pb "go_common/grpc/helloworld_demo/proto"
	"log"

	"google.golang.org/grpc"
)

func StartClientIntercept() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("127.0.0.1:8090", opts...)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := pb.NewHelloServiceClient(conn)

	r, err := c.HelloWorld(context.Background(), &pb.HelloRequest{})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	fmt.Println(r.Response)
	defer conn.Close()
}
