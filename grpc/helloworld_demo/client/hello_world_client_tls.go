package client

import (
	"context"
	"fmt"
	pb "go_common/grpc/helloworld_demo/proto"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func StartClientTLS() {
	// TLS认证
	creds, err := credentials.NewClientTLSFromFile("E:\\server.pem", "")

	conn, err := grpc.Dial("127.0.0.1:8090", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println(err)
		return
	}
	c := pb.NewHelloServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.HelloWorldClientAndServerStream(ctx, grpc.EmptyCallOption{})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	for i := 0; i < 10; i++ {
		r.Send(&pb.HelloRequest{Request: "my is golang gRpc client " + strconv.Itoa(i)})
	}
	r.CloseSend()
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
