package basic

import (
	"bytes"
	"fmt"
	"net"
)

func StartClient() {

	conn, _ := net.Dial("tcp4", ":8080")
	len, _ := conn.Write([]byte("您好"))
	fmt.Println("发送成功,字节数为:", len)
	buffer := bytes.NewBuffer(make([]byte, 1024, 1024))
	l, _ := conn.Read(buffer.Bytes())
	fmt.Println("读取服务端数据:", buffer.String(), "字节数:", l)
	defer conn.Close()

}
