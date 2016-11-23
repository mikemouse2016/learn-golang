package main

import (
	"math/rand"
	"net/http"
	"time"
	//"log"
	"strconv"

	"golang.org/x/net/websocket"
)

// Send random data to WebSocket.
func RandServer(ws *websocket.Conn) {
	defer ws.Close()
	rand.Seed(42)
	for {
		i := rand.Intn(100) - 50
		//io.WriteString(ws, strconv.Itoa(i))
		ws.Write([]byte(strconv.Itoa(i)))
		//log.Println(strconv.Itoa(i))
		time.Sleep(500 * time.Millisecond)
	}

}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/rand", websocket.Handler(RandServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
