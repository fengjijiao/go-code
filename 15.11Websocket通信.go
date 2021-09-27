package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
)

////websocket_server.go
func server(ws *websocket.Conn) {
	fmt.Println("new connection")
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			panic(err)
		}
	}
	println("closing connection\n")
	ws.Close()
}

func runServer() {
	http.Handle("/ws", websocket.Handler(server))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}


////websocket_client.go
func readFromServer(ws *websocket.Conn) {
	buf := make([]byte, 1000)
	for {
		if _, err := ws.Read(buf); err != nil {
			panic(err)
		}
	}
}

func runClient() {
	ws, err := websocket.Dial("ws://127.0.0.1:8000/ws", "", "http://127.0.0.1/")
	if err != nil {
		panic(err)
	}
	go readFromServer(ws)
	time.Sleep(5e9)
	ws.Close()
}