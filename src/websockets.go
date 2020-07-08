package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

type UsersManager struct {
	OnlineUsers       map[uint]OnlineUser
	RegisterChannel   chan OnlineUser
	UnregisterChannel chan OnlineUser
}

type BroadcastMessage struct {
	ID        string     `json:"id,omitempty"`
	Sender    OnlineUser `json:"sender"`
	Content   string     `json:"content,omitempty"`
	Type      string     `json:"type"`
	Timestamp string     `json:"timestamp,omitempty"`
}

type OnlineUser struct {
	SocketID      *websocket.Conn       `json:"socketId,omitempty"`
	ChatChannel   chan BroadcastMessage `json:"omitempty"`
	TypingChannel chan BroadcastMessage `json:"omitempty"`
	Username      string                `json:"user_name"`
	UserID        uint                  `json:"user_id"`
}

func (manager *UsersManager) handleWS(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r c.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ws)

	defer ws.Close()

	// client := Client{
	// 	SocketID: ws,
	// }
	//
	// fmt.Printf("Client: %v", client)
	//
	// socketerr := ws.ReadJSON(&client)
	//
	// fmt.Printf("Client: %v", client)
	//
	// if socketerr != nil {
	// 	log.Printf("error: %v", err)
	// }
	//
	// manager.RegisterChannel <- client
	//
	// for {
	// 	fmt.Println("Recorriendo handleConnections")
	// 	var msg Message
	//
	// 	fmt.Println("almost reading")
	// 	err := ws.ReadJSON(&msg)
	//
	// 	msg.Sender = client
	//
	// 	if err != nil {
	// 		log.Printf("error: %v", err)
	// 		delete(manager.Clients, client)
	// 		break
	// 	}
	// 	fmt.Println("Message: %v", msg)
	//
	// 	switch msg.Type {
	// 	case "broadcast":
	// 		manager.Broadcast <- msg
	// 	case "typing":
	// 		manager.Typing <- msg
	// 	}
	// }
	//
	// manager.UnregisterChannel <- client
	//
	// fmt.Println("Out handleConnections")

}

func (manager *UsersManager) run() {
	for {
		select {
		case client := <-manager.RegisterChannel:
			manager.Clients[client] = true
			newMsg := Message{
				Sender:  client,
				Message: "",
				Type:    "join",
			}
			manager.send(newMsg)
		case client := <-manager.UnregisterChannel:
			manager.Clients[client] = false
			newMsg := Message{
				Sender:  client,
				Message: "",
				Type:    "leave",
			}
			manager.send(newMsg)
		// case msg := <-manager.Broadcast:
		// 	manager.send(msg)
		// case msg := <-manager.Typing:
		// 	manager.send(msg)
		}
	}
}

func (manager *UsersManager) send(msg Message) {
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
