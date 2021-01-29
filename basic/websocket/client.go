package websocket

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

func ClientWebSocket() {
	var origin = "http://localhost:8080"
	var url = "wss://localhost:8080"

	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 4; i++ {
		message := []byte("\u0002{\"bussinessType\":1,\"messageType\":5}\u0003")
		_, err = ws.Write(message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Send: %s\n", message)

		var msg = make([]byte, 512)
		m, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receive: %s\n", msg[:m])
	}
	ws.Close() //关闭连接
}
