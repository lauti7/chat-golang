package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"./api/models"
)

var upgrader = websocket.Upgrader{}

type UsersManager struct {
	OnlineUsers       map[uint]OnlineUser
	RegisterChannel   chan OnlineUser
	UnregisterChannel chan OnlineUser
}

type BroadcastMessage struct {
	ID        string     `json:"id,omitempty"`
	ChatID    uint        `json:"chat_id"`
	Sender    OnlineUser `json:"sender"`
	ReceiverUserID uint `json:"receiver_user_id"`
	Content   string     `json:"content,omitempty"`
	Type      string     `json:"type"`
	Timestamp string     `json:"timestamp,omitempty"`
}

type OnlineUser struct {
	SocketID      *websocket.Conn       `json:"socket_id,omitempty"`
	ChatChannel   chan BroadcastMessage `json:"omitempty"`
	TypingChannel chan BroadcastMessage `json:"omitempty"`
	Username      string                `json:"user_name"`
	UserID        uint                  `json:"user_id" form:"user_id"`
}

type OnlineUserMessage struct {
	Type   string `json:"type"`
	UserID uint `json:"user_id"`
}

func (manager *UsersManager) handleWS(w http.ResponseWriter, r *http.Request, authUser interface{}) {

	onlineUser := authUser.(OnlineUser)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	onlineUser.SocketID = ws

	defer ws.Close()

	manager.RegisterChannel <- onlineUser

	go onlineUser.chat(manager)

	for {
		fmt.Println("::handleConnections")
		var msg BroadcastMessage

		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("error: %v", err)
			delete(manager.OnlineUsers, onlineUser.UserID)
			break
		}

		msg.Sender = onlineUser

		receiverUser, online := manager.OnlineUsers[msg.ReceiverUserID]

		if online {
			switch msg.Type {
			case "chat":
				receiverUser.ChatChannel <- msg
			case "typing":
				receiverUser.TypingChannel <- msg
			}
		}

		newMessage := models.Message{
			ChatID: msg.ChatID,
			UserID: msg.Sender.UserID,
			Content: msg.Content,
			ReceiverID:
			msg.ReceiverUserID,
		}
		newMessage.CreateMessage()
		continue
	}
	manager.UnregisterChannel <- onlineUser

}

func (manager *UsersManager) registration() {
	for {
		select {
		case user := <-manager.RegisterChannel:
			manager.OnlineUsers[user.UserID] = user
			u := OnlineUserMessage{
				UserID: user.UserID,
				Type: "new_user_online",
			}
			manager.newUserOnline(u)
		case user := <-manager.UnregisterChannel:
			manager.OnlineUsers[user.UserID] = user
			u := OnlineUserMessage{
				UserID: user.UserID,
				Type: "new_user_offline",
			}
			manager.newUserOnline(u)
		}
	}
}

func (user *OnlineUser) chat(manager *UsersManager) {
	for {
		select {
		case msg := <-user.ChatChannel:
			user.receiveMsg(msg, manager)
		case msg := <-user.TypingChannel:
			user.receiveMsg(msg, manager)
		}
	}
}

func (user *OnlineUser) receiveMsg(msg BroadcastMessage, manager *UsersManager) {

	err := user.SocketID.WriteJSON(msg)

	if err != nil {
		log.Printf("Error Routine: %v", err)
		user.SocketID.Close()
		delete(manager.OnlineUsers, user.UserID)
	}
}

func (manager *UsersManager) newUserOnline(u OnlineUserMessage) {
	for _, user := range manager.OnlineUsers {
		if user.UserID != u.UserID {
			err := user.SocketID.WriteJSON(u)
			if err != nil {
				log.Printf("Error Routine: %v", err)
				user.SocketID.Close()
				delete(manager.OnlineUsers, user.UserID)
			}
		}
	}
}


func checkAndFindUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var onlineUser OnlineUser

		c.Bind(&onlineUser)

		username := models.GetUsername(onlineUser.UserID)

		onlineUser.Username = username
		onlineUser.ChatChannel = make(chan BroadcastMessage)
		onlineUser.TypingChannel = make(chan BroadcastMessage)

		c.Set("user", onlineUser)
		c.Next()
	}
}

//
// func (manager *UsersManager) send(msg Message) {
// 	for client := range manager.Clients {
// 		if client.SocketID != msg.Sender.SocketID {
// 			err := client.SocketID.WriteJSON(msg)
// 			if err != nil {
// 				log.Printf("Error Routine: %v", err)
// 				client.SocketID.Close()
// 				delete(manager.Clients, client)
// 			}
// 		}
// 	}
// }
