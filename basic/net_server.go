package basic

import (
	"fmt"
	"io"
	"net"
)

func StartServer() {
	addr, _ := net.ResolveTCPAddr("tcp4", ":8080")
	listener, _ := net.ListenTCP("tcp4", addr)
	fmt.Println("服务器", listener.Addr().Network(), "启动成功")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受客户端连接错误:", err)
			continue
		}
		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	bs := make([]byte, 1024, 1024)
	for {
		n, err := conn.Read(bs)
		if err != nil {
			if err != io.EOF {
			}
			break
		}
		fmt.Println("读取到客户端数据:", string(bs), "数据大小:", n)
	}
	l, _ := conn.Write([]byte("服务端收到您发送的消息了"))
	fmt.Println("服务端发送给客户端消息成功,字节数为:", l)
	defer conn.Close()
}
