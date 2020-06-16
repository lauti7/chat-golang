package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)// broadcast channel
var typing = make(chan Message) //typing channel

var upgrader = websocket.Upgrader{}

type ClientManager struct {
  Clients map[*websocket.Conn]bool
  Broadcast chan Message
  Typing chan Message
  // Register chan
  // Unregister chan
}

type Message struct {
  Email    string `json:"email"`
  Username string `json:"username"`
  Message  string `json:"message"`
  Status   string `json:"status"`
  ClientID *websocket.Conn
}


func main() {

  clientManager := ClientManager{
    Clients: make(map[*websocket.Conn]bool),
    Broadcast: make(chan Message),
    Typing: make(chan Message),
  }


  fs := http.FileServer(http.Dir("../chatapp/build"))

  http.Handle("/", fs)

  http.HandleFunc("/ws", clientManager.handleConnections)

  go clientManager.handleBroadcastedMessages()
  go clientManager.handleUserTyping()

  err := http.ListenAndServe(":8000", nil)


  if err != nil {
          log.Fatal("ListenAndServe: ", err)
  }

  log.Println("http server started on :8000")

}

//entender READJSON. SE QUEDA ESPERANDO ? LOS CHANNELS CUANDO RECIBEN SON BLOQUENATES SE QUEDAN ESPERANDO

func (manager *ClientManager) handleConnections(w http.ResponseWriter, r *http.Request) {

  upgrader.CheckOrigin = func (r *http.Request) bool {return true}
  ws, err := upgrader.Upgrade(w, r, nil)

  if err != nil {
    log.Fatal(err)
  }

  defer ws.Close()

  manager.Clients[ws] = true

  for {
    fmt.Println("Recorriendo handleConnections")
    var msg Message

    fmt.Println("almost reading")
    err := ws.ReadJSON(&msg)


    msg.ClientID = ws

    if err != nil {
      log.Printf("error: %v", err)
      delete(clients, ws)
      break
    }
    fmt.Println("Message: %v", msg)

    if msg.Status == "typing" {
      manager.Typing <- msg
    } else {
      manager.Broadcast <- msg
    }
  }

  fmt.Println("Out handleConnections")


}

func (manager *ClientManager) handleBroadcastedMessages() {
  for {
    msg := <-manager.Broadcast
    fmt.Println(msg)

    for client := range manager.Clients {
      if client == msg.ClientID {
        continue
      }
      err := client.WriteJSON(msg)
      if err != nil {
        log.Printf("Broadcast Error Routine: %v", err)
        client.Close()
        delete(clients, client)
      }
    }
  }
}

func (manager *ClientManager) handleUserTyping(){
  for {
    msg := <-manager.Typing

    fmt.Println(msg)

    for client := range manager.Clients {
      err := client.WriteJSON(msg)
      if err != nil {
        log.Printf("Typing Error Routine: %v", err)
        client.Close()
        delete(clients, client)
      }
    }
  }
}
