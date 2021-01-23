package websocket

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

var origin = "http://localhost:8080"
var url = "wss://localhost:8080"

func ClientWebSocket() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	go sendMsg(ws)

	go receiveMsg(ws)

}

func receiveMsg(ws *websocket.Conn) {
	for {
		var msg = make([]byte, 512)
		m, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Receive:", string(msg[:m]))
	}

}

func sendMsg(ws *websocket.Conn) {
	for {
		message := []byte("\u0002{\"bussinessType\":1,\"messageType\":5}\u0003")
		_, err := ws.Write(message)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Second)
	}
}
