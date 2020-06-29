package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func main() {

	clientManager := ClientManager{
		Clients:    make(map[Client]bool),
		Broadcast:  make(chan Message),
		Typing:     make(chan Message),
		Register:   make(chan Client),
		Unregister: make(chan Client),
	}

	fs := http.FileServer(http.Dir("../chatapp/build"))

	http.Handle("/", fs)

	go clientManager.run()

	http.HandleFunc("/ws", clientManager.handleConnections)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}

	fmt.Println("http server started on :8000")

}
