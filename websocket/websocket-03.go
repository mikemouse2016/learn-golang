package main

import (
	"io"
	"math/rand"
	"net/http"
	"time"
	//"log"
	"strconv"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
rand.Seed(42)
	for {
	i := rand.Intn(30)
		io.WriteString(ws, strconv.Itoa(i))
		//log.Println(strconv.Itoa(i))
		//io.WriteString(ws, "1")
		time.Sleep(500 * time.Millisecond)
	}

}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
