package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type ClientManager struct {
	Clients    map[Client]bool
	Broadcast  chan Message
	Typing     chan Message
	Register   chan Client
	Unregister chan Client
}

type Message struct {
	ID        string `json:"id,omitempty"`
	Sender    Client `json:"sender"`
	Message   string `json:"message,omitempty"`
	Type      string `json:"type"`
	Timestamp string `json:"timestamp,omitempty"`
}

type Client struct {
	SocketID *websocket.Conn `json:"socketId,omitempty"`
	Username string          `json:"userName"`
}

func (manager *ClientManager) handleConnections(c *gin.Context) {

	upgrader.CheckOrigin = func(r *c.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	client := Client{
		SocketID: ws,
	}

	fmt.Printf("Client: %v", client)

	socketerr := ws.ReadJSON(&client)

	fmt.Printf("Client: %v", client)

	if socketerr != nil {
		log.Printf("error: %v", err)
	}

	manager.Register <- client

	for {
		fmt.Println("Recorriendo handleConnections")
		var msg Message

		fmt.Println("almost reading")
		err := ws.ReadJSON(&msg)

		msg.Sender = client

		if err != nil {
			log.Printf("error: %v", err)
			delete(manager.Clients, client)
			break
		}
		fmt.Println("Message: %v", msg)

		switch msg.Type {
		case "broadcast":
			manager.Broadcast <- msg
		case "typing":
			manager.Typing <- msg
		}
	}

	manager.Unregister <- client

	fmt.Println("Out handleConnections")

}

func (manager *ClientManager) run() {
	for {
		select {
		case client := <-manager.Register:
			manager.Clients[client] = true
			newMsg := Message{
				Sender:  client,
				Message: "",
				Type:    "join",
			}
			manager.send(newMsg)
		case client := <-manager.Unregister:
			manager.Clients[client] = false
			newMsg := Message{
				Sender:  client,
				Message: "",
				Type:    "leave",
			}
			manager.send(newMsg)
		case msg := <-manager.Broadcast:
			manager.send(msg)
		case msg := <-manager.Typing:
			manager.send(msg)
		}
	}
}

func (manager *ClientManager) send(msg Message) {
	for client := range manager.Clients {
		if client.SocketID != msg.Sender.SocketID {
			err := client.SocketID.WriteJSON(msg)
			if err != nil {
				log.Printf("Error Routine: %v", err)
				client.SocketID.Close()
				delete(manager.Clients, client)
			}
		}
	}
}
